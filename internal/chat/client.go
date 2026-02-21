package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Supabase publishable (anon) key — safe to embed in client code.
// Security is enforced server-side via Row Level Security (RLS) policies
// on the "messages" table, not by keeping this key secret.
const (
	supabaseURL = "https://qfamiwiqxhrdtcipmeju.supabase.co"
	supabaseKey = "sb_publishable_jS1bqdOlGVqxu2feYxGXJA_ZZeZnfPD"
	tableName   = "messages"
	realtimeURL = "wss://qfamiwiqxhrdtcipmeju.supabase.co/realtime/v1/websocket?apikey=" + supabaseKey + "&vsn=1.0.0"
)

// Client is a Supabase chat client using REST + Realtime WebSocket.
type Client struct {
	mu       sync.Mutex
	messages []Message
	httpC    *http.Client
	online   bool
	sendMsg  func(interface{})
	wsConn   *websocket.Conn
	stopCh   chan struct{}
	stopOnce sync.Once
}

// NewMessagesMsg is sent when new messages arrive.
type NewMessagesMsg struct {
	Messages []Message
}

// ChatOfflineMsg is sent when chat can't connect.
type ChatOfflineMsg struct{}

// NewClient creates a new chat client.
func NewClient() *Client {
	return &Client{
		httpC:  &http.Client{Timeout: 10 * time.Second},
		online: false,
	}
}

// SetSendFunc sets the function used to send messages to bubbletea.
func (c *Client) SetSendFunc(fn func(interface{})) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sendMsg = fn
}

// Start loads initial messages via REST, then connects WebSocket for realtime.
func (c *Client) Start() {
	c.stopCh = make(chan struct{})

	go func() {
		// Load history via REST
		msgs, err := c.fetchRecent(100)
		if err != nil {
			c.mu.Lock()
			c.online = false
			c.mu.Unlock()
			c.send(ChatOfflineMsg{})
		} else {
			c.mu.Lock()
			c.messages = msgs
			c.online = true
			c.mu.Unlock()
			c.send(NewMessagesMsg{Messages: msgs})
		}

		// Realtime WebSocket loop with reconnect
		for {
			select {
			case <-c.stopCh:
				return
			default:
			}

			err := c.connectRealtime()
			if err != nil {
				log.Printf("realtime: %v", err)
			}

			c.mu.Lock()
			c.online = false
			c.mu.Unlock()
			c.send(ChatOfflineMsg{})

			// Wait before reconnect
			select {
			case <-time.After(3 * time.Second):
			case <-c.stopCh:
				return
			}
		}
	}()
}

// Stop stops the realtime connection.
func (c *Client) Stop() {
	c.stopOnce.Do(func() {
		if c.stopCh != nil {
			close(c.stopCh)
		}
		c.mu.Lock()
		ws := c.wsConn
		c.mu.Unlock()
		if ws != nil {
			ws.Close()
		}
	})
}

// SendMessage posts a message to the chat via REST.
func (c *Client) SendMessage(name, text string) error {
	payload := map[string]string{
		"name": name,
		"text": text,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", supabaseURL+"/rest/v1/"+tableName, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Prefer", "return=minimal")

	resp, err := c.httpC.Do(req)
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("send message: HTTP %d: %s", resp.StatusCode, string(b))
	}
	return nil
}

func (c *Client) send(msg interface{}) {
	c.mu.Lock()
	fn := c.sendMsg
	c.mu.Unlock()
	if fn != nil {
		fn(msg)
	}
}

// Phoenix channel message format
type phxMsg struct {
	Topic   string      `json:"topic"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
	Ref     string      `json:"ref,omitempty"`
}

type joinConfig struct {
	Config joinConfigInner `json:"config"`
}

type joinConfigInner struct {
	Broadcast       map[string]interface{} `json:"broadcast"`
	Presence        map[string]interface{} `json:"presence"`
	PostgresChanges []pgChangeFilter       `json:"postgres_changes"`
}

type pgChangeFilter struct {
	Event  string `json:"event"`
	Schema string `json:"schema"`
	Table  string `json:"table"`
}

func (c *Client) connectRealtime() error {
	conn, _, err := websocket.DefaultDialer.Dial(realtimeURL, nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	c.mu.Lock()
	c.wsConn = conn
	c.mu.Unlock()
	defer func() {
		c.mu.Lock()
		c.wsConn = nil
		c.mu.Unlock()
		conn.Close()
	}()

	topic := "realtime:public:" + tableName
	ref := 1

	// Send phx_join with postgres_changes subscription
	joinPayload := joinConfig{
		Config: joinConfigInner{
			Broadcast: map[string]interface{}{"self": false},
			Presence:  map[string]interface{}{"key": ""},
			PostgresChanges: []pgChangeFilter{
				{Event: "INSERT", Schema: "public", Table: tableName},
			},
		},
	}

	err = conn.WriteJSON(phxMsg{
		Topic:   topic,
		Event:   "phx_join",
		Payload: joinPayload,
		Ref:     strconv.Itoa(ref),
	})
	if err != nil {
		return fmt.Errorf("join: %w", err)
	}
	ref++

	c.mu.Lock()
	c.online = true
	c.mu.Unlock()

	// Heartbeat goroutine
	heartDone := make(chan struct{})
	go func() {
		defer close(heartDone)
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.mu.Lock()
				r := ref
				ref++
				c.mu.Unlock()
				err := conn.WriteJSON(phxMsg{
					Topic:   "phoenix",
					Event:   "heartbeat",
					Payload: map[string]interface{}{},
					Ref:     strconv.Itoa(r),
				})
				if err != nil {
					return
				}
			case <-c.stopCh:
				return
			}
		}
	}()

	// Read messages
	for {
		select {
		case <-c.stopCh:
			return nil
		default:
		}

		var raw json.RawMessage
		if err := conn.ReadJSON(&raw); err != nil {
			<-heartDone
			return fmt.Errorf("read: %w", err)
		}

		// Parse the Phoenix message
		var msg struct {
			Topic   string          `json:"topic"`
			Event   string          `json:"event"`
			Payload json.RawMessage `json:"payload"`
		}
		if err := json.Unmarshal(raw, &msg); err != nil {
			continue
		}

		if msg.Event == "postgres_changes" && msg.Topic == topic {
			c.handleRealtimeChange(msg.Payload)
		}
	}
}

func (c *Client) handleRealtimeChange(payload json.RawMessage) {
	// Payload structure: {"data": {"record": {...}, "type": "INSERT", ...}}
	var p struct {
		Data struct {
			Record struct {
				ID        json.Number `json:"id"`
				Name      string      `json:"name"`
				Text      string      `json:"text"`
				CreatedAt string      `json:"created_at"`
			} `json:"record"`
			Type string `json:"type"`
		} `json:"data"`
	}
	if err := json.Unmarshal(payload, &p); err != nil {
		return
	}
	if p.Data.Type != "INSERT" {
		return
	}

	id, _ := p.Data.Record.ID.Int64()
	createdAt, _ := time.Parse(time.RFC3339Nano, p.Data.Record.CreatedAt)

	msg := Message{
		ID:        int(id),
		Name:      p.Data.Record.Name,
		Text:      p.Data.Record.Text,
		CreatedAt: createdAt,
	}

	c.mu.Lock()
	c.messages = append(c.messages, msg)
	if len(c.messages) > 200 {
		c.messages = c.messages[len(c.messages)-200:]
	}
	allMsgs := make([]Message, len(c.messages))
	copy(allMsgs, c.messages)
	c.mu.Unlock()

	c.send(NewMessagesMsg{Messages: allMsgs})
}

func (c *Client) fetchRecent(limit int) ([]Message, error) {
	url := fmt.Sprintf("%s/rest/v1/%s?select=id,name,text,created_at&order=id.desc&limit=%d",
		supabaseURL, tableName, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)

	resp, err := c.httpC.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch recent: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("fetch recent: HTTP %d: %s", resp.StatusCode, string(b))
	}

	var msgs []Message
	if err := json.NewDecoder(resp.Body).Decode(&msgs); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	// Reverse (they come desc, we want asc)
	sort.Slice(msgs, func(i, j int) bool {
		return msgs[i].ID < msgs[j].ID
	})

	return msgs, nil
}

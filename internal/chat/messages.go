package chat

import (
	"fmt"
	"math/rand"
	"time"
)

// Message represents a chat message from Supabase.
type Message struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// GenerateAnonName generates a random anonymous nickname like "anon_1234".
func GenerateAnonName() string {
	digits := 2 + rand.Intn(3) // 2-4 digits
	max := 1
	for i := 0; i < digits; i++ {
		max *= 10
	}
	return fmt.Sprintf("cli_anon_%d", rand.Intn(max))
}

package data

import (
	"fmt"
	"math/rand"
	"time"
)

type Track struct {
	Title    string
	Duration time.Duration
	URL      string
}

type Album struct {
	Title  string
	Genre  string
	Tracks []Track
}

// trackURL builds a CDN URL for a standard Dopo Goto track.
func trackURL(album string, num int, title string) string {
	return fmt.Sprintf("https://cdn.dopogoto.com/Dopo Goto - %s/Dopo Goto - %s - %02d %s.mp3", album, album, num, title)
}

// saraDamarisURL builds a CDN URL for a Dopo Goto, Sara Damaris collaboration track.
func saraDamarisURL(album string, num int, title string) string {
	return fmt.Sprintf("https://cdn.dopogoto.com/Dopo Goto, Sara Damaris - %s/Dopo Goto, Sara Damaris - %s - %02d %s.mp3", album, album, num, title)
}

// Albums contains the full Dopo Goto catalog.
var Albums = []Album{
	// ── Album 1 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Pillbox",
		Genre: "AMB",
		Tracks: []Track{
			{Title: "A Song Number One", Duration: 0, URL: trackURL("The Songs From The Pillbox", 1, "A Song Number One")},
			{Title: "A Song To Wake Up To", Duration: 0, URL: trackURL("The Songs From The Pillbox", 2, "A Song To Wake Up To")},
			{Title: "A Song To Recall Memories Of Missed People", Duration: 0, URL: trackURL("The Songs From The Pillbox", 3, "A Song To Recall Memories Of Missed People")},
			{Title: "A Song To Stop Worrying About Your Future", Duration: 0, URL: trackURL("The Songs From The Pillbox", 4, "A Song To Stop Worrying About Your Future")},
			{Title: "A Song To Stop Worrying About Your Life Decisions", Duration: 0, URL: trackURL("The Songs From The Pillbox", 5, "A Song To Stop Worrying About Your Life Decisions")},
			{Title: "A Song To Accept Your Wrong Decisions", Duration: 0, URL: trackURL("The Songs From The Pillbox", 6, "A Song To Accept Your Wrong Decisions")},
			{Title: "A Song To Go Back Home Even Though It Is No Longer There", Duration: 0, URL: trackURL("The Songs From The Pillbox", 7, "A Song To Go Back Home Even Though It Is No Longer There")},
			{Title: "A Song To Feel A Relief From Pattern Seeking Mind", Duration: 0, URL: trackURL("The Songs From The Pillbox", 8, "A Song To Feel A Relief From Pattern Seeking Mind")},
			{Title: "A Song To Listen To At Half Past Three In The Morning", Duration: 0, URL: trackURL("The Songs From The Pillbox", 9, "A Song To Listen To At Half Past Three In The Morning")},
			{Title: "A Song To Take A Nap", Duration: 0, URL: trackURL("The Songs From The Pillbox", 10, "A Song To Take A Nap")},
			{Title: "A Song To Remember That Their Faces Were Wholly Burned And Their Eyesockets Were Hollow", Duration: 0, URL: trackURL("The Songs From The Pillbox", 11, "A Song To Remember That Their Faces Were Wholly Burned And Their Eyesockets Were Hollow")},
			{Title: "A Song To Develop Your Negative Trait", Duration: 0, URL: trackURL("The Songs From The Pillbox", 12, "A Song To Develop Your Negative Trait")},
			{Title: "A Song To Find The Main Core", Duration: 0, URL: trackURL("The Songs From The Pillbox", 13, "A Song To Find The Main Core")},
			{Title: "A Song To Enjoy Your Own Forgetfulness", Duration: 0, URL: trackURL("The Songs From The Pillbox", 14, "A Song To Enjoy Your Own Forgetfulness")},
			{Title: "A Song To Spot A Cognitive Distortion", Duration: 0, URL: trackURL("The Songs From The Pillbox", 15, "A Song To Spot A Cognitive Distortion")},
			{Title: "A Song To Repeat Your Empathic Failures", Duration: 0, URL: trackURL("The Songs From The Pillbox", 16, "A Song To Repeat Your Empathic Failures")},
			{Title: "A Song For Evolutionary Relief", Duration: 0, URL: trackURL("The Songs From The Pillbox", 17, "A Song For Evolutionary Relief")},
			{Title: "A Song Is Complex", Duration: 0, URL: trackURL("The Songs From The Pillbox", 18, "A Song Is Complex")},
			{Title: "A Song Is Displaced", Duration: 0, URL: trackURL("The Songs From The Pillbox", 19, "A Song Is Displaced")},
			{Title: "A Song To Find Yourself Lost", Duration: 0, URL: trackURL("The Songs From The Pillbox", 20, "A Song To Find Yourself Lost")},
			{Title: "A Song To Stop Self Torture", Duration: 0, URL: trackURL("The Songs From The Pillbox", 21, "A Song To Stop Self Torture")},
			{Title: "A Song To Forgive Yourself", Duration: 0, URL: trackURL("The Songs From The Pillbox", 22, "A Song To Forgive Yourself")},
			{Title: "A Song To Regret Your Life", Duration: 0, URL: trackURL("The Songs From The Pillbox", 23, "A Song To Regret Your Life")},
			{Title: "A Song To Calibrate A New Mindset", Duration: 0, URL: trackURL("The Songs From The Pillbox", 24, "A Song To Calibrate A New Mindset")},
			{Title: "A Song To Destroy An Implemented Values", Duration: 0, URL: trackURL("The Songs From The Pillbox", 25, "A Song To Destroy An Implemented Values")},
			{Title: "A Song To Enjoy Your Insomnia", Duration: 0, URL: trackURL("The Songs From The Pillbox", 26, "A Song To Enjoy Your Insomnia")},
			{Title: "A Song To Find Your Angst", Duration: 0, URL: trackURL("The Songs From The Pillbox", 27, "A Song To Find Your Angst")},
			{Title: "A Song For The Unremembered", Duration: 0, URL: trackURL("The Songs From The Pillbox", 28, "A Song For The Unremembered")},
			{Title: "A Song To Slow Down Time", Duration: 0, URL: trackURL("The Songs From The Pillbox", 29, "A Song To Slow Down Time")},
			{Title: "A Song To Listen Before The Storm", Duration: 0, URL: trackURL("The Songs From The Pillbox", 30, "A Song To Listen Before The Storm")},
			{Title: "A Song When Nothing Is Gonna Be The Same Again", Duration: 0, URL: trackURL("The Songs From The Pillbox", 31, "A Song When Nothing Is Gonna Be The Same Again")},
			{Title: "A Song To Fall In Love With Chat Bot", Duration: 0, URL: trackURL("The Songs From The Pillbox", 32, "A Song To Fall In Love With Chat Bot")},
			{Title: "A Song To Mourn Your Loved Ones", Duration: 0, URL: trackURL("The Songs From The Pillbox", 33, "A Song To Mourn Your Loved Ones")},
			{Title: "A Song, Where the Tears Turn Into Rivers", Duration: 0, URL: trackURL("The Songs From The Pillbox", 34, "A Song, Where the Tears Turn Into Rivers")},
			{Title: "A Song To Walk On The Edge Of Your Borderline Personality Disorder", Duration: 0, URL: trackURL("The Songs From The Pillbox", 35, "A Song To Walk On The Edge Of Your Borderline Personality Disorder")},
			{Title: "A Song To Repair Your Spirit", Duration: 0, URL: trackURL("The Songs From The Pillbox", 36, "A Song To Repair Your Spirit")},
			{Title: "A Song To Remember Your School Teacher", Duration: 0, URL: trackURL("The Songs From The Pillbox", 37, "A Song To Remember Your School Teacher")},
			{Title: "A Song Is A Prelude In E Minor Written By Dmitri Shostakovich", Duration: 0, URL: trackURL("The Songs From The Pillbox", 38, "A Song Is A Prelude In E Minor Written By Dmitri Shostakovich")},
			{Title: "A Song From The Empty Room", Duration: 0, URL: trackURL("The Songs From The Pillbox", 39, "A Song From The Empty Room")},
			{Title: "A Song Is Compulsive", Duration: 0, URL: trackURL("The Songs From The Pillbox", 40, "A Song Is Compulsive")},
			{Title: "A Song Is Born From Unanything", Duration: 0, URL: trackURL("The Songs From The Pillbox", 41, "A Song Is Born From Unanything")},
			{Title: "A Song Is Obsessive", Duration: 0, URL: trackURL("The Songs From The Pillbox", 42, "A Song Is Obsessive")},
			{Title: "A Song Of Conversion", Duration: 0, URL: trackURL("The Songs From The Pillbox", 43, "A Song Of Conversion")},
			{Title: "A Song So Light", Duration: 0, URL: trackURL("The Songs From The Pillbox", 44, "A Song So Light")},
			{Title: "A Song Never Seen", Duration: 0, URL: trackURL("The Songs From The Pillbox", 45, "A Song Never Seen")},
			{Title: "A Song So Pure", Duration: 0, URL: trackURL("The Songs From The Pillbox", 46, "A Song So Pure")},
			{Title: "A Song Before The Tide Turns Red", Duration: 0, URL: trackURL("The Songs From The Pillbox", 47, "A Song Before The Tide Turns Red")},
			{Title: "A Song For Watching Urban Decay", Duration: 0, URL: trackURL("The Songs From The Pillbox", 48, "A Song For Watching Urban Decay")},
			{Title: "A Song To Listen On 31st of August", Duration: 0, URL: trackURL("The Songs From The Pillbox", 49, "A Song To Listen On 31st of August")},
			{Title: "A Song To Say Good Night", Duration: 0, URL: trackURL("The Songs From The Pillbox", 50, "A Song To Say Good Night")},
			{Title: "A Song To The Siren", Duration: 0, URL: trackURL("The Songs From The Pillbox", 51, "A Song To The Siren")},
		},
	},
	// ── Album 2 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Hard Drive",
		Genre: "DNB",
		Tracks: []Track{
			{Title: "A Song That Goes Boom", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 1, "A Song That Goes Boom")},
			{Title: "A Song For Doing Long Division", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 2, "A Song For Doing Long Division")},
			{Title: "A Song To Rage Quit To", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 3, "A Song To Rage Quit To")},
			{Title: "A Song For When Water Tastes Weird And You Don\u2019t Know Why", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 4, "A Song For When Water Tastes Weird And You Don\u2019t Know Why")},
			{Title: "A Song To Donate Your Plasma", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 5, "A Song To Donate Your Plasma")},
			{Title: "A Song That Could Beat Up A Kangaroo", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 6, "A Song That Could Beat Up A Kangaroo")},
			{Title: "A Song For Watching Urban Decay featuring BIGSKULL", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 7, "A Song For Watching Urban Decay featuring BIGSKULL")},
			{Title: "A Song For Chewing", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 8, "A Song For Chewing")},
			{Title: "A Song To Journey Through The Eye Of A Needle", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 9, "A Song To Journey Through The Eye Of A Needle")},
			{Title: "A Song To Let Slip The Dogs Of War", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 10, "A Song To Let Slip The Dogs Of War")},
			{Title: "A Song To Caulk A Sink", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 11, "A Song To Caulk A Sink")},
			{Title: "A Song To Start Your Annual Tax Calculations", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 12, "A Song To Start Your Annual Tax Calculations")},
			{Title: "A Song For Netrunners featuring Posthuman Lab", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 13, "A Song For Netrunners featuring Posthuman Lab")},
			{Title: "A Song To Decrease Your Credit Rating To", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 14, "A Song To Decrease Your Credit Rating To")},
			{Title: "A Song To Butter The Royal Crumpets featuring William Smith", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 15, "A Song To Butter The Royal Crumpets featuring William Smith")},
			{Title: "A Song I Got Off Napster", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 16, "A Song I Got Off Napster")},
			{Title: "A Song To Expand Horizons To", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 17, "A Song To Expand Horizons To")},
			{Title: "A Song For An Eternal Deployment", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 18, "A Song For An Eternal Deployment")},
			{Title: "A Song To Gurn To", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 19, "A Song To Gurn To")},
			{Title: "A Song That Takes You Underwater", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 20, "A Song That Takes You Underwater")},
			{Title: "A Song For Low Gravity Systems", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 21, "A Song For Low Gravity Systems")},
			{Title: "A Song To Watch Dust Motions In A Sunbeam", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 22, "A Song To Watch Dust Motions In A Sunbeam")},
			{Title: "A Song To Clinically End Your Straightness", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 23, "A Song To Clinically End Your Straightness")},
			{Title: "A Song To Be Reyt", Duration: 0, URL: trackURL("The Songs From The Hard Drive", 24, "A Song To Be Reyt")},
		},
	},
	// ── Album 3 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Magnetic Core",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song For TJ", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 1, "A Song For TJ")},
			{Title: "A Song To Live To", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 2, "A Song To Live To")},
			{Title: "A Song For Shooting At Stars", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 3, "A Song For Shooting At Stars")},
			{Title: "A Song To See-Travel The World", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 4, "A Song To See-Travel The World")},
			{Title: "A Song To Crash Out", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 5, "A Song To Crash Out")},
			{Title: "A Song Made Of Liquid", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 6, "A Song Made Of Liquid")},
			{Title: "A Song To Do A Flip To", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 7, "A Song To Do A Flip To")},
			{Title: "A Song To Divide By Zero", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 8, "A Song To Divide By Zero")},
			{Title: "A Song To Ascend To", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 9, "A Song To Ascend To")},
			{Title: "A Song To Watch Paint Dry", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 10, "A Song To Watch Paint Dry")},
			{Title: "A Song To Escape featuring Sara Damaris", Duration: 0, URL: trackURL("The Songs From The Magnetic Core", 11, "A Song To Escape featuring Sara Damaris")},
		},
	},
	// ── Album 4 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Marine Biology Class Instructional DVD",
		Genre: "AMB",
		Tracks: []Track{
			{Title: "A Song Is Rolling Along With The Waves", Duration: 0, URL: trackURL("The Songs From The Marine Biology Class Instructional DVD", 1, "A Song Is Rolling Along With The Waves")},
			{Title: "A Song From Subaquatic Tidal Simulation", Duration: 0, URL: trackURL("The Songs From The Marine Biology Class Instructional DVD", 2, "A Song From Subaquatic Tidal Simulation")},
			{Title: "A Song That Subdivides The Ocean Surface", Duration: 0, URL: trackURL("The Songs From The Marine Biology Class Instructional DVD", 3, "A Song That Subdivides The Ocean Surface")},
			{Title: "A Song To Get A Sunburn", Duration: 0, URL: trackURL("The Songs From The Marine Biology Class Instructional DVD", 4, "A Song To Get A Sunburn")},
			{Title: "A Song To Be Washed Out To The Shore", Duration: 0, URL: trackURL("The Songs From The Marine Biology Class Instructional DVD", 5, "A Song To Be Washed Out To The Shore")},
			{Title: "A Song From Aqua Seafoam Dreams", Duration: 0, URL: trackURL("The Songs From The Marine Biology Class Instructional DVD", 6, "A Song From Aqua Seafoam Dreams")},
			{Title: "A Song Painted In Ultramarine Saturated Colors", Duration: 0, URL: trackURL("The Songs From The Marine Biology Class Instructional DVD", 7, "A Song Painted In Ultramarine Saturated Colors")},
			{Title: "A Song Of Calm Current", Duration: 0, URL: trackURL("The Songs From The Marine Biology Class Instructional DVD", 8, "A Song Of Calm Current")},
		},
	},
	// ── Album 5 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Disc Two",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song Called Genesis", Duration: 0, URL: trackURL("The Songs From The Disc Two", 1, "A Song Called Genesis")},
			{Title: "A Song To Reach For The Sky", Duration: 0, URL: trackURL("The Songs From The Disc Two", 2, "A Song To Reach For The Sky")},
			{Title: "A Song That Simulates Air Particles", Duration: 0, URL: trackURL("The Songs From The Disc Two", 3, "A Song That Simulates Air Particles")},
			{Title: "A Song Is Untitled", Duration: 0, URL: trackURL("The Songs From The Disc Two", 4, "A Song Is Untitled")},
			{Title: "A Song Weights Three-Quarters Of An Ounce", Duration: 0, URL: trackURL("The Songs From The Disc Two", 5, "A Song Weights Three-Quarters Of An Ounce")},
			{Title: "A Song Is Blue", Duration: 0, URL: trackURL("The Songs From The Disc Two", 6, "A Song Is Blue")},
			{Title: "A Song Of Oceans Deep With Hope", Duration: 0, URL: trackURL("The Songs From The Disc Two", 7, "A Song Of Oceans Deep With Hope")},
			{Title: "A Song To Learn To Fly", Duration: 0, URL: trackURL("The Songs From The Disc Two", 8, "A Song To Learn To Fly")},
			{Title: "A Song Of A Swallow's Tail", Duration: 0, URL: trackURL("The Songs From The Disc Two", 9, "A Song Of A Swallow's Tail")},
			{Title: "A Song To Overcome Space", Duration: 0, URL: trackURL("The Songs From The Disc Two", 10, "A Song To Overcome Space")},
			{Title: "A Song To Overcome Time", Duration: 0, URL: trackURL("The Songs From The Disc Two", 11, "A Song To Overcome Time")},
			{Title: "A Song Made Of Oxygène", Duration: 0, URL: trackURL("The Songs From The Disc Two", 12, "A Song Made Of Oxygène")},
			{Title: "A Song To Dream About The Future", Duration: 0, URL: saraDamarisURL("The Songs From The Disc Two", 13, "A Song To Dream About The Future")},
			{Title: "A Song With No Limit", Duration: 0, URL: trackURL("The Songs From The Disc Two", 14, "A Song With No Limit")},
			{Title: "A Song To Float Down Victoria Falls", Duration: 0, URL: trackURL("The Songs From The Disc Two", 15, "A Song To Float Down Victoria Falls")},
			{Title: "A Song To Remember", Duration: 0, URL: trackURL("The Songs From The Disc Two", 16, "A Song To Remember")},
		},
	},
	// ── Album 6 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs Are Non-Destructive",
		Genre: "IDM",
		Tracks: []Track{
			{Title: "A Song That Rejects Entropy", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 1, "A Song That Rejects Entropy")},
			{Title: "A Song She Danced To", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 2, "A Song She Danced To")},
			{Title: "A Song To Fall Out Of Orbit", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 3, "A Song To Fall Out Of Orbit")},
			{Title: "A Song To Offer A Brief Respite From Soul Crushing Reality", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 4, "A Song To Offer A Brief Respite From Soul Crushing Reality")},
			{Title: "A Song To Remember That Dream Where You Were Floating", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 5, "A Song To Remember That Dream Where You Were Floating")},
			{Title: "A Song To Drink Crystal Pepsi", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 6, "A Song To Drink Crystal Pepsi")},
			{Title: "A Song For Falling Into The Void", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 7, "A Song For Falling Into The Void")},
			{Title: "A Song For Daydreaming At Work", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 8, "A Song For Daydreaming At Work")},
			{Title: "A Song To Break Rocks To", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 9, "A Song To Break Rocks To")},
			{Title: "A Song To Dissolve In Water", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 10, "A Song To Dissolve In Water")},
			{Title: "A Song Of Crossroads You Can No Longer Find", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 11, "A Song Of Crossroads You Can No Longer Find")},
			{Title: "A Song For The Sentient Dust Motes In Your Eyelashes", Duration: 0, URL: trackURL("The Songs Are Non-Destructive", 12, "A Song For The Sentient Dust Motes In Your Eyelashes")},
		},
	},
	// ── Album 7 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Memory Card",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song To Hyper Light Drift", Duration: 0, URL: trackURL("The Songs From The Memory Card", 1, "A Song To Hyper Light Drift")},
			{Title: "A Song To Transcend Into The Digital Existence", Duration: 0, URL: trackURL("The Songs From The Memory Card", 2, "A Song To Transcend Into The Digital Existence")},
			{Title: "A Song To Save Your Game To", Duration: 0, URL: trackURL("The Songs From The Memory Card", 3, "A Song To Save Your Game To")},
			{Title: "A Song To Lay Your Head On", Duration: 0, URL: trackURL("The Songs From The Memory Card", 4, "A Song To Lay Your Head On")},
			{Title: "A Song To Recompile", Duration: 0, URL: trackURL("The Songs From The Memory Card", 5, "A Song To Recompile")},
			{Title: "A Song For The Books We Read When We Were Children", Duration: 0, URL: trackURL("The Songs From The Memory Card", 6, "A Song For The Books We Read When We Were Children")},
			{Title: "A Song That Makes You Feel", Duration: 0, URL: trackURL("The Songs From The Memory Card", 7, "A Song That Makes You Feel")},
			{Title: "A Song To Drink Aged Milk", Duration: 0, URL: trackURL("The Songs From The Memory Card", 8, "A Song To Drink Aged Milk")},
			{Title: "A Song For Giving Financial Advice", Duration: 0, URL: trackURL("The Songs From The Memory Card", 9, "A Song For Giving Financial Advice")},
			{Title: "A Song To Stare At The Sun A Little Too Long To", Duration: 0, URL: trackURL("The Songs From The Memory Card", 10, "A Song To Stare At The Sun A Little Too Long To")},
			{Title: "A Song To Remember To Forget", Duration: 0, URL: trackURL("The Songs From The Memory Card", 11, "A Song To Remember To Forget")},
			{Title: "A Song To Take A Shuttle Bus Home", Duration: 0, URL: trackURL("The Songs From The Memory Card", 12, "A Song To Take A Shuttle Bus Home")},
			{Title: "A Song That Causes You To Spontaneously Remember Your Dreams In Vivid Detail", Duration: 0, URL: trackURL("The Songs From The Memory Card", 13, "A Song That Causes You To Spontaneously Remember Your Dreams In Vivid Detail")},
			{Title: "A Song To Remember That Day Of Days", Duration: 0, URL: trackURL("The Songs From The Memory Card", 14, "A Song To Remember That Day Of Days")},
			{Title: "A Song For Our Hearts", Duration: 0, URL: saraDamarisURL("The Songs From The Memory Card", 15, "A Song For Our Hearts")},
		},
	},
	// ── Album 8 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs That Are Far Out",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song For Oni", Duration: 0, URL: trackURL("The Songs That Are Far Out", 1, "A Song For Oni")},
			{Title: "A Song To Feel Unreal", Duration: 0, URL: trackURL("The Songs That Are Far Out", 2, "A Song To Feel Unreal")},
			{Title: "A Song To Catch Feelings", Duration: 0, URL: trackURL("The Songs That Are Far Out", 3, "A Song To Catch Feelings")},
			{Title: "A Song To Fling Yourself Into The Myst", Duration: 0, URL: trackURL("The Songs That Are Far Out", 4, "A Song To Fling Yourself Into The Myst")},
			{Title: "A Song To Watch Lotus Bloom Twice", Duration: 0, URL: trackURL("The Songs That Are Far Out", 5, "A Song To Watch Lotus Bloom Twice")},
			{Title: "A Song To Play Before The War", Duration: 0, URL: trackURL("The Songs That Are Far Out", 6, "A Song To Play Before The War")},
			{Title: "A Song To Be Last Seen Online", Duration: 0, URL: trackURL("The Songs That Are Far Out", 7, "A Song To Be Last Seen Online")},
		},
	},
	// ── Album 9 ──────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Unknown Storage",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song To Experience Zero Gravity", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 1, "A Song To Experience Zero Gravity")},
			{Title: "A Song In Slow Motion", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 2, "A Song In Slow Motion")},
			{Title: "A Song To Format Your Memory Card", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 3, "A Song To Format Your Memory Card")},
			{Title: "A Song For Natalie", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 4, "A Song For Natalie")},
			{Title: "A Song To Make Out With Low Poly Girlfriend", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 5, "A Song To Make Out With Low Poly Girlfriend")},
			{Title: "A Song To Finally Sell Your Prius", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 6, "A Song To Finally Sell Your Prius")},
			{Title: "A Song For Spline Based Path Planning For Unmanned Air Vehicles", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 7, "A Song For Spline Based Path Planning For Unmanned Air Vehicles")},
			{Title: "A Song To Download WinAmp Skins To", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 8, "A Song To Download WinAmp Skins To")},
			{Title: "A Song About Times New Roman", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 9, "A Song About Times New Roman")},
			{Title: "A Song To Bring Down Mishima Zaibatsu", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 10, "A Song To Bring Down Mishima Zaibatsu")},
			{Title: "A Song For Outer Space", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 11, "A Song For Outer Space")},
			{Title: "A Song That You've Probably Heard In The Mighty Boosh", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 12, "A Song That You've Probably Heard In The Mighty Boosh")},
			{Title: "A Song From Outer Dark", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 13, "A Song From Outer Dark")},
			{Title: "A Song To Restart", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 14, "A Song To Restart")},
			{Title: "A Song To Take A Train To Manchester", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 15, "A Song To Take A Train To Manchester")},
			{Title: "A Song In Real Time", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 16, "A Song In Real Time")},
			{Title: "A Song For The Sunbleached Memories Of Us", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 17, "A Song For The Sunbleached Memories Of Us")},
			{Title: "A Song To Hydrate", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 18, "A Song To Hydrate")},
			{Title: "A Song To Defragment Your Hard Drive To", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 19, "A Song To Defragment Your Hard Drive To")},
			{Title: "A Song To Dunk Your Hot Poptarts In Cold Milk To", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 20, "A Song To Dunk Your Hot Poptarts In Cold Milk To")},
			{Title: "A Song Of Gamma Hydroxy Sensibility", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 21, "A Song Of Gamma Hydroxy Sensibility")},
			{Title: "A Song Of Beauty And Loss", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 22, "A Song Of Beauty And Loss")},
			{Title: "A Song To Be Lost And Never Found", Duration: 0, URL: trackURL("The Songs From The Unknown Storage", 23, "A Song To Be Lost And Never Found")},
		},
	},
	// ── Album 10 ─────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Disc 22",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song To Remember Saturday Morning Cartoons", Duration: 0, URL: trackURL("The Songs From The Disc 22", 1, "A Song To Remember Saturday Morning Cartoons")},
			{Title: "A Song To Shop At GameStop", Duration: 0, URL: trackURL("The Songs From The Disc 22", 2, "A Song To Shop At GameStop")},
			{Title: "A Song To Travel From Rouen To Paris", Duration: 0, URL: trackURL("The Songs From The Disc 22", 3, "A Song To Travel From Rouen To Paris")},
			{Title: "A Song To Date A Goth", Duration: 0, URL: trackURL("The Songs From The Disc 22", 4, "A Song To Date A Goth")},
			{Title: "A Song To Grind", Duration: 0, URL: trackURL("The Songs From The Disc 22", 5, "A Song To Grind")},
			{Title: "A Song To Rewatch Night Of The Living Dead", Duration: 0, URL: trackURL("The Songs From The Disc 22", 6, "A Song To Rewatch Night Of The Living Dead")},
			{Title: "A Song To Transcend", Duration: 0, URL: trackURL("The Songs From The Disc 22", 7, "A Song To Transcend")},
			{Title: "A Song That Was Made Specifically For Sony Walkman", Duration: 0, URL: trackURL("The Songs From The Disc 22", 8, "A Song That Was Made Specifically For Sony Walkman")},
			{Title: "A Song For Popping Bloons", Duration: 0, URL: trackURL("The Songs From The Disc 22", 9, "A Song For Popping Bloons")},
			{Title: "A Song To Fall In Love With Peter Steele", Duration: 0, URL: trackURL("The Songs From The Disc 22", 10, "A Song To Fall In Love With Peter Steele")},
			{Title: "A Song To Drive To The Ocean", Duration: 0, URL: trackURL("The Songs From The Disc 22", 11, "A Song To Drive To The Ocean")},
			{Title: "A Song To Keep Runnin' Away", Duration: 0, URL: trackURL("The Songs From The Disc 22", 12, "A Song To Keep Runnin' Away")},
			{Title: "A Song To Never Let Go", Duration: 0, URL: trackURL("The Songs From The Disc 22", 13, "A Song To Never Let Go")},
			{Title: "A Song To Accept Inevitable", Duration: 0, URL: trackURL("The Songs From The Disc 22", 14, "A Song To Accept Inevitable")},
			{Title: "A Song To Look On The Bright Side of Life", Duration: 0, URL: trackURL("The Songs From The Disc 22", 15, "A Song To Look On The Bright Side of Life")},
			{Title: "A Song To Feel Ecstatic", Duration: 0, URL: trackURL("The Songs From The Disc 22", 16, "A Song To Feel Ecstatic")},
			{Title: "A Song To Say Goodbye", Duration: 0, URL: trackURL("The Songs From The Disc 22", 17, "A Song To Say Goodbye")},
			{Title: "A Song When The Destination Is Unknown", Duration: 0, URL: trackURL("The Songs From The Disc 22", 18, "A Song When The Destination Is Unknown")},
			{Title: "A Song To Forgive All", Duration: 0, URL: trackURL("The Songs From The Disc 22", 19, "A Song To Forgive All")},
		},
	},
	// ── Album 11 ─────────────────────────────────────────────────────────
	{
		Title: "The Songs To Undress The Robot",
		Genre: "BRC",
		Tracks: []Track{
			{Title: "A Song To Have Intercourse With Horny Aliens", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 1, "A Song To Have Intercourse With Horny Aliens")},
			{Title: "A Song For Pink", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 2, "A Song For Pink")},
			{Title: "A Song To Get Knocked Down", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 3, "A Song To Get Knocked Down")},
			{Title: "A Song To Fine-Tune The Universe (Ding-a-Ling-Ding Dong)", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 4, "A Song To Fine-Tune The Universe (Ding-a-Ling-Ding Dong)")},
			{Title: "A Song To Burn On The Edge Of Something Beautiful", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 5, "A Song To Burn On The Edge Of Something Beautiful")},
			{Title: "A Song To Light The Purest Heart", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 6, "A Song To Light The Purest Heart")},
			{Title: "A Song To Order An Ok Pizza In Cleveland", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 7, "A Song To Order An Ok Pizza In Cleveland")},
			{Title: "A Song To Resist", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 8, "A Song To Resist")},
			{Title: "A Song To Subtract 303 From 909", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 9, "A Song To Subtract 303 From 909")},
			{Title: "A Song To Expand Your Mind", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 10, "A Song To Expand Your Mind")},
			{Title: "A Song That I Heard When I Saw God", Duration: 0, URL: trackURL("The Songs To Undress The Robot", 11, "A Song That I Heard When I Saw God")},
		},
	},
	// ── Album 12 ─────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Early 2000s Translucent Flash Drive",
		Genre: "DNB",
		Tracks: []Track{
			{Title: "A Song To Sing Along to the AOL Dial-Up Noise", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 1, "A Song To Sing Along to the AOL Dial-Up Noise")},
			{Title: "A Song To Stop Flickering and Clean Up Granny's Old Junk", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 2, "A Song To Stop Flickering and Clean Up Granny's Old Junk")},
			{Title: "A Song To Teleport to Hillwood", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 3, "A Song To Teleport to Hillwood")},
			{Title: "A Song To Get Frosted Tips", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 4, "A Song To Get Frosted Tips")},
			{Title: "A Song To Squat The Acid House", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 5, "A Song To Squat The Acid House")},
			{Title: "A Song To Be On ...", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 6, "A Song To Be On ...")},
			{Title: "A Song To Write A Letter From Too Far", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 7, "A Song To Write A Letter From Too Far")},
			{Title: "A Song To Fight Invisible Monsters", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 8, "A Song To Fight Invisible Monsters")},
			{Title: "A Song To Command & Conquer", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 9, "A Song To Command & Conquer")},
			{Title: "A Song To Chase Cel-Shaded Agent", Duration: 0, URL: trackURL("The Songs From The Early 2000s Translucent Flash Drive", 10, "A Song To Chase Cel-Shaded Agent")},
		},
	},
	// ── Album 13 ─────────────────────────────────────────────────────────
	{
		Title: "The Songs From The System Folder",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song To Listen To On Your Way To Blockbuster", Duration: 0, URL: trackURL("The Songs From The System Folder", 1, "A Song To Listen To On Your Way To Blockbuster")},
			{Title: "A Song To Y2K Panic", Duration: 0, URL: trackURL("The Songs From The System Folder", 2, "A Song To Y2K Panic")},
			{Title: "A Song To Forgive Grimes Coachella Set", Duration: 0, URL: trackURL("The Songs From The System Folder", 3, "A Song To Forgive Grimes Coachella Set")},
			{Title: "A Song To Vote For Bill Clinton", Duration: 0, URL: trackURL("The Songs From The System Folder", 4, "A Song To Vote For Bill Clinton")},
			{Title: "A Song To Remember That Limp Bizkit Are Cool", Duration: 0, URL: trackURL("The Songs From The System Folder", 5, "A Song To Remember That Limp Bizkit Are Cool")},
			{Title: "A Song To Stop Being A Drugstore Cowboy", Duration: 0, URL: trackURL("The Songs From The System Folder", 6, "A Song To Stop Being A Drugstore Cowboy")},
			{Title: "A Song To Upgrade Your Civic", Duration: 0, URL: trackURL("The Songs From The System Folder", 7, "A Song To Upgrade Your Civic")},
			{Title: "A Song To Spun", Duration: 0, URL: trackURL("The Songs From The System Folder", 8, "A Song To Spun")},
			{Title: "A Song To Hit Me One More Time", Duration: 0, URL: trackURL("The Songs From The System Folder", 9, "A Song To Hit Me One More Time")},
			{Title: "A Song To Call Your Grandmother", Duration: 0, URL: trackURL("The Songs From The System Folder", 10, "A Song To Call Your Grandmother")},
			{Title: "A Song To Runaway From Your Problems", Duration: 0, URL: trackURL("The Songs From The System Folder", 11, "A Song To Runaway From Your Problems")},
			{Title: "A Song To Feed All The Stray Cats", Duration: 0, URL: trackURL("The Songs From The System Folder", 12, "A Song To Feed All The Stray Cats")},
			{Title: "A Song To Depersonalize", Duration: 0, URL: trackURL("The Songs From The System Folder", 13, "A Song To Depersonalize")},
		},
	},
	// ── Album 14 ─────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Disc Three",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song To Nuttertools", Duration: 0, URL: trackURL("The Songs From The Disc Three", 1, "A Song To Nuttertools")},
			{Title: "A Song To Nailgun", Duration: 0, URL: trackURL("The Songs From The Disc Three", 2, "A Song To Nailgun")},
			{Title: "A Song To Jet Set Radio", Duration: 0, URL: trackURL("The Songs From The Disc Three", 3, "A Song To Jet Set Radio")},
			{Title: "A Song To GetPainkillers", Duration: 0, URL: trackURL("The Songs From The Disc Three", 4, "A Song To GetPainkillers")},
			{Title: "A Song To Use Impulse 101", Duration: 0, URL: trackURL("The Songs From The Disc Three", 5, "A Song To Use Impulse 101")},
			{Title: "A Song To Crash", Duration: 0, URL: trackURL("The Songs From The Disc Three", 6, "A Song To Crash")},
			{Title: "A Song To pointbreak", Duration: 0, URL: trackURL("The Songs From The Disc Three", 7, "A Song To pointbreak")},
		},
	},
	// ── Album 15 ─────────────────────────────────────────────────────────
	{
		Title: "The Songs From The Disc One",
		Genre: "JNG",
		Tracks: []Track{
			{Title: "A Song To Fall Through Textures", Duration: 0, URL: trackURL("The Songs From The Disc One", 1, "A Song To Fall Through Textures")},
			{Title: "A Song To Insert Disc 2", Duration: 0, URL: trackURL("The Songs From The Disc One", 2, "A Song To Insert Disc 2")},
			{Title: "A Song To Remember Aeon Flux", Duration: 0, URL: trackURL("The Songs From The Disc One", 3, "A Song To Remember Aeon Flux")},
			{Title: "A Song To Type The Motherlode Code", Duration: 0, URL: trackURL("The Songs From The Disc One", 4, "A Song To Type The Motherlode Code")},
			{Title: "A Song To Recall Childhood Memories", Duration: 0, URL: trackURL("The Songs From The Disc One", 5, "A Song To Recall Childhood Memories")},
			{Title: "A Song To Turn On Waypoint On Noclip", Duration: 0, URL: trackURL("The Songs From The Disc One", 6, "A Song To Turn On Waypoint On Noclip")},
			{Title: "A Song To Solve The Tibia Mysteries", Duration: 0, URL: trackURL("The Songs From The Disc One", 7, "A Song To Solve The Tibia Mysteries")},
			{Title: "A Song To Unpack A Brand New Motorola", Duration: 0, URL: trackURL("The Songs From The Disc One", 8, "A Song To Unpack A Brand New Motorola")},
			{Title: "A Song To Leave Behind All Of The Ghosts", Duration: 0, URL: trackURL("The Songs From The Disc One", 9, "A Song To Leave Behind All Of The Ghosts")},
		},
	},
}

// ShuffledAlbums returns a shuffled copy of the album list.
func ShuffledAlbums() []Album {
	shuffled := make([]Album, len(Albums))
	copy(shuffled, Albums)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

// FormatDuration formats a duration as M:SS or H:MM:SS.
func FormatDuration(d time.Duration) string {
	total := int(d.Seconds())
	if total >= 3600 {
		h := total / 3600
		m := (total % 3600) / 60
		s := total % 60
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
	m := total / 60
	s := total % 60
	return fmt.Sprintf("%d:%02d", m, s)
}

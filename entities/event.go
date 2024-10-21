package entities

import "time"

type Event struct {
	OrderType  string
	SessionID  string
	Card       string
	EventDate  time.Time
	WebsiteURL string
}

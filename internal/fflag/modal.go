package fflag

import "time"

// ToDo how to handle schedules?
type Flag struct {
	ID        string    `json:"id"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	// Schedule  *Schedule `json:"schedule"`
}

// type Schedule struct {
// 	ToState       bool      `json:"to_state"`
// 	EffectiveFrom time.Time `json:"effective_from"`
// }

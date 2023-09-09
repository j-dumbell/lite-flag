package fflag

import "time"

// ToDo how to handle schedules?
type Flag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Transition struct {
	ID            int
	FlagId        int
	ToState       bool
	EffectiveFrom time.Time
}

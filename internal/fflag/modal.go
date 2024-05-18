package fflag

// ToDo how to handle schedules?
type Flag struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

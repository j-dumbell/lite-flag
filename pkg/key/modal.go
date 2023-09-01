package key

import "time"

type ApiKey struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(name string) ApiKey {
	return ApiKey{Name: name, ApiKey: newKey(), CreatedAt: time.Now()}
}

package key

import "time"

type ApiKey struct {
	ID        string    `json:"id"`
	ApiKey    string    `json:"apiKey"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

// ToDo where should this live?
type ApiKeyRedacted struct {
	ID        string    `json:"id"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

func RemoveKey(apiKey ApiKey) ApiKeyRedacted {
	return ApiKeyRedacted{
		ID:        apiKey.ID,
		CreatedAt: apiKey.CreatedAt,
		Role:      apiKey.Role,
	}
}

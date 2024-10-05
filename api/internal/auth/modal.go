package auth

type Role string

const (
	RoleRoot     Role = "root"
	RoleAdmin    Role = "admin"
	RoleReadonly Role = "readonly"
)

var AllRoles = []Role{RoleRoot, RoleAdmin, RoleReadonly}

type ApiKey struct {
	Name string `json:"name"`
	Key  string `json:"apiKey"`
	Role Role   `json:"role"`
}

func (apiKey *ApiKey) Redact() ApiKeyRedacted {
	return ApiKeyRedacted{
		Name: apiKey.Name,
		Role: apiKey.Role,
	}
}

type ApiKeyRedacted struct {
	Name string `json:"name"`
	Role Role   `json:"role"`
}

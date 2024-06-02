package sdk

// Client is an API client.
type Client struct {
	host   string
	apiKey *string
}

// NewClient initializes a new Client.  If a non-nil apiKey is provided,
// this API key will be set in all requests.
func NewClient(host string, apiKey *string) Client {
	return Client{
		host:   host,
		apiKey: apiKey,
	}
}

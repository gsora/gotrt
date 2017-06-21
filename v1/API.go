package v1

const (
	// TheRockTrading base endpoint for v1 API
	baseEndpoint = "https://api.therocktrading.com/v1/"
)

// Session is a structure where API key and secret are contained.
// All the main methods used to interact with TheRockTrading's API will be referenced through this structure.
type Session struct {
	APIKey    string
	APISecret string
}

// NewSession returns a Session with API key and secret specified.
func NewSession(APIKey string, APISecret string) Session {
	return Session{APIKey, APISecret}
}

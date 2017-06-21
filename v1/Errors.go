package v1

// Error represents an error thrown by the API.
type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

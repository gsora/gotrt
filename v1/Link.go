package v1

// Link represents a link to a page.
type Link struct {
	Page int    `json:"page"`
	Href string `json:"href"`
}

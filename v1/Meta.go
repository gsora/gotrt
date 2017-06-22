package v1

// Meta represents informations about paging of some data.
type Meta struct {
	TotalCount int  `json:"total_count"`
	First      Link `json:"first"`
	Previous   Link `json:"previous"`
	Current    Link `json:"current"`
	Next       Link `json:"next"`
	Last       Link `json:"last"`
}

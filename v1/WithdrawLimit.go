package v1

// WithdrawLimit represents a limit in withdrawal of a given currency, for an account.
type WithdrawLimit struct {
	Currency  string  `json:"currency"`
	Initial   float64 `json:"initial"`
	Available float64 `json:"available"`
	Errors    []Error `json:"errors"`
}

// WithdrawLimits is an aggregation of WithdrawLimit.
type WithdrawLimits struct {
	WithdrawLimits []WithdrawLimit `json:"withdrawlimits"`
}

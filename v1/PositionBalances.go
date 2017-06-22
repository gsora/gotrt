package v1

import "time"

const (
	positionsBalancesEndpoint = "/funds/:fund_id/position_balances"
	positionsEndpoint         = "/funds/:fund_id/positions"
	positionEndpoint          = positionsEndpoint + "/:id"
)

// PositionBalances is an overall view of margin trading personal balances, both base and trade currency, for each fund/currency pairs.
type PositionBalances struct {
	Balances []Balance `json:"position_balances"`
}

// Position represents a position opened on a given account.
type Position struct {
	ID         int       `json:"id"`
	Status     string    `json:"status"`
	Date       time.Time `json:"date"`
	Amount     float64   `json:"amount"`
	BasePrice  float64   `json:"base_price"`
	Type       string    `json:"type"`
	Leverage   float64   `json:"leverage"`
	ProfitLoss float64   `json:"profit_loss"`
	TradeID    int       `json:"trade_id"`
	OrderID    int       `json:"order_id"`
}

// Positions is a list of your margin positions.
type Positions struct {
	Positions []Position `json:"positions"`
	Meta      Meta       `json:"meta"`
}

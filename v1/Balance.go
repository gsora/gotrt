package v1

import (
	"encoding/json"

	"github.com/gsora/gotrt/v1/currency"
)

const (
	balancesEndpoint = "/balances/"
)

// Balance represents a balance for an account, for a single currency.
type Balance struct {
	Currency       string  `json:"currency"`
	Balance        float64 `json:"balance"`
	TradingBalance float64 `json:"trading_balance"`
	Errors         []Error `json:"errors"`
}

// Balances is an aggregation of Balance.
type Balances struct {
	Balances []Balance `json:"balances"`
}

// BalancePerCurrency returns the account's balance for the currency specified.
func (s *Session) BalancePerCurrency(currency currency.Currency) (Balance, error) {
	response, err := s.doTRTRequest(baseEndpoint + balancesEndpoint + currency.String())
	defer response.Close()

	if err != nil {
		return Balance{}, err
	}

	balanceDecoder := json.NewDecoder(response)
	var balance Balance
	err = balanceDecoder.Decode(&balance)
	if err != nil {
		return balance, err
	}

	return balance, nil
}

// Balances returns all the account balances.
func (s *Session) Balances() (Balances, error) {
	response, err := s.doTRTRequest(baseEndpoint + balancesEndpoint)
	defer response.Close()

	if err != nil {
		return Balances{}, err
	}

	balancesDecoder := json.NewDecoder(response)
	var balances Balances
	err = balancesDecoder.Decode(&balances)
	if err != nil {
		return balances, err
	}

	return balances, nil
}

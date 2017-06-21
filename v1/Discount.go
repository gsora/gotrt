package v1

import (
	"encoding/json"

	"github.com/gsora/gotrt/v1/currency"
)

const (
	discountsEndpoint = "/discounts/"
)

// Discount represents a discount for a given cryptocurrency for an account.
type Discount struct {
	Currency string  `json:"currency"`
	Discount float64 `json:"discount"`
	Errors   []Error `json:"errors"`
}

// Discounts is an aggregation of Discount.
type Discounts struct {
	Discounts []Discount `json:"discounts"`
}

// DiscountPerCurrency returns the account's balance for the currency specified.
func (s *Session) DiscountPerCurrency(currency currency.Currency) (Discount, error) {
	response, err := s.doTRTRequest(baseEndpoint + balancesEndpoint + currency.String())
	defer response.Close()

	if err != nil {
		return Discount{}, err
	}

	balanceDecoder := json.NewDecoder(response)
	var balance Discount
	err = balanceDecoder.Decode(&balance)
	if err != nil {
		return balance, err
	}

	return balance, nil
}

// Discounts returns all the account balances.
func (s *Session) Discounts() (Discounts, error) {
	response, err := s.doTRTRequest(baseEndpoint + balancesEndpoint)
	defer response.Close()

	if err != nil {
		return Discounts{}, err
	}

	balancesDecoder := json.NewDecoder(response)
	var balances Discounts
	err = balancesDecoder.Decode(&balances)
	if err != nil {
		return balances, err
	}

	return balances, nil
}

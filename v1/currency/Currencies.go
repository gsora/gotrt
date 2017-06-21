package currency

//go:generate stringer -type=Currency

// Currency is a... currency.
type Currency int

const (
	// BTC = Bitcoin
	BTC Currency = iota
	// ETH = Ethereum
	ETH
	// LTC = Litecoin
	LTC
	// PPC = Peercoin
	PPC
	// ZEC = Zcash
	ZEC
)

package currency

// Fund represents a tuple of currencies.
type Fund int

const (
	// BTCEUR = BTC->EUR
	BTCEUR Fund = iota
	// BTCUSD = BTC->USD
	BTCUSD
	// LTCEUR = LTC->EUR
	LTCEUR
	// LTCBTC = LTC->BTC
	LTCBTC
	// BTCXRP = BTC->XRP
	BTCXRP
	// EURXRP = EUR->XRP
	EURXRP
	// USDXRP = USD->XRP
	USDXRP
	// PPCEUR = PPC->EUR
	PPCEUR
	// PPCBTC = PPC->BTC
	PPCBTC
	// ETHEUR = ETH->EUR
	ETHEUR
	// ETHBTC = ETH->BTC
	ETHBTC
	// ZECBTC = ZEC->BTC
	ZECBTC
	// ZECEUR = ZEC->EUR
	ZECEUR
)

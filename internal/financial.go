package financial

type Asset struct {
	Name          string
	Symbol        string
	// Currency      Currency
	// Holding       Holding
	Quote         Quote
	// QuoteSource   QuoteSource
	// Exchange      Exchange
	// Meta          Meta
}


// // Currency is the original and converted currency if applicable
// type Currency struct {
// 	// Code is the original currency code of the asset
// 	FromCurrencyCode string
// 	// CodeConverted is the currency code that pricing and values have been converted into
// 	ToCurrencyCode string
// }

type Quote struct {
	Price            float64
	PricePrevClose   float64
	PriceOpen        float64
	PriceDayHigh     float64
	PriceDayLow      float64
	Change           float64
	ChangePercent    float64
	FiftyTwoWeekHigh float64
	FiftyTwoWeekLow  float64
	MarketCap        float64
	Volume           float64
}


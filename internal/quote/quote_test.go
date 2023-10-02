package quote_test

import (
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	g "github.com/onsi/gomega/gstruct"
	. "github.com/wakatara/mgbq/internal/quote"
	. "github.com/wakatara/mgbq/test/http"
)

func mockResponseCurrencyGOOG() {
	response := `{
		"quoteResponse": {
			"result": [
				{
					"currency": "USD",
					"symbol": "GOOG"
				}
			],
			"error": null
		}
	}
	`
	responseURL := `=~\/finance\/quote.*symbols\=GOOG`
	httpmock.RegisterResponder("GET", responseURL, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, response)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
}

func mockResponseCurrencyEURUSD() {
	response := `{
		"quoteResponse": {
			"result": [
				{
					"quoteType": "CURRENCY",
					"quoteSourceName": "Delayed Quote",
					"currency": "EUR",
					"regularMarketPrice": {"raw": 0.8891,"fmt": "0.8891"},
					"sourceInterval": 15,
					"exchangeDataDelayedBy": 0,
					"exchange": "CCY",
					"fullExchangeName": "CCY",
					"symbol": "USDEUR=X"
				}
			],
			"error": null
		}
	}
	
	`
	responseURL := `=~\/finance\/quote.*symbols\=USDEUR.*X`

	httpmock.RegisterResponder("GET", responseURL, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, response)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
}

var _ = Describe("Quote", func() {

	var (
		dep f.Dependencies
	)

	BeforeEach(func() {
		dep = f.Dependencies{
			HttpClients: f.DependenciesHttpClients{
				Default: client,
				Yahoo:   client,
			},
		}
		MockResponseYahooQuotes()
	})

	Describe("GetAssetGroupQuote", func() {

		It("should get price quotes for each asset based on it's data source", func() {

			input := f.AssetGroup{
				SymbolsBySource: []c.AssetGroupSymbolsBySource{
					{
						Source: f.QuoteSourceYahoo,
						Symbols: []string{
							"GOOG",
							"RBLX",
						},
					},
					{
						Source: f.QuoteSourceCoingecko,
						Symbols: []string{
							"bitcoin",
							"solana",
						},
					},
					{
						Source: f.QuoteSourceUserDefined,
						Symbols: []string{
							"CASH",
							"PRIVATESHARES",
						},
					},
				},
			}
			output := GetAssetGroupQuote(dep)(input)

			Expect(output).To(g.MatchFields(g.IgnoreExtras, g.Fields{
				"AssetQuotes": HaveLen(2),
			}))

		})

	})

	Describe("GetAssetGroupsCurrencyRates", func() {

		It("should get currency conversion rates for each type of data source", func() {

			mockResponseCurrencyEURUSD()
			mockResponseCurrencyGOOG()
			input := []c.AssetGroup{
				{
					SymbolsBySource: []c.AssetGroupSymbolsBySource{
						{
							Source: f.QuoteSourceYahoo,
							Symbols: []string{
								"GOOG",
							},
						},
					},
				},
			}
			output, _ := GetAssetGroupsCurrencyRates(*client, input, "EUR")
			Expect(output).To(g.MatchAllKeys(g.Keys{
				"USD": g.MatchFields(g.IgnoreExtras, g.Fields{
					"FromCurrency": Equal("USD"),
					"ToCurrency":   Equal("EUR"),
					"Rate":         Equal(0.8891),
				}),
			}))

		})

	})

})

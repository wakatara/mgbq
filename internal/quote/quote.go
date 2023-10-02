package quote

import (
	"github.com/go-resty/resty/v2"
	quoteYahoo "github.com/wakatara/mgbq/internal/quote/yahoo"
)

func getQuoteBySource(dep f.Dependencies, symbolBySource f.AssetGroupSymbolsBySource) []c.AssetQuote {

	if symbolBySource.Source == f.QuoteSourceYahoo {
		return quoteYahoo.GetAssetQuotes(*dep.HttpClients.Yahoo, symbolBySource.Symbols)()
	}

	if symbolBySource.Source == f.QuoteSourceCoingecko {
		return quoteCoingecko.GetAssetQuotes(*dep.HttpClients.Yahoo, symbolBySource.Symbols)
	}

	return []c.AssetQuote{}
}

// GetAssetGroupQuote gets price quotes for groups of assets by data source
func GetAssetGroupQuote(dep f.Dependencies) func(c.AssetGroup) f.AssetGroupQuote {

	return func(assetGroup f.AssetGroup) f.AssetGroupQuote {

		var assetQuotes []c.AssetQuote

		for _, symbolBySource := range assetGroup.SymbolsBySource {

			assetQuotebySource := getQuoteBySource(dep, symbolBySource)
			assetQuotes = append(assetQuotes, assetQuotebySource...)

		}

		return f.AssetGroupQuote{
			AssetQuotes: assetQuotes,
			AssetGroup:  assetGroup,
		}
	}
}

func getUniqueSymbolsBySource(assetGroups []c.AssetGroup) []c.AssetGroupSymbolsBySource {

	symbols := make(map[c.QuoteSource]map[string]bool)
	symbolsUnique := make(map[c.QuoteSource][]string)
	assetGroupSymbolsBySource := make([]c.AssetGroupSymbolsBySource, 0)
	for _, assetGroup := range assetGroups {

		for _, symbolGroup := range assetGroup.SymbolsBySource {

			for _, symbol := range symbolGroup.Symbols {

				source := symbolGroup.Source

				if symbols[source] == nil {
					symbols[source] = map[string]bool{}
				}

				if !symbols[source][symbol] {
					symbols[source][symbol] = true
					symbolsUnique[source] = append(symbolsUnique[source], symbol)
				}
			}

		}

	}

	for source, symbols := range symbolsUnique {
		assetGroupSymbolsBySource = append(assetGroupSymbolsBySource, f.AssetGroupSymbolsBySource{
			Source:  source,
			Symbols: symbols,
		})
	}

	return assetGroupSymbolsBySource

}

// GetAssetGroupsCurrencyRates gets the currency rates by source across all asset groups
func GetAssetGroupsCurrencyRates(client resty.Client, assetGroups []c.AssetGroup, targetCurrency string) (c.CurrencyRates, error) {

	var err error
	var currencyRates f.CurrencyRates
	uniqueSymbolsBySource := getUniqueSymbolsBySource(assetGroups)

	for _, source := range uniqueSymbolsBySource {

		if source.Source == f.QuoteSourceYahoo && err == nil {
			currencyRates, err = quoteYahoo.GetCurrencyRates(client, source.Symbols, targetCurrency)
		}

	}

	return currencyRates, err
}

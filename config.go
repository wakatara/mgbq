package main

import (
	"fmt"

	"github.com/adrg/xdg"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

// Should be in config not common
type Config struct {
	Watchlist                         []string           `yaml:"watchlist"`
	// Lots                              []Lot              `yaml:"lots"`
	// Separate                          bool               `yaml:"show-separator"`
	// ExtraInfoExchange                 bool               `yaml:"show-tags"`
	// ExtraInfoFundamentals             bool               `yaml:"show-fundamentals"`
	// ShowSummary                       bool               `yaml:"show-summary"`
	// ShowHoldings                      bool               `yaml:"show-holdings"`
	// Proxy                             string             `yaml:"proxy"`
	// Sort                              string             `yaml:"sort"`
	// Currency                          string             `yaml:"currency"`
	// CurrencyConvertSummaryOnly        bool               `yaml:"currency-summary-only"`
	// CurrencyDisableUnitCostConversion bool               `yaml:"currency-disable-unit-cost-conversion"`
	// ColorScheme                       ConfigColorScheme  `yaml:"colors	// AssetGroup                        []ConfigAssetGroup `yaml:"groups"`
}


client := yahooClient.New(resty.New(), resty.New())
yahooClient.RefreshSession(client, resty.New())

func getConfigPath(fs afero.Fs, configPathOption string) (string, error) {
	var err error
	if configPathOption != "" {
		return configPathOption, nil
	}

	home, _ := homedir.Dir()

	v := viper.New()
	v.SetFs(fs)
	v.SetConfigType("yaml")
	v.AddConfigPath(home)
	v.AddConfigPath("./config")
	v.AddConfigPath(xdg.ConfigHome)
	v.AddConfigPath(xdg.ConfigHome + "/mgbq")
	v.SetConfigName("mgbq.config")
	err = v.ReadInConfig()

	if err != nil {
		return "", fmt.Errorf("invalid config: %w", err)
	}

	return v.ConfigFileUsed(), nil
}






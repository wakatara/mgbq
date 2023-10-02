/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mgbq",
	Short: "Moats, goats, boats, and quotes. Stock tracker, picker, and optimizer.",
	Long: `mgbq is a command line tool to help track, analyze and pick, and use
algorithms to maximize returns while constraining risk.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)
	// rootCmd.Flags().StringVar(&configPath, "config", "", "config file (default is $HOME/.ticker.yaml)")
	// rootCmd.Flags().StringVarP(&options.Watchlist, "watchlist", "w", "", "comma separated list of symbols to watch")
	// rootCmd.Flags().BoolVar(&options.Separate, "show-separator", false, "layout with separators between each quote")
	// rootCmd.Flags().BoolVar(&options.ExtraInfoExchange, "show-tags", false, "display currency, exchange name, and quote delay for each quote")
	// rootCmd.Flags().BoolVar(&options.ExtraInfoFundamentals, "show-fundamentals", false, "display open price, high, low, and volume for each quote")
	// rootCmd.Flags().BoolVar(&options.ShowSummary, "show-summary", false, "display summary of total gain and loss for positions")
	// rootCmd.Flags().BoolVar(&options.ShowHoldings, "show-holdings", false, "display average unit cost, quantity, portfolio weight")
	// rootCmd.Flags().StringVar(&options.Sort, "sort", "", "sort quotes on the UI. Set \"alpha\" to sort by ticker name. Set \"value\" to sort by position value. Keep empty to sort according to change percent")
	// rootCmd.AddCommand(printCmd)
	// printCmd.Flags().StringVar(&optionsPrint.Format, "format", "", "output format for printing holdings. Set \"csv\" to print as a CSV or \"json\" for JSON. Defaults to JSON.")
	// printCmd.Flags().StringVar(&configPath, "config", "", "config file (default is $HOME/.ticker.yaml)")
}

func initConfig() {
	// dep = cli.GetDependencies()
	// ctx, err = cli.GetContext(dep, options, configPath)

}


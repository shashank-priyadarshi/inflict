package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "inflict [amount] [year]",
	Short: "Input an amount and the year it was earned, to fetch the current value of the amount",
	Args:  cobra.MinimumNArgs(2),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Invalid value. Please provide an amount and the year it was earned in.")
			cmd.Help()
			os.Exit(1)
		}

		country, _ := cmd.Flags().GetStringSlice("countries")
		inflationType, _ := cmd.Flags().GetStringSlice("types")

		ctx := context.WithValue(context.Background(), "countries", country)
		ctx = context.WithValue(ctx, "types", inflationType)
		cmd.SetContext(ctx)
	},
	Run: func(cmd *cobra.Command, args []string) {

		var year_string string
		var amount float64
		calculatorArgs := make(map[string][]string)

		countries, types := cmd.Context().Value("countries"), cmd.Context().Value("types")

		if countries != nil {
			calculatorArgs["countries"] = countries.([]string)
		}

		if types != nil {
			calculatorArgs["types"] = types.([]string)
		}

		year_string = args[1]

		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("Invalid value. Please provide a valid amount.")
			os.Exit(1)
		}

		if err := parse(); err != nil {
			fmt.Println("error while parse inflation data: %v", err)
			return
		}
		calculator(amount, year_string, calculatorArgs)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version of inflict",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) != 0 {
			cmd.Help()
			return fmt.Errorf("trailing args detected")
		}

		fmt.Printf("0.0.1\n") // TODO: Semantic versioning
		return nil
	},
}

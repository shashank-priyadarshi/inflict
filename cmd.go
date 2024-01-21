package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "inflict [country] [value] [year] [inflation_type]",
	Short: "Input a country name, an amount, the year it was earned and the type of inflation to adjust for, to fetch the current value of the amount",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 4 {
			fmt.Println("Invalid value. Please provide")
			os.Exit(1)
		}

		var sheet sheet_type
		var country, year_string string
		var amount float64

		sheet, country, year_string = sheet_type(args[0]), args[1], args[2]

		amount, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			fmt.Println("Invalid value. Please provide a valid integer.")
			os.Exit(1)
		}

		if err := parse(); err != nil {
			fmt.Println("error while parse inflation data: %v", err)
			return
		}
		calculator(sheet, country, year_string, amount)
	},
}

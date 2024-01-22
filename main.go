package main

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
	"os"
)

func main() {

	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "help",
		Short:  "Help for inflict",
		Hidden: false,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	})

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(versionCmd)

	rootCmd.PersistentFlags().StringP("version", "v", "", "")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "")
	rootCmd.PersistentFlags().StringSliceP("types", "t", []string{"ccpi_a_e"}, "Specify one or more inflation types")
	rootCmd.PersistentFlags().StringSliceP("countries", "c", []string{"GLOBAL"}, "Specify one or more countries")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

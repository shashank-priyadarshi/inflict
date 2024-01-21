package main

import (
	"fmt"
	_ "github.com/spf13/cobra"
	"os"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-inventory",
	Short: "A CLI tool for stock keeping",
	Long: `Go Inventory is a tool that allows users to store item data.
  It gives yours a simple way to keep track of stock.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(storeCmd, addCmd, subtractCmd, deleteCmd, updateCmd)
	rootCmd.Flags().BoolP("toogle", "t", false, "Help message for toogle")
}

package main

import (
	"github.com/spf13/cobra"
	"github.com/jr-selphius/starwars-apiclient-go/internal/cli"
)

func main() {
	rootCmd := &cobra.Command{Use: "starwars-cli"}
	rootCmd.AddCommand(cli.InitPlanetsCmd())
	rootCmd.Execute()
}
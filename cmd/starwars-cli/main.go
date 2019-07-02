package main

import (
	"github.com/jr-selphius/starwars-apiclient-go/internal/cli"
	"github.com/jr-selphius/starwars-apiclient-go/internal/fetching"
	"github.com/jr-selphius/starwars-apiclient-go/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	fetchingService := fetching.NewService()
	csvRepository := csv.NewCsvRepository()

	rootCmd := &cobra.Command{Use: "starwars-cli"}
	rootCmd.AddCommand(cli.InitPlanetsCmd(fetchingService, csvRepository))
	rootCmd.Execute()
}

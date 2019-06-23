package main

import (
	"github.com/jr-selphius/starwars-apiclient-go/internal"
	"github.com/jr-selphius/starwars-apiclient-go/internal/cli"
	"github.com/jr-selphius/starwars-apiclient-go/internal/storage/csv"
	"github.com/jr-selphius/starwars-apiclient-go/internal/storage/remote"
	"github.com/spf13/cobra"
)

func main() {

	var remoteRepository internal.PlanetRepo
	remoteRepository = remote.NewRemoteRepository()

	var csvRepository internal.PlanetRepo
	csvRepository = csv.NewCsvRepository()

	rootCmd := &cobra.Command{Use: "starwars-cli"}
	rootCmd.AddCommand(cli.InitPlanetsCmd(remoteRepository, csvRepository))
	rootCmd.Execute()
}

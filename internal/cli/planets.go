package cli

import (
	"fmt"
	"github.com/jr-selphius/starwars-apiclient-go/internal"
	"github.com/jr-selphius/starwars-apiclient-go/internal/storage/remote"
	"github.com/jr-selphius/starwars-apiclient-go/internal/storage/csv"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

const idFlag string = "id"

func InitPlanetsCmd() *cobra.Command {
	planetsCmd := &cobra.Command{
		Use:   "planets",
		Short: "Save planet to csv file",
		Run:   runPlanetsFn(),
	}

	planetsCmd.Flags().StringP(idFlag, "i", "", "id of the planet")

	return planetsCmd
}

const APIEndpoint string = "https://swapi.co/api/"
const APIResource string = "planets/"

func runPlanetsFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {

		IDResource, _ := cmd.Flags().GetString(idFlag)

		if IDResource == "" {
			IDResource = "1"
		}

		var remoteRepository internal.PlanetRepo
		remoteRepository = remote.NewRemoteRepository()
		planet, err := remoteRepository.GetPlanet(IDResource)
		if err != nil {
			fmt.Println("There was an error getting the planet "+ IDResource +" remotely")
		}

		var csvRepository internal.PlanetRepo
		csvRepository = csv.NewCsvRepository()
		err = csvRepository.AddPlanet(planet)
		if err != nil {
			fmt.Println("There was an error saving the planet "+ IDResource + "to disk")
		}
	}
}

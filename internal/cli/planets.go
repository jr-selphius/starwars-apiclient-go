package cli

import (
	"fmt"

	. "github.com/jr-selphius/starwars-apiclient-go/internal"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

const idFlag string = "id"

func InitPlanetsCmd(remoteRepository PlanetRepo, csvRepository PlanetRepo) *cobra.Command {
	planetsCmd := &cobra.Command{
		Use:   "planets",
		Short: "Save planet to csv file",
		Run:   runPlanetsFn(remoteRepository, csvRepository),
	}

	planetsCmd.Flags().StringP(idFlag, "i", "", "id of the planet")

	return planetsCmd
}

const APIEndpoint string = "https://swapi.co/api/"
const APIResource string = "planets/"

func runPlanetsFn(remoteRepository PlanetRepo, csvRepository PlanetRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		IDResource, _ := cmd.Flags().GetString(idFlag)

		if IDResource == "" {
			IDResource = "1"
		}

		planet, err := remoteRepository.GetPlanet(IDResource)
		if err != nil {
			fmt.Println("There was an error getting the planet " + IDResource + " remotely")
		}

		err = csvRepository.AddPlanet(planet)
		if err != nil {
			fmt.Println("There was an error saving the planet " + IDResource + "to disk")
		}
	}
}

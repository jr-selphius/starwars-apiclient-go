package cli

import (
	"encoding/csv"
	"os"
	"fmt"

	"github.com/jr-selphius/starwars-apiclient-go/internal"
	"github.com/jr-selphius/starwars-apiclient-go/internal/storage/remote"
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
		//repo = csv.NewRepository()
		//resp, _ := http.Get(APIEndpoint + APIResource + IDResource)
		//body, _ := ioutil.ReadAll(resp.Body)
		//json.Unmarshal(body, &planet)

		f, _ := os.Create("planet." + IDResource + ".csv")
		defer f.Close()

		csvWriter := csv.NewWriter(f)

		csvWriter.Write(planet.ToArray())
		csvWriter.Flush()
	}
}

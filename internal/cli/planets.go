package cli

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

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

		var planet Planet

		resp, _ := http.Get(APIEndpoint + APIResource + IDResource)
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &planet)

		f, _ := os.Create("planet." + IDResource + ".csv")
		defer f.Close()

		csvWriter := csv.NewWriter(f)

		csvWriter.Write(planet.ToArray())
		csvWriter.Flush()
	}
}

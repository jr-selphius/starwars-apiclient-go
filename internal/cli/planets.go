package cli

import(
	"github.com/spf13/cobra"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type CobraFn func(cmd *cobra.Command, args []string)

func InitPlanetsCmd() *cobra.Command {
	planetsCmd := &cobra.Command {
		Use: "planets",
		Short: "Save planet to csv file",
		Run: runPlanetsFn(),
	}
	return planetsCmd
}

const APIEndpoint string = "https://swapi.co/api/"
const APIResource string = "planets/"

func runPlanetsFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {

		var planet Planet

		IDResource := "1"

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

type Planet struct {
	Name           string `json:"name"`
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Diameter       string `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	SurfaceWater   string `json:"surface_water"`
	Population     string `json:"population"`
	Created        string `json:"created"`
	Edited         string `json:"edited"`
	URL            string `json:"url"`
}

func (p *Planet) ToArray() []string {
	return []string{p.Name, p.RotationPeriod, p.OrbitalPeriod, p.Diameter, p.Climate, p.Gravity, p.SurfaceWater, p.Population, p.Created, p.Edited, p.URL}
}

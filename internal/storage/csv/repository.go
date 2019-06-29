package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jr-selphius/starwars-apiclient-go/internal"
)

type planetRepo struct {
}

func NewCsvRepository() internal.PlanetRepo {
	return &planetRepo{}
}

func (p *planetRepo) AddPlanet(planet *internal.Planet) error {
	f, err := os.Create("planet." + planet.ID + ".csv")
	defer f.Close()

	csvWriter := csv.NewWriter(f)

	csvWriter.Write(planet.ToArray())
	csvWriter.Flush()

	return err
}

func (p *planetRepo) GetPlanet(id string) (planet *internal.Planet, err error) {
	fmt.Printf("Not implemented")
	panic(1)
}

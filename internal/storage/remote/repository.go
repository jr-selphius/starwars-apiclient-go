package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jr-selphius/starwars-apiclient-go/internal/cli"
)

const (
	APIEndpoint string = "https://swapi.co/api/"
	APIResource string = "planets/"
)

type planetRepo struct {
	url string
}

func NewRemoteRepository() cli.PlanetRepo {
	return &planetRepo{url: APIEndpoint}
}

func (p *planetRepo) GetPlanet(id string) (planet *cli.Planet, err error) {

	resp, err := http.Get(APIEndpoint + APIResource + id)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &planet)
	if err != nil {
		return nil, err
	}

	return
}

func (p *planetRepo) AddPlanet(planet *cli.Planet) {
	fmt.Printf("Not implemented")
	panic(1)
}

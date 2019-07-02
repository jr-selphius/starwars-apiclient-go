package fetching

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jr-selphius/starwars-apiclient-go/internal"
)

const (
	APIEndpoint string = "https://swapi.co/api/"
	APIResource string = "planets/"
)

type Service interface {
	GetPlanet(id string) (*internal.Planet, error)
}

type service struct {
	url string
}

func NewService() Service {
	return &service{}
}

func (s *service) GetPlanet(id string) (planet *internal.Planet, err error) {

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

	planet.ID = id

	return
}

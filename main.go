package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const APIEndpoint string = "https://swapi.co/api/"
const APIResource string = "planets/"

func main() {

	var planet Planet

	IDResource := "1"

	resp, _ := http.Get(APIEndpoint + APIResource + IDResource)
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &planet)
	fmt.Println(planet.ToCsv())
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

func (p *Planet) ToCsv() string {
	return p.Name + "," + p.RotationPeriod + "," + p.OrbitalPeriod + "," + p.Diameter + "," + p.Climate + "," + p.Gravity + "," + p.SurfaceWater + "," + p.SurfaceWater + "," + p.Population + "," + p.Created + "," + p.Edited + "," + p.URL
}

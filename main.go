package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"fmt"
)

const APIEndpoint string = "https://swapi.co/api/"
const APIResource string = "planets/"

func main() {

	var planet Planet

	resp, _ := http.Get(APIEndpoint + APIResource + "1")
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &planet)
	fmt.Println(planet)
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

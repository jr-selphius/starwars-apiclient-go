package internal

type Planet struct {
	ID             string
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

type PlanetRepo interface {
	GetPlanet(id string) (*Planet, error)
	AddPlanet(planet *Planet) error
}

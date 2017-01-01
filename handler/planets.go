package handler

import (
  "io/ioutil"
  "fmt"
  "encoding/json"

	"gopkg.in/gin-gonic/gin.v1"
)

type Planet struct {
	Name           string `json:"name"`
	RotationPeriod int    `json:"rotation_period"`
	OrbitalPeriod  int    `json:"orbital_period"`
	Diameter       int    `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	Terrain        string `json:"terrain"`
	SurfaceWater   int    `json:"surface_water"`
	Population     int    `json:"population"`
}

func Planets(c *gin.Context) {
  data, _ := ioutil.ReadFile("./planets.json")
  var animals []Planet
  err := json.Unmarshal(data, &animals)
  if err != nil {
    fmt.Println("error:", err)
  }
  c.JSON(200, animals)
}

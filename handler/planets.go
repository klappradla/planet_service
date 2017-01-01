package handler

import (
  "encoding/json"

	"gopkg.in/gin-gonic/gin.v1"

  "github.com/klappradla/planet_service/data"
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
  var planets []Planet
  err := json.Unmarshal(data.Planets(), &planets)
  if err != nil {
    c.AbortWithError(500, err)
    return
  }
  c.JSON(200, planets)
}

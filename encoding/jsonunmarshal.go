package main

import (
	"encoding/json"
	"fmt"
)

type EngineConfig struct {
	Count  int    `json:"count"`
	Engine string `json:"engine"`
}

type Rocket struct {
	Name         string       `json:"name"`
	Manufacturer string       `json:"manufacturer"`
	Engines      EngineConfig `json:"engineconfig"`
	Thrust       float32      `json:"thrust"`
	Length       int          `json:"length,omitempty"`
	Reusable     bool         `json:"resuable"`
	Cost         float32      `json:"-"`
}

func main() {
	var data = `[
        {
                "name": "Saturn V",
                "manufacturer": "Nasa",
                "engineconfig": {
                        "count": 5,
                        "engine": "F-1"
                },
                "thrust": 1500000.00,
                "length": 42,
                "resuable": false
        },
        {
                "name": "Falcon Heavy",
                "manufacturer": "Space-X",
                "engineconfig": {
                        "count": 3,
                        "engine": "Merlin 1-D"
                },
                "thrust": 3400000.00,
                "length": 70,
                "resuable": true
        },
        {
                "name": "Electron",
                "manufacturer": "Rocket Lab",
                "engineconfig": {
                        "count": 2,
                        "engine": "Rutherford"
                },
                "thrust": 50400.00,
                "resuable": true
        }
]`

	var rockets []Rocket

	if err := json.Unmarshal([]byte(data), &rockets); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", rockets)
	fmt.Printf("%d\n", rockets[2].Engines.Count)
	fmt.Printf("%.2f\n", rockets[2].Thrust)
}

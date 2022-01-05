package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type EngineConfig struct {
	Count  int    `xml:"count"`
	Engine string `xml:"engine"`
}

type Rocket struct {
	XMLName      xml.Name     `xml:"rocket"`
	Name         string       `xml:"name,attr"`
	Manufacturer string       `xml:"manufacturer"`
	Engines      EngineConfig `xml:"engineconfig"`
	Thrust       float32      `xml:"thrust"`
	Length       int          `xml:"length,omitempty"`
	Reusable     bool         `xml:"resuable"`
	Cost         float32      `xml:"-"`
}

type Rockets struct {
	XMLName xml.Name `xml:"rockets"`
	Rockets []Rocket `xml:"rocket"`
}

const HTTPBIN_URL = "https://httpbin.org/post?course=golang&author=jeremycook"

func PostRequest(data string) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var jsonStr = []byte(data)
	req, _ := http.NewRequest("POST", HTTPBIN_URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("Accept", "*/*")
	resp, err := client.Do(req)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		// timeout error
		panic("TIMEOUT!!")
	} else if err != nil {
		// error, but not a timeout
		panic("ERROR!!")
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Printf(string(body))
}

func main() {
	saturnV := Rocket{
		Name:         "Saturn V",
		Manufacturer: "Nasa",
		Engines:      EngineConfig{5, "F-1"},
		Thrust:       1500000,
		Length:       42,
		Reusable:     false,
		Cost:         500.0,
	}

	falconHeavy := Rocket{
		Name:         "Falcon Heavy",
		Manufacturer: "Space-X",
		Engines:      EngineConfig{3, "Merlin 1-D"},
		Thrust:       3400000,
		Length:       70,
		Reusable:     true,
		Cost:         135.0,
	}

	electron := Rocket{
		Name:         "Electron",
		Manufacturer: "Rocket Lab",
		Engines:      EngineConfig{2, "Rutherford"},
		Thrust:       50400,
		Reusable:     true,
		Cost:         7.5,
	}

	rockets := Rockets{
		Rockets: []Rocket{saturnV, falconHeavy, electron},
	}

	//xml, err := json.Marshal(rockets)
	xml, err := xml.MarshalIndent(rockets, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n\n", string(xml))
	fmt.Printf("=======================\n\n")

	PostRequest(string(xml))
}

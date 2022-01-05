package main

import (
	"encoding/xml"
	"fmt"
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

func main() {
	var data = `<rockets>
	<rocket name="Saturn V">
			<manufacturer>Nasa</manufacturer>
			<engineconfig>
					<count>5</count>
					<engine>F-1</engine>
			</engineconfig>
			<thrust>1.5e+06</thrust>
			<length>42</length>
			<resuable>false</resuable>
	</rocket>
	<rocket name="Falcon Heavy">
			<manufacturer>Space-X</manufacturer>
			<engineconfig>
					<count>3</count>
					<engine>Merlin 1-D</engine>
			</engineconfig>
			<thrust>3.4e+06</thrust>
			<length>70</length>
			<resuable>true</resuable>
	</rocket>
	<rocket name="Electron">
			<manufacturer>Rocket Lab</manufacturer>
			<engineconfig>
					<count>2</count>
					<engine>Rutherford</engine>
			</engineconfig>
			<thrust>50400</thrust>
			<resuable>true</resuable>
	</rocket>
</rockets>`

	var rockets Rockets

	if err := xml.Unmarshal([]byte(data), &rockets); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", rockets)
	fmt.Printf("%s\n", rockets.Rockets[2].Name)
	fmt.Printf("%d\n", rockets.Rockets[2].Engines.Count)
	fmt.Printf("%.2f\n", rockets.Rockets[2].Thrust)
}

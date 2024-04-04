package main

import (
	"encoding/xml"
	"fmt"
)

// Plant is mapped to XML
// Like JSON field tags contain directives for encoding and decoding
// XMLName field is a special feature which dictates the xml element name
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// MarshalIndent is like Marshal but adds line breaks and indentation
	out, _ := xml.MarshalIndent(coffee, "", "  ")
	fmt.Println(string(out))

	// We can add the XML header manually
	fmt.Println(xml.Header + string(out))

	// Unmarshal to a new Plant struct
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// Nesting elements
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}

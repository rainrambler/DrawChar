package main

import (
	"encoding/xml"
	"fmt"
)

// Dictionary was generated 2023-04-24 03:03:54 by https://xml-to-go.github.io/
type Dictionary struct {
	XMLName   xml.Name `xml:"dictionary"`
	Text      string   `xml:",chardata"`
	Name      string   `xml:"name,attr"`
	Character []struct {
		Text    string `xml:",chardata"`
		Utf8    string `xml:"utf8"`
		Strokes struct {
			Text   string `xml:",chardata"`
			Stroke []struct {
				Text  string `xml:",chardata"`
				Point []struct {
					Text string `xml:",chardata"`
					X    string `xml:"x,attr"`
					Y    string `xml:"y,attr"`
				} `xml:"point"`
			} `xml:"stroke"`
		} `xml:"strokes"`
	} `xml:"character"`
}

// Dictionary was generated 2023-04-24 03:03:54 by https://xml-to-go.github.io/
type Dictionary2 struct {
	XMLName   xml.Name     `xml:"dictionary"`
	Text      string       `xml:",chardata"`
	Name      string       `xml:"name,attr"`
	Character []Character2 `xml:"character"`
}

type Character2 struct {
	Text    string   `xml:",chardata"`
	Utf8    string   `xml:"utf8"`
	Strokes Strokes2 `xml:"strokes"`
}

type Strokes2 struct {
	Text   string    `xml:",chardata"`
	Stroke []Stroke2 `xml:"stroke"`
}

type Stroke2 struct {
	Text  string   `xml:",chardata"`
	Point []Point2 `xml:"point"`
}

type Point2 struct {
	Text string `xml:",chardata"`
	X    string `xml:"x,attr"`
	Y    string `xml:"y,attr"`
}

func readDictXml(filename string) {
	bs, err := ReadBinFile(filename)
	if err != nil {
		fmt.Printf("[WARN]Cannot read file: %s!\n", filename)
		return
	}

	dict := &Dictionary2{}
	xml.Unmarshal(bs, dict)
	fmt.Println(dict)
}

func drawDemo() {
	var oc Character2
	oc.Text = "11"
	oc.Utf8 = "22"

	drawChar(&oc)
}

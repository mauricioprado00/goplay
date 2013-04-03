package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Action struct {
	Method string   `xml:"method,attr"`
	Params []string `xml:",any"`
}

type Block struct {
	XMLName  xml.Name `xml:"block"`
	Blocks   []Block  `xml:"block"`
	Type     string   `xml:"type,attr"`
	As       string   `xml:"as,attr"`
	Template string   `xml:"template,attr"`
	Actions  []Action `xml:"action"`
}

type Reference struct {
	Blocks  []Block  `xml:"block"`
	Actions []Action `xml:"action"`
}

type Layout struct {
	Output     string      `xml:"Config>Output"`
	Groups     []string    `xml:"Config>Group>Value"`
	References []Reference `xml:"reference"`
}

func (this *Layout) showReferences() {
	for _, reference := range this.References {
		for _, block := range reference.Blocks {
			fmt.Println(block.XMLName.Local + " " + block.Type + " " + block.As + " " + block.Template)
		}
		for _, action := range reference.Actions {
			fmt.Println("action: " + action.Method)
			for _, param := range action.Params {
				fmt.Println("   param:" + param)
			}
		}
	}
}

func main() {
	var filename string = "layout.xml"
	filebyte, _ := ioutil.ReadFile(filename)

	v := Layout{}

	err := xml.Unmarshal(filebyte, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(string(filebyte))
	fmt.Printf("Name: %q\n", v.Output)
	fmt.Printf("Groups: %v\n", v.Groups)
	v.showReferences()
	//fmt.Printf("%d", len(v.References))
	//fmt.Printf("%v", v.References)
}

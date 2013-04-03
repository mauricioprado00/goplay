package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
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
	XMLName xml.Name `xml:"reference"`
	Name    string   `xml:"name,attr"`
	Blocks  []Block  `xml:"block"`
	Actions []Action `xml:"action"`
}

type Layout struct {
	Output     string      `xml:"Config>Output"`
	Groups     []string    `xml:"Config>Group>Value"`
	References []Reference `xml:"reference"`
}

type LayoutElement interface {
	GetReferences() []Reference
	GetBlocks() []Block
	GetActions() []Action
}

/** Block methods */
func (this *Block) GetReferences() []Reference {
	return nil
}

func (this *Block) GetBlocks() []Block {
	return this.Blocks
}

func (this *Block) GetActions() []Action {
	return this.Actions
}

/** Reference methods */
func (this *Reference) GetReferences() []Reference {
	return nil
}

func (this *Reference) GetBlocks() []Block {
	return this.Blocks
}

func (this *Reference) GetActions() []Action {
	return this.Actions
}

/** Layout methods */
func (this *Layout) GetReferences() []Reference {
	return this.References
}

func (this *Layout) GetBlocks() []Block {
	return nil
}

func (this *Layout) GetActions() []Action {
	return nil
}

func dumpLayoutElement(layoutElement LayoutElement, indent int) {
	indentation := strings.Repeat("  ", indent)
	for _, reference := range layoutElement.GetReferences() {
		fmt.Println(indentation + reference.XMLName.Local + " " + reference.Name)
		//fmt.Println(indentation + "Reference: ")
		dumpLayoutElement(&reference, indent+1)
	}
	for _, block := range layoutElement.GetBlocks() {
		fmt.Println(indentation + block.XMLName.Local + " " + block.Type + " " + block.As + " " + block.Template)
		dumpLayoutElement(&block, indent+1)
	}
	for _, action := range layoutElement.GetActions() {
		fmt.Println(indentation + "action: " + action.Method)
		for _, param := range action.Params {
			fmt.Println(indentation + "   param:" + param)
		}
	}
}

func (this *Block) dump() {
	dumpLayoutElement(this, 0)
}
func (this *Reference) dump() {
	dumpLayoutElement(this, 0)
}
func (this *Layout) dump() {
	dumpLayoutElement(this, 0)
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
	v.dump()
	//fmt.Printf("%d", len(v.References))
	//fmt.Printf("%v", v.References)
}

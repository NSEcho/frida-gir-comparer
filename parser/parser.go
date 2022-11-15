package parser

import (
	"encoding/xml"
	"os"
)

type Repository struct {
	XMLName   xml.Name  `xml:"repository"`
	Package   Package   `xml:"package"`
	Include   []Include `xml:"include"`
	Namespace Namespace `xml:"namespace"`
}

type Package struct {
	Name string `xml:"name,attr"`
}

type Include struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
}

type CInclude struct {
	Name string `xml:"name,attr"`
}

type Namespace struct {
	Name         string        `xml:"name,attr"`
	Version      string        `xml:"version,attr"`
	Enumerations []Enumeration `xml:"enumeration"`
	Functions    []Function    `xml:"function"`
	Classes      []Class       `xml:"class"`
}

type Class struct {
	Name    string   `xml:"name,attr"`
	CType   string   `xml:"type,attr"`
	Methods []Method `xml:"method"`
	Signals []Signal `xml:"signal"`
}

type Method struct {
	Name        string      `xml:"name,attr"`
	Identifier  string      `xml:"identifier,attr"`
	ReturnValue ReturnValue `xml:"return-value"`
}

type Signal struct {
	Name        string      `xml:"name,attr"`
	ReturnValue ReturnValue `xml:"return-value"`
}

type ReturnValue struct {
	Type struct {
		Name  string `xml:"name,attr"`
		CType string `xml:"type,attr"`
	} `xml:"type"`
}

type Parameters struct {
	Params []Parameter `xml:"parameter"`
}

type Parameter struct {
	Name string `xml:"name,attr"`
	Type struct {
		Name  string `xml:"name,attr"`
		CType string `xml:"type,attr"`
	} `xml:"type"`
}

type Function struct {
	Name        string      `xml:"name,attr"`
	Identifier  string      `xml:"identifier,attr"`
	ReturnValue ReturnValue `xml:"return-value"`
	Parameters  Parameters  `xml:"parameters"`
}

type Enumeration struct {
	Name     string   `xml:"name,attr"`
	CType    string   `xml:"type,attr"`
	GlibType string   `xml:"type-name,attr"`
	Members  []Member `xml:"member"`
}

type Member struct {
	Name       string `xml:"name,attr"`
	Identifier string `xml:"identifier,attr"`
	Value      string `xml:"value,attr"`
}

type Parser struct {
	rep *Repository
}

func NewParser(girPath string) (*Parser, error) {
	f, err := os.Open(girPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var repo *Repository
	if err := xml.NewDecoder(f).Decode(&repo); err != nil {
		return nil, err
	}

	return &Parser{repo}, nil
}

func (p *Parser) Classes() []Class {
	return p.rep.Namespace.Classes
}

func (p *Parser) Functions() []Function {
	return p.rep.Namespace.Functions
}

func (p *Parser) Enumerations() []Enumeration {
	return p.rep.Namespace.Enumerations
}

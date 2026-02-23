package parser

import (
	"encoding/xml"
	"github.com/nsecho/fgcomparer/helper"
	"github.com/nsecho/fgcomparer/types"
	"os"
	"strings"
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
	Parameters  Parameters  `xml:"parameters"`
	ClsName     string
}

func (m *Method) String() string {
	var out strings.Builder
	out.WriteString(helper.ConvertToCamelCase(m.Name) + "(")

	// write arguments
	for i, param := range m.Parameters.Params {
		if param.Type.CType != "GCancellable*" {
			out.WriteString(param.Name + " " + convertToGoType(param.Type.CType))
			if i < len(m.Parameters.Params)-1 {
				out.WriteString(", ")
			}
		}
	}
	out.WriteString(") ")

	hasReturn := false
	isFridaReturn := false

	// write return type
	if m.ReturnValue.Type.CType != "void" {
		hasReturn = true
		out.WriteString(convertToGoType(m.ReturnValue.Type.CType))
		if strings.Contains(m.ReturnValue.Type.CType, "Frida") {
			isFridaReturn = true
		}
	}
	out.WriteString(" {\n")

	newParamNames := map[string]string{}

	for _, param := range m.Parameters.Params {
		if param.Type.CType != "GCancellable*" {
			typeFn, ok := types.FnParams[stripType(param.Type.CType)]
			if ok {
				paramDesc, paramNewName := typeFn(param.Name)
				newParamNames[param.Name] = paramNewName
				out.WriteString("\t" + paramDesc + "\n\n")
			}
		}
	}

	if hasReturn {
		out.WriteString("\trt := ")
	} else {
		out.WriteString("\t")
	}

	// call function
	out.WriteString("C." + m.Identifier + "(")
	out.WriteString(m.ClassName() + "." + m.ClassName())

	if len(m.Parameters.Params) > 0 {
		out.WriteString(", ")
	}

	// pass params to C function along with new parameters
	for i, param := range m.Parameters.Params {
		if param.Name == "cancellable" {
			out.WriteString("nil")
		} else {
			if pname, ok := newParamNames[param.Name]; ok {
				out.WriteString(pname)
			} else {
				if strings.Contains(param.Type.CType, "Frida") {
					out.WriteString(param.Name + "." + strings.ToLower(string(stripType(param.Type.CType)[0])))
				} else {
					out.WriteString(param.Name)
				}
			}
		}
		if i < len(m.Parameters.Params)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")\n")
	if m.ReturnValue.Type.CType != "void" {
		tp := convertToGoType(m.ReturnValue.Type.CType)
		fn, ok := types.RetTypesFns[tp]
		if ok {
			out.WriteString("\treturn " + fn() + "\n")
		} else {
			if isFridaReturn {
				out.WriteString("\t return &" + strings.ReplaceAll(stripType(m.ReturnValue.Type.CType), "*", "") + "{rt}\n")
			} else {
				out.WriteString("\treturn rt\n")
			}
		}
		// out.WriteString(types.RetTypesFns[convertToGoType(m.ReturnValue.Type.CType)]() + "\n")
	}
	out.WriteString("}\n")
	return out.String()
}

func stripType(ctype string) string {
	s := strings.ReplaceAll(ctype, "const", "")
	s = strings.ReplaceAll(s, "Frida", "")
	s = strings.TrimSpace(s)
	return s
}

func convertToGoType(ctype string) string {
	s := stripType(ctype)
	goType, ok := types.GoTypes[s]
	if ok {
		return goType
	} else {
		if strings.HasSuffix(s, "*") {
			return "*" + s[:len(s)-1]
		}
		return s
	}
}

func (m *Method) ClassName() string {
	return strings.ToLower(string(m.ClsName[0]))
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

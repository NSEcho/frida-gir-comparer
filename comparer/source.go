package comparer

import (
	"github.com/nsecho/fgcomparer/parser"
	"strings"
)

type Source struct {
	NewClasses          []parser.Class
	RemovedClasses      []string
	NewEnumerations     []parser.Enumeration
	RemovedEnumerations []string
}

func (s *Source) String() string {
	var out strings.Builder

	out.WriteString("// New classes:\n")
	for _, class := range s.NewClasses {
		out.WriteString("type " + class.Name + " struct {\n\t")
		out.WriteString(strings.ToLower(string(class.Name[0])))
		out.WriteString(" *C." + class.CType + "\n}\n\n")
		parsedMethods := make(map[string]*parser.Method)
		// out.WriteString(class.Name + "\n")
		for i, method := range class.Methods {
			parsedMethods[method.Name] = &class.Methods[i]
		}
		parsed := make(map[string]struct{})
		for method := range parsedMethods {
			if _, parsedAlready := parsed[method]; !parsedAlready {
				/*if _, ok := parsedMethods[method+"_finish"]; ok {
					if _, k := parsedMethods[method+"_sync"]; k {
						parsed[method] = struct{}{}
						parsed[method+"_finish"] = struct{}{}
						parsed[method+"_sync"] = struct{}{}
						//out.WriteString("sync " + method + "_sync" + parsedMethods[method].ReturnValue.Type.CType + "\n")
						// out.WriteString(syncMethods[method].String())
					}
				} else {*/
				parsedMethods[method].ClsName = class.Name
				out.WriteString("func (" + strings.ToLower(string(class.Name[0])) + " *" + class.Name + ") " + parsedMethods[method].String() + "\n")
				// out.WriteString(parsedMethods[method].String())
				//}
			}
		}
	}

	out.WriteString("\n\n")

	/*out.WriteString("// New enumerations:\n")
	for _, enum := range s.NewEnumerations {
		startChar := strings.ToLower(string(enum.Name[0]))
		enumString := fmt.Sprintf("func (%s %s) String() string {\n",
			startChar, enum.Name)
		enumString += "\treturn [...]string{\""
		out.WriteString(fmt.Sprintf("// Enumeration %s\n", enum.Name))
		out.WriteString(fmt.Sprintf("type %s int\n\n", enum.Name))
		out.WriteString("const (\n")
		out.WriteString("\t" + helper.ConvertToCamelCase(enum.Members[0].Identifier) + " " + enum.Name + " = iota\n")
		enumString += enum.Members[0].Name + "\",\n"
		for i := 1; i < len(enum.Members); i++ {
			out.WriteString("\t" + helper.ConvertToCamelCase(enum.Members[i].Identifier) + "\n")
			enumString += "\t\t\"" + enum.Members[i].Name
			if i != len(enum.Members)-1 {
				enumString += "\",\n"
			} else {
				enumString += "\"}[" + startChar + "]\n}\n"
			}
		}
		out.WriteString(")\n\n")

		out.WriteString(enumString)
	}*/

	out.WriteString("// Removed enumerations:\n")
	for _, enum := range s.RemovedEnumerations {
		out.WriteString("* " + enum + "\n")
	}
	out.WriteString("\n\n")

	return out.String()
}

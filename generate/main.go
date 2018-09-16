package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TypeProperty struct {
	Name   string
	Type   string
	Setter bool
}

func (p TypeProperty) GetterName() string {
	return "Get" + strings.Title(p.Name)
}

func (p TypeProperty) SetterName() string {
	return "Set" + strings.Title(p.Name)
}

type FunctionArgument struct {
	Name string
	Type string
}

type TypeFunction struct {
	Name       string
	Arguments  []FunctionArgument
	ReturnType string
}

func (f TypeFunction) GoName() string {
	return strings.Title(f.Name)
}

type TypeStructure struct {
	Type       string
	ImportJS   bool
	Interface  bool
	Inherits   []string
	Implements []string
	Properties []TypeProperty
	Functions  []TypeFunction
}

func (t TypeStructure) Filename() string {
	return fmt.Sprintf("%s.go", strings.ToLower(t.Type))
}

func (t TypeStructure) ShortName() string {
	return strings.ToLower(string(t.Type[0]))
}

func (t TypeStructure) writeArguments(out *bytes.Buffer, args []FunctionArgument) {
	for i, arg := range args {
		if i > 0 {
			out.WriteString(",")
		}
		out.WriteString(fmt.Sprintf("%s %s", arg.Name, arg.Type))
	}
}

func (t TypeStructure) writeFunctions(out *bytes.Buffer, funcs []TypeFunction) error {
	for _, f := range funcs {
		out.WriteString(fmt.Sprintf("func (%s *%s) %s(", t.ShortName(), t.Type, f.GoName()))
		for i, arg := range f.Arguments {
			if i > 0 {
				out.WriteString(",")
			}
			out.WriteString(fmt.Sprintf("%s %s", arg.Name, arg.Type))
		}
		out.WriteString(fmt.Sprintf(") %s {\r\n", f.ReturnType))

		if f.ReturnType != "" {
			out.WriteString("val := ")
		}
		out.WriteString(fmt.Sprintf("%s.Call(\"%s\"", t.ShortName(), f.Name))
		for _, arg := range f.Arguments {
			out.WriteString(",")
			out.WriteString(fmt.Sprintf("ToValue(%s)", arg.Name))
		}
		out.WriteString(")\r\n")

		switch f.ReturnType {
		case "string":
			out.WriteString("return val.String()\r\n")
		case "*Element":
			out.WriteString("return &Element{ Value: val }\r\n")
		case "[]*Element":
			out.WriteString("elms := make([]*Element, 0)\r\n")
			out.WriteString("for i := 0; i < val.Length(); i += 1 {\r\n")
			out.WriteString("\telms = append(elms, &Element{Value: val.Index(i)})\r\n")
			out.WriteString("}\r\n")
			out.WriteString("return elms\r\n")
		case "":
			// skip
		default:
			return fmt.Errorf("Unknown function return type %s", f.ReturnType)
		}

		out.WriteString("}\r\n")
	}
	return nil
}

func (t TypeStructure) writeProperties(out *bytes.Buffer, props []TypeProperty) error {
	for _, p := range props {
		out.WriteString(fmt.Sprintf("func (%s *%s) %s() %s {\r\n", t.ShortName(), t.Type, p.GetterName(), p.Type))
		out.WriteString(fmt.Sprintf("\tval := %s.Get(\"%s\")\r\n", t.ShortName(), p.Name))
		switch p.Type {
		case "string":
			out.WriteString("return val.String()\r\n")
		case "*Element":
			out.WriteString("return &Element{ Value: val }\r\n")
		case "[]*Element":
			out.WriteString("elms := make([]*Element, 0)\r\n")
			out.WriteString("for i := 0; i < val.Length(); i += 1 {\r\n")
			out.WriteString("\telms = append(elms, &Element{Value: val.Index(i)})\r\n")
			out.WriteString("}\r\n")
			out.WriteString("return elms\r\n")
		case "":
			// skip
		default:
			return fmt.Errorf("Unknown property return type %s", p.Type)
		}
		out.WriteString(fmt.Sprintf("}\r\n"))

		if p.Setter {
			out.WriteString(fmt.Sprintf("func (%s *%s) %s(v %s){\r\n", t.ShortName(), t.Type, p.SetterName(), p.Type))
			out.WriteString(fmt.Sprintf("\t%s.Set(\"%s\", v)\r\n", t.ShortName(), p.Name))
			out.WriteString("}\r\n")
		}
	}
	return nil
}

func (t TypeStructure) generateInterface(out *bytes.Buffer, types map[string]TypeStructure) error {
	out.WriteString(fmt.Sprintf("type %s interface {\r\n", t.Type))

	for _, f := range t.Functions {
		out.WriteString(fmt.Sprintf("\t%s(", f.GoName()))
		t.writeArguments(out, f.Arguments)
		out.WriteString(fmt.Sprintf(") %s\r\n", f.ReturnType))
	}

	out.WriteString("}\r\n")
	return nil
}

func (t TypeStructure) generateStruct(out *bytes.Buffer, types map[string]TypeStructure) (err error) {
	out.WriteString(fmt.Sprintf("type %s struct {\r\n", t.Type))

	out.WriteString("\tjs.Value\r\n")
	for _, name := range t.Inherits {
		out.WriteString(fmt.Sprintf("\t%s\r\n", name))
	}

	out.WriteString("}\r\n")

	// Properties
	err = t.writeProperties(out, t.Properties)
	if err != nil {
		return
	}

	// Functions
	err = t.writeFunctions(out, t.Functions)
	if err != nil {
		return
	}

	// Interfaces
	for _, it := range t.Implements {
		i, ok := types[it]
		if !ok {
			return fmt.Errorf("%q cannot implement unknown type %q", t.Type, it)
		}

		// Functions
		err = t.writeFunctions(out, i.Functions)
		if err != nil {
			return
		}

		// Properties
		err = t.writeProperties(out, i.Properties)
		if err != nil {
			return
		}
	}

	return
}

func (t TypeStructure) Generate(types map[string]TypeStructure) ([]byte, error) {
	out := bytes.NewBuffer(make([]byte, 0))
	out.WriteString("// DO NOT EDIT - generated file\r\n")
	out.WriteString("package dom\r\n")

	if t.Interface == false || t.ImportJS == true {
		out.WriteString("import \"syscall/js\"\r\n")
	}

	var err error
	if t.Interface {
		err = t.generateInterface(out, types)
	} else {
		err = t.generateStruct(out, types)
	}
	if err != nil {
		return nil, err
	}

	return format.Source(out.Bytes())
}

func loadFile(fname string) (t TypeStructure, err error) {
	fp, err := os.Open(fname)
	if err != nil {
		return
	}
	defer fp.Close()

	val, err := ioutil.ReadAll(fp)
	if err != nil {
		return
	}

	err = json.Unmarshal(val, &t)
	return
}

func processType(t TypeStructure, types map[string]TypeStructure) error {
	contents, err := t.Generate(types)
	if err != nil {
		return err
	}

	out, err := os.OpenFile(t.Filename(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = out.Write(contents)
	return err
}

func main() {
	matches, err := filepath.Glob("generate/*.json")
	if err != nil {
		log.Panic(err)
	}

	types := make(map[string]TypeStructure)
	for _, fname := range matches {
		typeStructure, err := loadFile(fname)
		if err != nil {
			log.Panic(err)
		}
		types[typeStructure.Type] = typeStructure
	}

	for _, t := range types {
		log.Printf("Generating %s\r\n", t.Type)
		err = processType(t, types)
		if err != nil {
			log.Panic(err)
		}
	}
}

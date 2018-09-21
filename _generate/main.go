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
	Name     string
	Type     string
	ReadOnly bool
}

func (p TypeProperty) GetterName() string {
	return "Get" + strings.Title(p.Name)
}

func (p TypeProperty) SetterName() string {
	return "Set" + strings.Title(p.Name)
}

type FunctionArgument struct {
	Name    string
	Type    string
	Varidic bool
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
	GlobalAPI  string
	ImportJS   bool
	Interface  bool
	Inherits   []string
	Implements []string
	Properties []TypeProperty
	Functions  []TypeFunction
	String     []string
}

func (t TypeStructure) Filename() string {
	return fmt.Sprintf("%s.go", strings.ToLower(t.Type))
}

func (t TypeStructure) APIFilename() string {
	return fmt.Sprintf("%s/%s.go", t.PackageName(), strings.ToLower(t.Type))
}

func (t TypeStructure) PackageName() string {
	return strings.ToLower(t.GlobalAPI)
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

func (t TypeStructure) writeReturnValue(out *bytes.Buffer, rt string) error {
	switch rt {
	case "string":
		out.WriteString("return val.String()\r\n")
	case "Value":
		out.WriteString("return val")
	case "NodeIFace":
		out.WriteString("return NewNode(val.JSValue())\r\n")
	case "*ShadowRoot":
		out.WriteString("return NewShadowRoot(val.JSValue())\r\n")
	case "*Promise":
		out.WriteString("return NewPromise(val.JSValue())\r\n")
	case "*Element":
		out.WriteString("return NewElement(val.JSValue())\r\n")
	case "[]*Element":
		out.WriteString("elms := make([]*Element, 0)\r\n")
		out.WriteString("for i := 0; i < val.Length(); i += 1 {\r\n")
		out.WriteString("\telms = append(elms, NewElement(val.Index(i)))\r\n")
		out.WriteString("}\r\n")
		out.WriteString("return elms\r\n")
	case "":
		out.WriteString("return val")
	default:
		return fmt.Errorf("Unknown function return type %s", rt)
	}

	return nil
}

func (t TypeStructure) writeFunctions(out *bytes.Buffer, funcs []TypeFunction) error {
	for _, f := range funcs {
		out.WriteString(fmt.Sprintf("func (%s *%s) %s(", t.ShortName(), t.Type, f.GoName()))
		for i, arg := range f.Arguments {
			if i > 0 {
				out.WriteString(",")
			}
			at := arg.Type
			if arg.Varidic {
				at = "..." + at
			}
			out.WriteString(fmt.Sprintf("%s %s", arg.Name, at))
		}
		rt := f.ReturnType
		if rt == "" {
			rt = "Value"
		}
		out.WriteString(fmt.Sprintf(") %s {\r\n", rt))

		for _, arg := range f.Arguments {
			if !arg.Varidic {
				continue
			}

			out.WriteString(fmt.Sprintf("%sVaridic := make([]interface{}, 0)\r\n", arg.Name))
			out.WriteString(fmt.Sprintf("for _, a := range %s {\r\n", arg.Name))
			out.WriteString(fmt.Sprintf("%sVaridic = append(%sVaridic, ToJSValue(a))\r\n", arg.Name, arg.Name))
			out.WriteString("}\r\n")
		}

		out.WriteString(fmt.Sprintf("val := Value{ Value: %s.Call(\"%s\"", t.ShortName(), f.Name))
		for _, arg := range f.Arguments {
			out.WriteString(",")
			if arg.Varidic {
				out.WriteString(fmt.Sprintf("%sVaridic...", arg.Name))
			} else {
				out.WriteString(fmt.Sprintf("ToJSValue(%s)", arg.Name))
			}
		}
		out.WriteString(")}\r\n")

		// Return value
		t.writeReturnValue(out, rt)

		out.WriteString("}\r\n")
	}
	return nil
}

func (t TypeStructure) writeAPIFunctions(out *bytes.Buffer, funcs []TypeFunction) error {
	for _, f := range funcs {
		out.WriteString(fmt.Sprintf("func %s(", f.GoName()))
		for i, arg := range f.Arguments {
			if i > 0 {
				out.WriteString(",")
			}
			rt := arg.Type
			if rt == "*Element" {
				rt = "*dom.Element"
			} else if rt == "JSValue" {
				rt = "dom.JSValue"
			} else if rt == "*Callback" {
				rt = "*dom.Callback"
			} else if rt == "*Promise" {
				rt = "*dom.Promise"
			} else if rt == "ShadowRootInit" {
				rt = "dom.ShadowRootInit"
			}

			if arg.Varidic {
				rt = "..." + rt
			}
			out.WriteString(fmt.Sprintf("%s %s", arg.Name, rt))
		}
		rt := f.ReturnType
		if rt == "" {
			rt = "Value"
		}
		if rt != "string" {
			if rt[0] == '*' {
				rt = "*dom." + strings.TrimLeft(rt, "*")
			} else if rt[0] == '[' {
				rt = "[]*dom." + strings.TrimLeft(rt, "[]*")
			} else {
				rt = "dom." + rt
			}
		}

		out.WriteString(fmt.Sprintf(") %s {\r\n", rt))
		out.WriteString(fmt.Sprintf("return %s.%s(", t.ShortName(), f.GoName()))
		for i, arg := range f.Arguments {
			if i > 0 {
				out.WriteString(",")
			}
			if arg.Varidic {
				out.WriteString(fmt.Sprintf("%s...", arg.Name))
			} else {
				out.WriteString(fmt.Sprintf("%s", arg.Name))
			}
		}
		out.WriteString(")}\r\n")
	}
	return nil
}

func (t TypeStructure) writeProperties(out *bytes.Buffer, props []TypeProperty) error {
	for _, p := range props {
		rt := p.Type
		if rt == "" {
			rt = "Value"
		}

		out.WriteString(fmt.Sprintf("func (%s *%s) %s() %s {\r\n", t.ShortName(), t.Type, p.GetterName(), rt))
		out.WriteString(fmt.Sprintf("val := Value{ Value: %s.Get(\"%s\")}\r\n", t.ShortName(), p.Name))

		// Return value
		t.writeReturnValue(out, rt)

		out.WriteString(fmt.Sprintf("}\r\n"))

		if !p.ReadOnly {
			out.WriteString(fmt.Sprintf("func (%s *%s) %s(v %s){\r\n", t.ShortName(), t.Type, p.SetterName(), rt))
			out.WriteString(fmt.Sprintf("\t%s.Set(\"%s\", v)\r\n", t.ShortName(), p.Name))
			out.WriteString("}\r\n")
		}
	}
	return nil
}

func (t TypeStructure) writeAPIProperties(out *bytes.Buffer, props []TypeProperty) error {
	for _, p := range props {
		rt := p.Type
		if rt == "" {
			rt = "Value"
		}
		if rt != "string" {
			if rt[0] == '*' {
				rt = "*dom." + strings.TrimLeft(rt, "*")
			} else if rt[0] == '[' {
				rt = "[]*dom." + strings.TrimLeft(rt, "[]*")
			} else {
				rt = "dom." + rt
			}
		}

		out.WriteString(fmt.Sprintf("func %s() %s {\r\n", p.GetterName(), rt))
		out.WriteString(fmt.Sprintf("return %s.%s()", t.ShortName(), p.GetterName()))
		out.WriteString(fmt.Sprintf("}\r\n"))

		if !p.ReadOnly {
			out.WriteString(fmt.Sprintf("func %s(v %s){\r\n", p.SetterName(), rt))
			out.WriteString(fmt.Sprintf("%s.%s(v)", t.ShortName(), p.SetterName()))
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

	out.WriteString("\tValue\r\n")
	for _, name := range t.Inherits {
		out.WriteString(fmt.Sprintf("\t%s\r\n", name))
	}

	out.WriteString("}\r\n")

	// DEV: Start with upper case means it is public
	if t.Type[0] == strings.ToUpper(string(t.Type[0]))[0] {
		// Constructor
		out.WriteString(fmt.Sprintf("func New%s(v js.Value) *%s {\r\n", t.Type, t.Type))
		out.WriteString("val := Value{ Value: v }\r\n")
		out.WriteString("if val.IsNull() || val.IsUndefined() { return nil }\r\n")
		out.WriteString(fmt.Sprintf("return val.To%s()\r\n", t.Type))
		out.WriteString("}\r\n")

		// Convert to specific type helper
		out.WriteString(fmt.Sprintf("func (v Value) To%s () *%s { return &%s{ Value: v }}\r\n", t.Type, t.Type, t.Type))
	}

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

	if t.ImportJS {
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

func (t TypeStructure) GenerateAPI(types map[string]TypeStructure) ([]byte, error) {
	pkg := strings.ToLower(t.GlobalAPI)
	out := bytes.NewBuffer(make([]byte, 0))
	out.WriteString("// DO NOT EDIT - generated file\r\n")
	out.WriteString(fmt.Sprintf("package %s\r\n", pkg))

	if t.ImportJS {
		out.WriteString("import \"syscall/js\"\r\n")
	}
	out.WriteString("import dom \"github.com/brettlangdon/go-dom/v1\"\r\n")

	out.WriteString(fmt.Sprintf("var %s *dom.%s\r\n", t.ShortName(), t.Type))
	out.WriteString("func init() {\r\n")
	out.WriteString(fmt.Sprintf("%s = dom.New%s(js.Global().Get(\"%s\"))\r\n", t.ShortName(), t.Type, t.GlobalAPI))
	out.WriteString("}\r\n")

	err := t.writeAPIProperties(out, t.Properties)
	if err != nil {
		return nil, err
	}

	err = t.writeAPIFunctions(out, t.Functions)
	if err != nil {
		return nil, err
	}

	// Interfaces
	for _, it := range t.Implements {
		i, ok := types[it]
		if !ok {
			return nil, fmt.Errorf("%q cannot implement unknown type %q", t.Type, it)
		}

		// Functions
		err = t.writeAPIFunctions(out, i.Functions)
		if err != nil {
			return nil, err
		}

		// Properties
		err = t.writeAPIProperties(out, i.Properties)
		if err != nil {
			return nil, err
		}
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
	if err != nil {
		return err
	}

	if t.GlobalAPI == "" {
		return nil
	}

	contents, err = t.GenerateAPI(types)
	if err != nil {
		return err
	}
	err = os.MkdirAll(t.PackageName(), os.ModePerm)
	if err != nil {
		return err
	}
	api, err := os.OpenFile(t.APIFilename(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer api.Close()
	_, err = api.Write(contents)
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

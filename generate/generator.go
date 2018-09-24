package main

import (
	"fmt"
	"go/format"
	"os"
	"strings"
)

type Generator struct {
	fname string
	specs SpecMap
}

func NewGenerator(fname string) (*Generator, error) {
	specs, err := LoadSpecs(fname)
	return &Generator{
		fname: fname,
		specs: specs,
	}, err
}

func (g *Generator) writeInterfaceMember(m Member, b *Builder) error {
	if m.Title() == "" {
		return nil
	}

	switch m.Type {
	case "attribute":
		rt := convertIDLType(m.IDLType)
		b.WriteF("Get%s() %s", m.Title(), rt)
		if !m.IsReadOnly() {
			b.WriteF("Set%s(%s)", m.Title(), rt)
		}
	case "operation":
		b.WriteF("%s(args ...interface{}) %s", m.Title(), convertIDLType(m.Body.IDLType))
	}
	return nil
}

func (g *Generator) writeStructMember(s Spec, m Member, b *Builder) error {
	if m.Title() == "" {
		return nil
	}

	switch m.Type {
	case "attribute":
		rt := convertIDLType(m.IDLType)
		if rt == "" {
			rt = "Value"
		}

		b.WriteF("func %s Get%s() %s {", s.Receiver(), m.Title(), rt)
		if rt != "" {
			b.WriteF("val := %s.Get(\"%s\")", s.Shortname(), m.Name)
			switch rt {
			case "string":
				b.WriteF("return val.String()")
			case "[]byte":
				b.WriteF("return []byte(val.String())")
			case "bool":
				b.WriteF("return val.Bool()")
			case "float64":
				b.WriteF("return val.Float()")
			case "int":
				b.WriteF("return val.Int()")
			case "Value":
				b.WriteString("return val")
			default:
				b.WriteF("return JSValueTo%s(val.JSValue())", rt)
			}
		} else {
			b.WriteF("%s.Get(\"%s\")", s.Shortname(), m.Name)
		}

		b.WriteString("}")

		if !m.IsReadOnly() {
			b.WriteF("func %s Set%s(val %s) {", s.Receiver(), m.Title(), rt)
			b.WriteF("%s.Set(\"%s\", val)", s.Shortname(), m.Name)
			b.WriteString("}")
		}
	case "operation":
		rt := convertIDLType(m.Body.IDLType)
		b.WriteF("func %s %s(args ...interface{}) %s {", s.Receiver(), m.Title(), rt)
		if rt != "" {
			b.WriteF("val := %s.Call(\"%s\", args...)", s.Shortname(), m.Body.Name.Value)
			switch rt {
			case "string":
				b.WriteF("return val.String()")
			case "[]byte":
				b.WriteF("return []byte(val.String())")
			case "bool":
				b.WriteF("return val.Bool()")
			case "float64":
				b.WriteF("return val.Float()")
			case "int":
				b.WriteF("return val.Int()")
			case "Value":
				b.WriteF("return val")
			default:
				b.WriteF("return JSValueTo%s(val.JSValue())", rt)
			}
		} else {
			b.WriteF("%s.Call(\"%s\", args...)", s.Shortname(), m.Body.Name.Value)
		}
		b.WriteString("}")
	}
	return nil

}

func (g *Generator) generateFileHeader(spec Spec, b *Builder, importJs bool) {
	b.WriteString("// Code generated DO NOT EDIT")
	b.WriteF("// %s", spec.Filename())
	b.WriteString("package dom")
	if importJs {
		b.WriteString("import \"syscall/js\"")
	}
}

func (g *Generator) writeFile(spec Spec, b *Builder) (err error) {
	source, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	out, err := os.OpenFile(spec.Filename(), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(source)
	return err
}

func (g *Generator) generateInterface(spec Spec) error {
	b := NewBuilder()
	g.generateFileHeader(spec, b, true)

	// Interface
	b.WriteF("type %sIFace interface {", spec.Name)
	mems, err := spec.ResolveMembers(g.specs, true)
	if err != nil {
		return err
	}
	for _, m := range mems {
		err := g.writeInterfaceMember(m, b)
		if err != nil {
			return err
		}
	}
	b.WriteString("}")

	// Implementation
	b.WriteF("type %s struct {", spec.Name)
	b.WriteString("Value")

	inheritance, err := spec.ResolveInheritance(g.specs)
	if err != nil {
		return err
	}
	for _, i := range inheritance {
		b.WriteString(i.Name)
	}

	b.WriteString("}")

	b.WriteF("func JSValueTo%s(val js.Value) %s { return %s{ Value: Value { Value: val }}}", spec.Name, spec.Name, spec.Name)
	b.WriteF("func (v Value) As%s() %s { return %s{Value: v} }", spec.Name, spec.Name, spec.Name)

	mems, err = spec.ResolveMembers(g.specs, false)
	if err != nil {
		return err
	}
	for _, m := range mems {
		err := g.writeStructMember(spec, m, b)
		if err != nil {
			return err
		}
	}

	return g.writeFile(spec, b)
}

func (g *Generator) generateCallbackInterface(spec Spec) error {
	b := NewBuilder()
	g.generateFileHeader(spec, b, true)

	mems, err := spec.ResolveMembers(g.specs, true)
	if err != nil {
		return err
	}
	consts := make([]Member, 0)
	funcs := make([]Member, 0)
	for _, m := range mems {
		switch m.Type {
		case "const":
			consts = append(consts, m)
		case "operation":
			funcs = append(funcs, m)
		default:
			fmt.Printf("%#v\r\n", m)

		}
	}

	if len(consts) > 0 {
		b.WriteString("const (")
		for _, c := range consts {
			n := fmt.Sprintf("%s%s", spec.Name, c.Title())
			t := convertIDLType(c.IDLType)
			b.WriteF("%s %s = %s", n, t, c.Value.Value)
		}
		b.WriteString(")")
	}

	for _, f := range funcs {
		n := fmt.Sprintf("%s%s", spec.Name, f.Title())
		args := ""
		params := ""
		for i, a := range f.Body.Arguments {
			if i > 0 {
				args += ","
				params += ","
			}
			t := convertIDLType(a.IDLType)
			if t == "" {
				t = "Value"
			}
			args += fmt.Sprintf("%s %s", a.Name, t)
			params += a.Name
		}

		b.WriteF("type %sCallback func(%s) %s", n, args, convertIDLType(f.IDLType))
		b.WriteF("type %s struct {", n)
		b.WriteString("Callback")
		b.WriteString("}")

		b.WriteF("func JSValueTo%s(val js.Value) %s {", n, n)
		b.WriteF("return %s{ Callback: JSValueToCallback(val) }", n)
		b.WriteString("}")
		b.WriteF("func New%s(c %sCallback) %s {", n, n, n)
		b.WriteF("callback := js.NewCallback(func (args []js.Value) {")

		for i, a := range f.Body.Arguments {
			if a.Name == "c" || a.Name == "args" {
				a.Name += "Arg"
			}
			t := convertIDLType(a.IDLType)
			if t == "" {
				t = "Value"
			}
			switch t {
			case "string":
				b.WriteF("%s := args[%d].String()", a.Name, i)
			case "[]byte":
				b.WriteF("%s := []byte(args[%d].String())", a.Name, i)
			case "bool":
				b.WriteF("%s := args[%d].Bool()", a.Name, i)
			case "float64":
				b.WriteF("%s := args[%d].Float()", a.Name, i)
			case "int":
				b.WriteF("%s := args[%d].Int()", a.Name, i)
			default:
				b.WriteF("%s := JSValueTo%s(args[%d])", a.Name, t, i)
			}
		}
		b.WriteF("c(%s)", params)
		b.WriteString("})")

		b.WriteF("return %s{ Callback: Callback{ Callback: callback } }", n)
		b.WriteString("}")
	}

	// Implementation
	b.WriteF("type %s struct {", spec.Name)
	b.WriteString("Value")

	for _, f := range funcs {
		b.WriteF("%s %s%sCallback", f.Title(), spec.Name, f.Title())
	}

	b.WriteString("}")

	b.WriteF("func JSValueTo%s(val js.Value) %s { return %s{ Value: Value { Value: val }}}", spec.Name, spec.Name, spec.Name)
	b.WriteF("func (v Value) As%s() %s { return %s{Value: v} }", spec.Name, spec.Name, spec.Name)
	return g.writeFile(spec, b)
}

func (g *Generator) generateCallback(spec Spec) (err error) {
	b := NewBuilder()
	g.generateFileHeader(spec, b, true)

	args := ""
	params := ""
	for i, a := range spec.Arguments {
		if i > 0 {
			args += ","
			params += ","
		}
		t := convertIDLType(a.IDLType)
		if t == "" {
			t = "Value"
		}
		args += fmt.Sprintf("%s %s", a.Name, t)
		params += a.Name
	}
	b.WriteF("type %sCallback func(%s)", spec.Name, args)

	b.WriteF("type %s struct {", spec.Name)
	b.WriteString("Callback")
	b.WriteString("}")

	b.WriteF("func JSValueTo%s(val js.Value) %s {", spec.Name, spec.Name)
	b.WriteF("return %s{ Callback: JSValueToCallback(val) }", spec.Name)
	b.WriteString("}")
	b.WriteF("func New%s(c %sCallback) %s {", spec.Name, spec.Name, spec.Name)
	b.WriteF("callback := js.NewCallback(func (args []js.Value) {")

	for i, a := range spec.Arguments {
		if a.Name == "c" || a.Name == "args" {
			a.Name += "Arg"
		}
		t := convertIDLType(a.IDLType)
		if t == "" {
			t = "Value"
		}
		switch t {
		case "string":
			b.WriteF("%s := args[%d].String()", a.Name, i)
		case "[]byte":
			b.WriteF("%s := []byte(args[%d].String())", a.Name, i)
		case "bool":
			b.WriteF("%s := args[%d].Bool()", a.Name, i)
		case "float64":
			b.WriteF("%s := args[%d].Float()", a.Name, i)
		case "int":
			b.WriteF("%s := args[%d].Int()", a.Name, i)
		default:
			b.WriteF("%s := JSValueTo%s(args[%d])", a.Name, t, i)
		}
	}
	b.WriteF("c(%s)", params)
	b.WriteString("})")

	b.WriteF("return %s{ Callback: Callback{ Callback: callback } }", spec.Name)
	b.WriteString("}")

	return g.writeFile(spec, b)
}

func (g *Generator) generateTypedef(spec Spec) (err error) {
	b := NewBuilder()
	g.generateFileHeader(spec, b, true)

	t := convertIDLType(spec.IDLType)
	if t == "" {
		t = "Value"
	}
	b.WriteF("type %s %s", spec.Name, t)

	b.WriteF("func JSValueTo%s(val js.Value) %s {", spec.Name, spec.Name)
	switch t {
	case "string":
		b.WriteF("return %s(val.String())", spec.Name)
	case "[]byte":
		b.WriteF("return %s([]byte(val.String()))", spec.Name)
	case "bool":
		b.WriteF("return %s(val.Bool())", spec.Name)
	case "float64":
		b.WriteF("return %s(val.Float())", spec.Name)
	case "int":
		b.WriteF("return %s(val.Int())", spec.Name)
	default:
		b.WriteF("return %s(JSValueTo%s(val))", spec.Name, t)
	}
	b.WriteString("}")

	p, ok := g.specs[t]
	if ok && p.Type == "callback" {
		b.WriteF("func New%s(c %sCallback) %s {", spec.Name, p.Name, spec.Name)
		b.WriteF("return %s(New%s(c))", spec.Name, p.Name)
		b.WriteString("}")
	}

	return g.writeFile(spec, b)
}

func (g *Generator) generateEnum(spec Spec) (err error) {
	b := NewBuilder()
	g.generateFileHeader(spec, b, true)

	t := convertIDLType(spec.IDLType)
	if t == "" {
		t = "string"
	}

	b.WriteF("type %s %s", spec.Name, t)

	b.WriteString("const (")
	for _, v := range spec.Values {
		if v.Value == "" {
			v.Value = "Empty"
		}
		n := fmt.Sprintf("%s%s", spec.Name, strings.Title(v.Value))
		n = strings.Replace(n, "-", "", -1)
		n = strings.Replace(n, "/", "", -1)
		n = strings.Replace(n, "+", "", -1)
		b.WriteF("%s %s = %q", n, spec.Name, v.Value)
	}
	b.WriteString(")")

	b.WriteF("func JSValueTo%s(val js.Value) %s {", spec.Name, spec.Name)
	switch t {
	case "string":
		b.WriteF("return %s(val.String())", spec.Name)
	case "[]byte":
		b.WriteF("return %s([]byte(val.String()))", spec.Name)
	case "bool":
		b.WriteF("return %s(val.Bool())", spec.Name)
	case "float64":
		b.WriteF("return %s(val.Float())", spec.Name)
	case "int":
		b.WriteF("return %s(val.Int())", spec.Name)
	default:
		b.WriteF("return %s(JSValueTo%s(val))", spec.Name, t)
	}
	b.WriteString("}")

	return g.writeFile(spec, b)
}

func (g *Generator) generateNamespace(spec Spec) (err error) {
	d := strings.ToLower(spec.Name)
	fname := fmt.Sprintf("%s/%s.go", d, d)
	err = os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return
	}

	b := NewBuilder()
	b.WriteString("// Code generated DO NOT EDIT")
	b.WriteF("// %s", fname)
	b.WriteF("package %s", d)
	b.WriteString("import dom \"github.com/brettlangdon/go-dom\"")
	b.WriteString("import \"syscall/js\"")

	vt := "Value"
	if spec.Type == "interface" {
		vt = spec.Name
	}
	b.WriteF("var value dom.%s", vt)

	b.WriteF("func init() { value = dom.JSValueTo%s(js.Global().Get(\"%s\")) }", vt, strings.ToLower(spec.Name))

	mems, err := spec.ResolveMembers(g.specs, true)
	if err != nil {
		return
	}

	for _, m := range mems {
		switch m.Type {
		case "attribute":
			t := convertIDLType(m.IDLType)
			if t == "" {
				t = "Value"
			}
			rt := t
			switch rt {
			case "string", "[]byte", "bool", "float64", "int":
			default:
				rt = "dom." + rt
			}

			b.WriteF("func Get%s() %s {", m.Title(), rt)
			if spec.Type == "interface" {
				b.WriteF("return value.Get%s()", m.Title())
			} else {
				if rt != "" {
					b.WriteF("val := value.Get(\"%s\")", m.Name)
					switch rt {
					case "string":
						b.WriteF("return val.String()")
					case "[]byte":
						b.WriteF("return []byte(val.String())")
					case "bool":
						b.WriteF("return val.Bool()")
					case "float64":
						b.WriteF("return val.Float()")
					case "int":
						b.WriteF("return val.Int()")
					case "Value":
						b.WriteString("return val")
					default:
						b.WriteF("return dom.JSValueTo%s(val.JSValue())", t)
					}
				} else {
					b.WriteF("value.Get(\"%s\")", m.Name)
				}
			}

			b.WriteString("}")

			if !m.IsReadOnly() {
				if spec.Type == "interface" {
					b.WriteF("func Set%s(val %s) { value.Set%s(val) }", m.Title(), rt, m.Title())
				} else {
					b.WriteF("func Set%s(val %s) { value.Set(\"%s\", val) }", m.Title(), rt, m.Name)
				}
			}
		case "operation":
			n := m.Body.Name.Value
			t := convertIDLType(m.Body.IDLType)
			rt := t
			switch rt {
			case "string", "[]byte", "bool", "float64", "int":
			case "":
			default:
				rt = "dom." + rt
			}

			b.WriteF("func %s(args ...interface{}) %s {", m.Title(), rt)
			if rt != "" {
				b.WriteF("val := value.Call(\"%s\", args...)", n)
				switch rt {
				case "string":
					b.WriteString("return val.String()")
				case "[]byte":
					b.WriteString("return []byte(val.String())")
				case "bool":
					b.WriteString("return val.Bool()")
				case "float64":
					b.WriteString("return val.Float()")
				case "int":
					b.WriteString("return val.Int()")
				case "dom.Value":
					b.WriteString("return val")
				default:
					b.WriteF("return dom.JSValueTo%s(val.JSValue())", t)
				}
			} else {
				b.WriteF("value.Call(\"%s\", args...)", n)
			}
			b.WriteString("}")
		}
	}

	source, err := format.Source(b.Bytes())
	if err != nil {
		return
	}
	out, err := os.OpenFile(fname, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = out.Write(source)
	return
}

func (g *Generator) generateSpec(spec Spec) (err error) {
	switch spec.Type {
	case "interface":
		err = g.generateInterface(spec)
	case "interface mixin":
		// fmt.Printf("INTERFACE MIXIN - %s\r\n", spec.Name)
	case "callback":
		err = g.generateCallback(spec)
	case "dictionary":
		// fmt.Printf("DICTIONARY - %s\r\n", spec.Name)
	case "enum":
		err = g.generateEnum(spec)
	case "typedef":
		err = g.generateTypedef(spec)
	case "callback interface":
		err = g.generateCallbackInterface(spec)
	case "namespace":
		err = g.generateNamespace(spec)
	default:
		err = fmt.Errorf("Unknown or unsupported spec type %q", spec.Type)
	}

	if spec.Name == "Document" || spec.Name == "Window" {
		if err == nil {
			err = g.generateNamespace(spec)
		}
	}

	return
}

func (g *Generator) Generate() error {
	for _, spec := range g.specs {
		err := g.generateSpec(spec)
		if err != nil {
			return err
		}
	}
	return nil
}

package main

type IDLType struct {
	BaseName        string           `json:"baseName"`
	ExtraAttributes *ExtraAttributes `json:"extAttrs"`
	Generic         interface{}      `json:"generic"`
	IDLType         interface{}      `json:"idlType"`
	Nullable        interface{}      `json:"nullable"`
	Postfix         interface{}      `json:"postfix"`
	Prefix          interface{}      `json:"prefix"`
	Separator       interface{}      `json:"separator"`
	Trivia          interface{}      `json:"trivia"`
	Type            string           `json:"type"`
	Union           bool             `json:"union"`
}

func convertIDLType(idlType interface{}) string {
	switch t := idlType.(type) {
	case IDLType:
		return convertIDLType(t.IDLType)
	case map[string]interface{}:
		return convertIDLType(t["idlType"])
	case string:
		switch t {
		case "USVString", "DOMString", "CSSOMString":
			t = "string"
		case "unsigned long", "long", "unsigned long long",
			"double", "unrestricted double", "long long":
			t = "float64"
		case "unsigned short", "short":
			t = "int"
		case "boolean":
			t = "bool"
		case "any", "object":
			t = "Value"
		case "null", "void":
			t = ""
		case "ByteString":
			t = "[]byte"
		case "ReadableStream", "WritableStream":
			t = "Value"
		}
		return t
	}
	return ""
}

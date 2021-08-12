package config

type Types struct {
	Struct    string
	Interface string
	Function  string
}

func (t *Types) fillTypes() {
	t.Struct = "struct"
	t.Interface = "interface"
	t.Function = "function"
}

func (t *Types) GetType(fileType string) string {
	t.fillTypes()

	switch fileType {
	case t.Struct:
		f := "package %v\n\n// %v - Your comment goes here\ntype %v struct {\n\t\n}\n"
		return f
	case t.Interface:
		f := "package %v\n\n// %v - Your comment goes here\ntype %v interface {\n\t\n}\n"
		return f
	case t.Function:
		f := "package %v\n\n// %v - Your comment goes here\nfunc %v () {\n\t\n}\n"
		return f
	default:
		f := "package %v\n\n// %v - Your comment goes here\nfunc %v () {\n\t\n}\n"
		return f
	}
}

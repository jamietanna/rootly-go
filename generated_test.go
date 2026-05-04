package rootly

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestNoInterfaceTypeAliases(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "schema.gen.go", nil, 0)
	if err != nil {
		t.Fatalf("failed to parse schema.gen.go: %v", err)
	}

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok || !typeSpec.Assign.IsValid() {
				continue
			}

			t.Run(typeSpec.Name.Name+" should not be an interface{}", func(t *testing.T) {
				iface, ok := typeSpec.Type.(*ast.InterfaceType)

				if ok && iface.Methods.NumFields() == 0 {
					t.Fail()
				}
			})
		}
	}
}

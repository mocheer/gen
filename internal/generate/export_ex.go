package generate

import (
	"fmt"

	"gorm.io/gen/internal/parser"
)

// AddMethod
func (g *QueryStructMeta) AddMethod(method interface{}) {
	g.addMethodFromAddMethodOpt(method)
}

func (g *QueryStructMeta) AddMethodWidthMapTest(value map[string]any) {
	method := &parser.Method{
		Receiver:   parser.Param{IsPointer: true, Type: g.ModelStructName},
		MethodName: "ToSchemaBuffer",
		Doc:        "",
		Params:     []parser.Param{},
		Result: []parser.Param{
			{
				Type: "[]byte",
			},
		},
		Body: fmt.Sprintf("{\n\treturn %s\n} ", value["body"].(string)),
	}
	g.ModelMethods = append(g.ModelMethods, method)
}

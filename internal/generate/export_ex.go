package generate

// AddMethod
func (g *QueryStructMeta) AddMethod(method interface{}) {
	g.addMethodFromAddMethodOpt(method)
}

package interfaces

type Router interface {
	Group(prefix string) Router
	GET(path string, hanlder Handler)
	POST(path string, hanlder Handler)
	PUT(path string, handler Handler)
	DELETE(path string, handler Handler)
	LoadHTMLFiles(files ...string) error
	Static(relativePath string, path string)
	Run(addr string) error
}

type Handler func(ctx Context)
type HandlerFunc func(ctx *Context)

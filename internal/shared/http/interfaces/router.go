package interfaces

type Router interface {
	Group(prefix string) Router
	GET(path string, hanlder Handler)
	POST(path string, hanlder Handler)
	DELETE(path string, handler Handler)
	LoadHTMLFiles(files ...string) error
	Run(addr string) error
}

type Handler func(ctx Context)
type HandlerFunc func(ctx *Context)

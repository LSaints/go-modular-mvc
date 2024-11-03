package adapters

import (
	"fmt"

	"github.com/LSaints/go-modular-mvc/internal/shared/http/interfaces"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
}

func NewGinRouter() *GinRouter {
	r := gin.Default()
	return &GinRouter{
		engine: r,
	}
}

func (g *GinRouter) Group(prefix string) interfaces.Router {
	group := g.engine.Group(prefix)
	return &GinRouter{engine: g.engine, routerGroup: group}
}

func (g *GinRouter) GET(path string, handler interfaces.Handler) {
	g.engine.GET(path, func(ctx *gin.Context) {
		handler(&GinContext{ctx})
	})
}

func (g *GinRouter) POST(path string, handler interfaces.Handler) {
	g.engine.POST(path, func(ctx *gin.Context) {
		handler(&GinContext{ctx})
	})
}

func (g *GinRouter) DELETE(path string, handler interfaces.Handler) {
	g.engine.DELETE(path, func(ctx *gin.Context) {
		handler(&GinContext{ctx})
	})
}

func (g *GinRouter) PUT(path string, handler interfaces.Handler) {
	g.engine.PUT(path, func(ctx *gin.Context) {
		handler(&GinContext{ctx})
	})
}

func (g *GinRouter) LoadHTMLFiles(files ...string) error {
	if g.engine == nil {
		return fmt.Errorf("GinRouter não está inicializado")
	}
	g.engine.LoadHTMLFiles(files...)
	return nil
}

func (g *GinContext) FormValue(name string) string {
	return g.Context.PostForm(name)
}

func (g *GinContext) Redirect(code int, location string) {
	g.Context.Redirect(code, location)
}

func (g *GinRouter) Run(addr string) error {
	return g.engine.Run(addr)
}

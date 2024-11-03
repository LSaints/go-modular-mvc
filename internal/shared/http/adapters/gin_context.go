package adapters

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	*gin.Context
}

func (g *GinContext) Param(name string) string {
	return g.Context.Param(name)
}

func (g *GinContext) HTML(code int, view string, data interface{}) {
	g.Context.HTML(code, view, data)
}

func (g *GinContext) Query(name string) string {
	return g.Context.Query(name)
}

func (g *GinContext) ShouldBind(obj any) error {
	err := g.Context.ShouldBind(obj)
	if err != nil {
		return err
	}
	return nil
}

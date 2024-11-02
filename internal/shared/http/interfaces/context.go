package interfaces

type Context interface {
	Param(name string) string
	HTML(code int, view string, data interface{})
	ShouldBind(obj interface{}) error
}

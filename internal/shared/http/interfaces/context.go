package interfaces

type Context interface {
	Param(name string) string
	HTML(code int, view string, data interface{})
	ShouldBind(obj interface{}) error
	FormValue(name string) string
	Redirect(code int, location string)
	Query(name string) string
}

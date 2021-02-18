package controller

type HttpContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
}

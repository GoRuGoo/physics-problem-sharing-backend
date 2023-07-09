package controller

type Context interface {
	Praram(string) string
	JSON(int, interface{})
}

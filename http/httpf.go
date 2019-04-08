package http

type ApiEndPoint interface {
	BindGET(api string, f interface{})
	BindPOST(api string, f interface{})
}

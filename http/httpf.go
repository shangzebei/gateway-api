package http

import "net/http"

type ApiEndPoint interface {
	BindGET(api string, f interface{})
	BindPOST(api string, f interface{})
	BindResource(web string, handler http.Handler)
}

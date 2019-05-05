package filterf

import (
	"github.com/shangzebei/gateway-api/backend"
	"github.com/shangzebei/gateway-api/http"
	"github.com/shangzebei/gateway-api/httpstatus"
	"github.com/shangzebei/gateway-api/pluginf"
	"net/url"
)

type RequestContext interface {
	GetPath() string
	SetPath(path string)
	GetHeader(key string) string
	SetHeader(key string, value string)
	DelHeader(key string)
	GetParam(key string) string
	GetAllParams() url.Values
	AppendParam(k string, v string)
	ReWriteUrl(match string, to string) error
	GetContentType() string
	GetHost() string
	SetTag(tag interface{})
	GetTag() interface{}
	SetValue(filter Filter, v interface{})
	GetValue(filter Filter) interface{}
}

type RouteContext interface {
	GetOrig() string
	SetTo(string)
	GetTo() string
	SetTag(tag interface{})
	GetTag() interface{}
	SetValue(filter Filter, v interface{})
	GetValue(filter Filter) interface{}
}

type ResponseContext interface {
	SetBody(body []byte)
	GetBody() []byte
	SetHeader(k string, v string)
	SetContentType(t string)
	DelHeader(k string)
	GetTag() interface{}
	SetTag(tag interface{})
	SetStatusCode(code int)
	SetValue(filter Filter, v interface{})
	GetValue(filter Filter) interface{}
}

type Context interface {
	GetMethod() string
	GetRequest() RequestContext
	GetResponse() ResponseContext
}

type Filter interface {
	Order() int
	Pre(ctx RequestContext) (httpstatus.HTTPSTATUS, bool)
	//now only chang remote url
	Route(ctx RouteContext) (httpstatus.HTTPSTATUS, bool)
	Post(ctx ResponseContext) (httpstatus.HTTPSTATUS, bool)
	Err(err error, ctx ResponseContext) bool
}

type Pre interface {
	Order() int
	Pre(ctx RequestContext) (httpstatus.HTTPSTATUS, bool)
}

type Route interface {
	Order() int
	Route(ctx RouteContext) (httpstatus.HTTPSTATUS, bool)
}

type Post interface {
	Order() int
	Post(ctx ResponseContext) (httpstatus.HTTPSTATUS, bool)
}

type Error interface {
	Order() int
	Err(err error, ctx ResponseContext) bool
}

type FilterEntry struct {
	Info    *pluginf.PlugInfo
	Http    http.ApiEndPoint
	Backend backend.Backend
}

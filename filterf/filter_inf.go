package filterf

import (
	"github.com/shangzebei/gateway-api/httpstatus"
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
	SetTag(tag string)
	GetTag() string
}

type ResponseContext interface {
	SetBody(body []byte)
	GetBody() []byte
	SetHeader(k string, v string)
	SetContentType(t string)
	DelHeader(k string)
	GetTag() string
	SetTag(tag string)
	SetStatusCode(code int)
}

type Context interface {
	GetMethod() string
	GetRequest() RequestContext
	GetResponse() ResponseContext
}

type Filter interface {
	Pre(ctx RequestContext) (httpstatus.HTTPSTATUS, bool)
	Route(ctx Context) (httpstatus.HTTPSTATUS, bool)
	Post(ctx ResponseContext) (httpstatus.HTTPSTATUS, bool)
	Err(err error, ctx ResponseContext) bool
}

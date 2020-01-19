package core

import (
	"github.com/lightjiang/OneBD/rfc"
	"net/http"
)

type Router interface {
	Set(prefix string, fc func() Handler, allowedMethods ...rfc.Method)
	SubRouter(name string) Router
	ServeHTTP(http.ResponseWriter, *http.Request)
	SetStatusFunc(status rfc.Status, rf MetaFunc)
	SetNotFoundFunc(fc MetaFunc)
	// 对于内部panic错误， 返回500 错误 并执行相应回调
	SetInternalErrorFunc(fc MetaFunc)
}

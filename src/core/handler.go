package core

import "github.com/kataras/iris/v12"

type HandlerBuilder interface {
	Authorize(authorities ...string) iris.Handler
	Scope(scope ...string) iris.Handler
	Handler(handler interface{}) iris.Handler
	Register(values ...interface{}) HandlerBuilder
}

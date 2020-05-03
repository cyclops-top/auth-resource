package context

import (
	"auth-resource/src/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"reflect"
)

type Context interface {
	iris.Context
	User() *model.User
	SetUser(user *model.User)
}


const (
	reflectValueContextKey = "_auth_context_context_reflect_value"
)

func New(app *iris.Application) *context.Pool {
	return context.New(func() context.Context {
		return &contextImpl{Context: context.NewContext(app)}
	})
}

type contextImpl struct {
	iris.Context
	user *model.User
}

func (c *contextImpl) User() *model.User {
	return c.user
}

func (c *contextImpl) SetUser(user *model.User) {
	c.user = user
}

func (c *contextImpl) Next() {
	context.Next(c)
}

func (c *contextImpl) Do(handlers context.Handlers) {
	context.Do(c, handlers)
}


func (c *contextImpl) ReflectValue() []reflect.Value {
	if v := c.Values().Get(reflectValueContextKey); v != nil {
		return v.([]reflect.Value)
	}
	v := []reflect.Value{reflect.ValueOf(c)}
	c.Values().Set(reflectValueContextKey, v)
	return v
}

func (c *contextImpl) EndRequest() {
	c.Context.EndRequest()
	c.SetUser(nil)
}

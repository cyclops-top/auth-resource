package workshop

import (
	"auth-resource/src/context"
	"auth-resource/src/core"
	"auth-resource/src/model"
	"auth-resource/src/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"sort"
	"strings"
)

type Workshop struct {
	h *hero.Hero
}

func New() *Workshop {
	w := &Workshop{h: hero.New()}
	w.h.Register(func(ctx iris.Context) *model.User {
		return ctx.(context.Context).User()
	})
	return w
}

func (w *Workshop) Clone() *Workshop {
	return &Workshop{
		h: w.h.Clone(),
	}
}

func havSame(a1 []string, a2 []string) bool {
	sort.Strings(a1)
	sort.Strings(a2)
	l1 := len(a1)
	l2 := len(a2)
	for i, j := 0, 0; i < l1 && j < l2; {
		switch strings.Compare(a1[i], a2[j]) {
		case 0:
			return true
		case 1:
			i++
			continue
		case -1:
			j++
			continue
		}
	}
	return false
}

func (w *Workshop) Authorize(authorities ...string) iris.Handler {
	return w.h.Handler(func(ctx context.Context) {
		user := ctx.User()
		if user == nil{
			ctx.StatusCode(401)
			_, _ = ctx.WriteString("not login")
			return
		}
		if user.Authorities == nil {
			ctx.StatusCode(403)
			_, _ = ctx.WriteString("authorities is not allowed")
			return
		}
		if havSame(user.Authorities, authorities) {
			ctx.Next()
			return
		}
		ctx.StatusCode(403)
		_, _ = ctx.WriteString("authorities is not allowed")
	})
}

func (w *Workshop) Scope(scope ...string) iris.Handler {
	return w.h.Handler(func(ctx context.Context) {
		user := ctx.User()
		if  user == nil{
			ctx.StatusCode(401)
			_, _ = ctx.WriteString("not login")
			return
		}
		if user.Scope == nil {
			ctx.StatusCode(403)
			_, _ = ctx.WriteString("scope is not allowed")
			return
		}
		if havSame(user.Scope, scope) {
			ctx.Next()
			return
		}
		ctx.StatusCode(403)
		_, _ = ctx.WriteString("scope is not allowed")
	})
}

func (w *Workshop) Register(values ...interface{}) core.HandlerBuilder {
	w.h.Register(values)
	return w
}

func (w *Workshop) Handler(fn interface{}) iris.Handler {
	return w.h.Handler(fn)
}

func UserLoader(authorizeUrl string) iris.Handler {
	return func(ctx iris.Context) {
		c := ctx.(context.Context)
		user, err := service.Authorize(c,authorizeUrl)
		if err != nil {
			c.StatusCode(401)
			_, _ = c.WriteString(err.Error())
			return
		}
		c.SetUser(user)
		c.Next()
	}
}

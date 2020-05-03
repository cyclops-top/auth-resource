package auth_resource

import (
	"github.com/cyclops-top/auth-resource/src/context"
	"github.com/cyclops-top/auth-resource/src/core"
	"github.com/cyclops-top/auth-resource/src/workshop"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type Application struct {
	*iris.Application
	*workshop.Workshop
}
type Party2 router.Party

func New(authorizeUrl string) *Application {
	irisApp := iris.New()
	app := &Application{
		irisApp, workshop.New(),
	}
	app.ContextPool = context.New(irisApp)
	app.Use(workshop.UserLoader(authorizeUrl))
	return app
}

type Party interface {
	Party2
	core.HandlerBuilder
	Part(relativePath string, handlers ...iris.Handler) Party
}

type party struct {
	Party2
	*workshop.Workshop
}

func (api *Application) Part(relativePath string, handlers ...iris.Handler) Party {
	p := api.Application.Party(relativePath, handlers...)
	return &party{p, api.Workshop.Clone()}
}

func (p *party) Part(relativePath string, handlers ...iris.Handler) Party {
	part := p.Party(relativePath, handlers...)
	return &party{part, p.Workshop.Clone()}
}

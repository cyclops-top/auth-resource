package main

import (
	_ "auth-resource/docs"
	authResource "auth-resource/src"
	"auth-resource/src/model"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

func main() {
	app:=authResource.New("http://127.0.0.1:9091/auth/authorize")
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},   //允许通过的主机名称
		AllowCredentials: true,
	})
	app.UseGlobal(crs)
	config := &swagger.Config{
		URL: "http://127.0.0.1:9092/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	app.Get("/hello/{name:string}",app.Authorize("ADMIN"), app.Handler(test))
	p := app.Part("/test")
	p2 := p.Part("/t2",p.Authorize("ADMIN"))
	p2.Get("/t3/{name:string}",p2.Scope("test"),p2.Handler(test))
	_ = app.Run(iris.Addr(fmt.Sprintf(":%d", 9092)), iris.WithoutServerError(iris.ErrServerClosed))

}
// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.User
// @Header 200 {string} Token "qwerty"
// @Router /accounts [get]
func test(user *model.User, name string) *model.User {
	fmt.Println(user.String())
	return user
}
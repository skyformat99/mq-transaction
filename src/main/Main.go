package main

import (
	"github.com/kataras/iris"
)

func main()  {
 startHttp()
}
func startHttp() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"hello": "yes hello!"})
	})
	app.Run(iris.Addr("localhost:8080"), iris.WithoutServerError(iris.ErrServerClosed))
}


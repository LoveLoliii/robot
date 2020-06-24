package main

import (
	"github.com/kataras/iris/v12"
)

// test new ssh key
func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	//app.Use(myMiddleware)
	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})

	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.Writef("Hello from the SECURE server")
	// })

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from the SECURE server on path /mypath")
	})
	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})
	// 使用nginx监听80/443 再转发到8888
	app.Listen(":8888")
}
func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

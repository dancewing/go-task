package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	app.Config.IsDevelopment = true

	// Load all templates from the "./templates" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package.
	app.RegisterView(iris.HTML("./templates", ".html"))
	//app.RegisterView(iris.BASIC("./templates", ".tmpl"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx context.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./templates/hello.html
		ctx.View("hello.html")
	})

	// Start the server using a network address and block.
	app.Run(iris.Addr(":8080"))
}

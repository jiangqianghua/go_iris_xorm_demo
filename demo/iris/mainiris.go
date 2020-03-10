package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 从当前路径加载html模板
	htmlEngine := iris.HTML("./", ".html")
	// 注册模板引擎
	app.RegisterView(htmlEngine)
	// http://localhost:8085/
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello iris...")
	})

	app.Get("hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "模板页面")
		//  模板引擎渲染
		ctx.View("template.html")
	})

	app.Run(iris.Addr(":8085"), iris.WithCharset("UTF-8"))
}

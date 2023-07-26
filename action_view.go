package main

import "github.com/kataras/iris/v12"

func action_welcome(ctx iris.Context) {
	ctx.View("welcome.html")
}

func action_fileman(ctx iris.Context) {
	if err := ctx.View("fileman_view.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

func action_scan(ctx iris.Context) {
	ctx.View("scan.html")
}

func action_print(ctx iris.Context) {
	ctx.View("print.html")
}

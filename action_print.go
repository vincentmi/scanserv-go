package main

import "github.com/kataras/iris/v12"

func action_print(ctx iris.Context) {
	ctx.View("print.html")
}

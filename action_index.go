package main

import "github.com/kataras/iris/v12"

func action_welcome(ctx iris.Context) {
	ctx.View("welcome.html")
}

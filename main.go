package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kataras/iris/v12"
)

//go:embed templates
var templateFs embed.FS

//go:embed static
var staticFs embed.FS

var app = iris.New()
var port int
var filesPath string
var scanCommand string

func check_dir(dir string) bool {
	s, err := os.Stat(dir)
	if err != nil || !s.IsDir() {
		app.Logger().Errorf(" dir [%s] invalid", dir)
		return false
	}
	return true
}

// func check_is_file(file string) bool {
// 	s, err := os.Stat(file)
// 	if err != nil || s.IsDir() {
// 		app.Logger().Errorf(" file [%s] invalid", file)
// 		return false
// 	}
// 	return true
// }

func main() {

	flag.IntVar(&port, "p", 8080, "监听端口")
	flag.StringVar(&filesPath, "f", "./file", "上传和中转文件路径")
	flag.StringVar(&scanCommand, "m", "/usr/bin/scanimage", "扫描命令")
	flag.Parse()

	if !check_dir(filesPath) {
		return
	}
	filesPath, _ = filepath.Abs(filesPath)

	app.UseRouter(func(ctx iris.Context) {
		app.Logger().Info(fmt.Sprintf("%s -> %s",
			ctx.Request().RemoteAddr,
			ctx.Request().RequestURI))
		ctx.Next()

	})

	app.Logger().SetLevel("debug")

	app.Logger().Info(fmt.Sprintf("listen at [%s]", strconv.Itoa(port)))
	app.Logger().Info(fmt.Sprintf("use  [%s] as temp folder ", filesPath))
	app.Logger().Info(fmt.Sprintf("use scan command  [%s]  ", scanCommand))

	//载入模板
	ftemp := iris.PrefixDir("templates", http.FS(templateFs))
	var tmpl = iris.HTML(ftemp, ".html")
	tmpl.Delims("{%", "%}")
	app.RegisterView(tmpl)
	//静态文件
	fstatic := iris.PrefixDir("static", http.FS(staticFs))
	app.HandleDir("/static", fstatic)
	app.HandleDir("/file", iris.Dir(filesPath))
	app.Use(iris.Compression)

	//注册路由
	app.Get("/fileman/view", action_fileman)
	app.Get("/fileman/list", action_file_list)
	app.Get("/fileman/delete", action_file_delete)
	app.Get("/serv/scan", action_scan)
	app.Get("/serv/do_scan", action_doscan)

	app.Get("/serv/print", action_print)
	app.Get("/", action_welcome)
	app.Get("/info", func(ctx iris.Context) {
		ctx.JSON(Success("pong"))
	})
	app.Listen(":" + strconv.Itoa(port))
}

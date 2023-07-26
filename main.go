package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kataras/iris/v12"
)

var app = iris.New()

var configFile string
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

func check_file(file string) bool {
	s, err := os.Stat(file)
	if err != nil || s.IsDir() {
		app.Logger().Errorf(" file [%s] invalid", file)
		return false
	}
	return true
}

func main() {

	flag.StringVar(&configFile, "c", "./config.yaml", "配置文件config.yaml")
	flag.IntVar(&port, "p", 8080, "监听端口")
	flag.StringVar(&filesPath, "f", "./file", "上传和中转文件路径")
	flag.StringVar(&scanCommand, "m", "/bin/scanimage", "扫描命令")
	flag.Parse()

	if !check_dir(filesPath) {
		return
	}

	app.UseRouter(func(ctx iris.Context) {
		app.Logger().Info(fmt.Sprintf("%s -> %s",
			ctx.Request().RemoteAddr,
			ctx.Request().RequestURI))
		ctx.Next()

	})

	app.Logger().Info(fmt.Sprintf("loading config [%s]", configFile))

	//载入模板
	var tmpl = iris.HTML("./templates", ".html")
	app.RegisterView(tmpl)
	//静态文件
	app.HandleDir("/static", iris.Dir("./static"))
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
	app.Listen(":" + strconv.Itoa(port))
}

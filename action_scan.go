package main

import (
	"os/exec"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

type DoScanResponse struct {
	Path string `json:"path"`
}

func action_doscan(ctx iris.Context) {
	resp := DoScanResponse{}
	args := ""

	filename := time.Now().Format("20060102150405_") + strconv.Itoa(time.Now().Nanosecond()) + randStr(5) + ".pdf"

	ctx.Application().Logger().Infof("scan file => %s", filename)

	cmd := exec.Command(scanCommand, args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		ctx.JSON(FailMsg("执行扫描命令错误:" + err.Error()))
		return
	}
	rest := string(out)

	ctx.Application().Logger().Infof("scan return => %s", rest)

	if time.Now().Second()%2 == 0 {
		resp.Path = "/file/a.jpg"
	} else {
		resp.Path = "/file/b.pdf"
	}

	ctx.JSON(Success(resp))
}

func action_scan(ctx iris.Context) {
	ctx.View("scan.html")
}

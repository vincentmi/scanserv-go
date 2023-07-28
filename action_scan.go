package main

import (
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

type DoScanResponse struct {
	Path string `json:"path"`
}

func check_contain(options []string, s string) bool {
	for _, item := range options {
		if item == s {
			return true
		}
	}
	return false
}
func check_value(s string, min int, max int) bool {
	s1, e := strconv.Atoi(s)
	if e != nil {
		return false
	}

	if s1 >= min && s1 <= max {
		return true
	}
	return false
}

func action_doscan(ctx iris.Context) {

	format := ctx.URLParamDefault("format", "jpeg")
	format = strings.ToLower(format)
	if !check_contain([]string{"pdf", "jpeg", "png"}, format) {
		format = "jpeg"
	}

	resolution := ctx.URLParamDefault("resolution", "300")
	if !check_contain([]string{"70", "100", "150", "200", "300", "600", "1200"}, resolution) {
		format = "300"
	}

	mode := ctx.URLParamDefault("mode", "Color")
	if !check_contain([]string{"Lineart", "Gray", "Color"}, mode) {
		format = "Color"
	}
	contrast := ctx.URLParamDefault("contrast", "6")
	if !check_value(contrast, 1, 11) {
		contrast = "6"
	}

	brightness := ctx.URLParamDefault("brightness", "100")
	if !check_value(brightness, 0, 200) {
		contrast = "100"
	}

	l := ctx.URLParamDefault("l", "0")
	if !check_value(brightness, 0, 215) {
		l = "0"
	}

	t := ctx.URLParamDefault("t", "0")
	if !check_value(brightness, 0, 381) {
		t = "0"
	}

	x := ctx.URLParamDefault("x", "215")
	if !check_value(brightness, 0, 215) {
		x = "100"
	}

	y := ctx.URLParamDefault("y", "381")
	if !check_value(brightness, 0, 381) {
		y = "100"
	}

	resp := DoScanResponse{}

	filename := time.Now().Format("20060102150405_") + strconv.Itoa(time.Now().Nanosecond()) + randStr(5) + "." + format

	output := filesPath + "/" + filename
	fileurl := "/file/" + filename

	args := " --format=" + format + "  --resolution=" + resolution
	args += " --source=Flatbed "
	args += " --mode=" + mode
	args += " --contrast=" + contrast
	args += " --brightness=" + brightness
	args += " -l  " + l
	args += " -t " + t
	args += " -x " + x
	args += " -y " + y
	//args += " --progress"
	args += " > " + output

	ctx.Application().Logger().Infof("scan file => %s", filename)
	ctx.Application().Logger().Infof("scan command => %s %s", scanCommand, args)

	cmd := exec.Command(scanCommand, args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		ctx.JSON(FailMsg("执行扫描命令错误:" + err.Error()))
		return
	}
	rest := string(out)

	ctx.Application().Logger().Infof("scan return => %s", rest)

	resp.Path = fileurl

	ctx.JSON(Success(resp))
}

func action_scan(ctx iris.Context) {
	ctx.View("scan.html")
}

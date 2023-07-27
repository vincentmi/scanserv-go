package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/kataras/iris/v12"
)

func action_fileman(ctx iris.Context) {
	if err := ctx.View("fileman_view.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

type PathDesc struct {
	Path      string `json:"path"`
	Url       string `json:"url"`
	IsFile    bool   `json:"is_file"`
	Name      string `json:"name"`
	CanDelete bool   `json:"can_delete"`
}

type FileListResponse struct {
	Root     string     `json:"root"`
	FileList []PathDesc `json:"file_list"`
	NavList  []PathDesc `json:"nav_list"`
}

func trim_and_check_path(path string) ([]string, error) {
	fixedDir := []string{}
	path = strings.TrimLeft(path, "/")
	for _, part := range strings.Split(path, "/") {
		part = strings.TrimSpace(part)
		if part == "." || part == ".." {
			return nil, errors.New("无效的路径")
		}
		if part != "" {
			fixedDir = append(fixedDir, part)
		}
	}
	return fixedDir, nil
}

func action_file_delete(ctx iris.Context) {
	dir := ctx.URLParamDefault("path", "/")
	fixedDir, err := trim_and_check_path(dir)

	if err != nil {
		ctx.JSON(FailMsg(err.Error()))
		return
	}
	relateDir := strings.Join(fixedDir, string(os.PathSeparator))
	realPath := filesPath + string(os.PathSeparator) + relateDir
	ctx.Application().Logger().Warnf("delete file %s", realPath)
	// if !check_file(realPath) {
	// 	ctx.JSON(FailMsg("目录无法删除"))
	// 	return
	// }

	e := os.Remove(realPath)

	if e != nil {
		ctx.JSON(FailMsg(e.Error()))
	} else {
		ctx.JSON(Success(nil))
	}

}

func action_file_list(ctx iris.Context) {
	dir := ctx.URLParamDefault("dir", "/")

	fixedDir, err := trim_and_check_path(dir)

	if err != nil {
		ctx.JSON(FailMsg(err.Error()))
		return
	}

	navList := []PathDesc{}
	tempFixDir := []string{}
	for _, part := range fixedDir {
		tempFixDir = append(tempFixDir, part)
		pathItem := PathDesc{}
		pathItem.Name = part
		pathItem.Path = strings.Join(tempFixDir, "/")
		pathItem.Url = pathItem.Path
		navList = append(navList, pathItem)
	}

	relateDir := strings.Join(fixedDir, string(os.PathSeparator))
	relateUrl := strings.Join(fixedDir, "/")
	realPath := filesPath + string(os.PathSeparator) + relateDir

	abspath, err := filepath.Abs(realPath)

	if err != nil {
		ctx.JSON(FailMsg("无效的路径r"))
		return
	}

	fileList := []PathDesc{}

	files, err := os.ReadDir(abspath)
	if err != nil {
		ctx.JSON(FailMsg(err.Error()))
		return
	}

	for _, info := range files {
		//ctx.Application().Logger().Info(fmt.Sprintf("-%s", info.Name()))
		pathDisc := PathDesc{}
		pathDisc.Name = info.Name()
		pathDisc.Path = relateUrl + "/" + info.Name()
		pathDisc.Url = "/file/" + info.Name()
		pathDisc.IsFile = !info.IsDir()
		pathDisc.CanDelete = true
		fileList = append(fileList, pathDisc)
	}

	resp := FileListResponse{}
	resp.Root = relateDir
	resp.FileList = fileList
	resp.NavList = navList
	ctx.JSON(Success(resp))

}

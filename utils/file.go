package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Exist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

// WalkDirFiles 从路径dirPth下获取全部的文件.
func WalkDirFiles(dirPath string, handler func(base string, fileName string, info os.FileInfo) error) error {
	var err error
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return err
	}
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf(`%s: file info is nil`, path)
		}
		if err := handler(dirPath, path, info); err != nil {
			return err
		}
		return nil
	})
}

var _pwd = ""

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_pwd = pwd
}
func FormatFile(filename string) string {
	if !filepath.IsAbs(filename) {
		return filename
	}
	rel, err := filepath.Rel(_pwd, filename)
	if err != nil {
		return filename
	}
	if strings.HasPrefix(rel, "..") {
		return filename
	}
	return rel
}

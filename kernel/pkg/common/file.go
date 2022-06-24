package common

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) bool {
	/*
	   判断文件或文件夹是否存在
	   如果返回的错误为nil,说明文件或文件夹存在
	   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	   如果返回的错误为其它类型,则不确定是否在存在
	*/
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CopyFile(from, to string) {
	fileInfo, err := os.Stat(from)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if fileInfo.IsDir() {
		if fileList, err := ioutil.ReadDir(from); err == nil {
			for _, item := range fileList {
				CopyFile(
					filepath.Join(from, item.Name()),
					filepath.Join(to, item.Name()))
			}
		}
	} else {
		path := filepath.Dir(to)
		if _, err := os.Stat(path); err != nil {
			if e := os.MkdirAll(path, 0777); e != nil {
				return
			}
		}
		ffile, err := os.Open(from)
		if err != nil {
			return
		}
		defer ffile.Close()

		tfile, err := os.Create(to)
		if err != nil {
			return
		}
		defer tfile.Close()

		io.Copy(tfile, bufio.NewReader(ffile))
	}
}

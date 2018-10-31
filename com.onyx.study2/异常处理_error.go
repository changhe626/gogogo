package main

import (
	"errors"
	"io/ioutil"
	"os"
)

func readFile(path string) ([]byte, error) {

	if path == "" {
		//errors.New是一个很常用的函数。
		// 在Go语言标准库的代码包中有很多由此函数创建出来的错误值，比如os.ErrPermission、io.EOF等变量的值。
		return nil, errors.New("path is null")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func main() {

}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.tianmao.com/",
	}

	for _, v := range urls {
		response, err := http.Head(v)
		if err != nil {
			fmt.Println("Error:", v, err)
		}
		fmt.Print(v, ": ", response.Status)

	}

}

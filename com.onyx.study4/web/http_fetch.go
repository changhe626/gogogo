package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("http://www.baidu.com")
	checkerror(err)
	data, err := ioutil.ReadAll(response.Body)
	checkerror(err)
	fmt.Printf("get : %q", string(data))

}
func checkerror(e error) {
	if e != nil {
		log.Fatalf("Get : %v", e)
	}
}

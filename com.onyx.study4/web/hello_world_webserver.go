package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside HelloServer handler")
	// fmt.Fprint  和 fmt.Fprintf  都是用来写入 http.ResponseWriter  的不错的函数
	fmt.Fprint(w, "hello ,"+r.URL.Path[1:])
}

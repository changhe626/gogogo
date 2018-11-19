// httpProxy project main.go
package main

import (
	_ "fmt"
	"io"
	_ "log"
	"net/http"
	"net/http/httputil"
	"net/url"
	_ "strings"
)

func main() {
	/*localHost := "127.0.0.1:9293"
	targetHost := "127.0.0.1:29383"
	httpsServer(localHost, targetHost)
	var err error = nil*/
	http.HandleFunc("/", ServeHTTP)
	//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//		if req.RequestURI == "/favicon.ico" {
	//			io.WriteString(w, "Request path Error")
	//			return
	//		}
	//		fmt.Println(req.RequestURI)
	//		fmt.Fprintf(w, req.RequestURI)
	//	})
	http.ListenAndServe(":8085", nil)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		io.WriteString(w, "Request path Error")
		return
	}
	remote, err := url.Parse("http://" + "127.0.0.1:80")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

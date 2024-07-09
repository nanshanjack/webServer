package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatalln(http.ListenAndServe(":9999", nil))
}

type Engine struct {
}

func (engine *Engine) ServeHttp(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":

	}
}
func helloHandler(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		_, err := fmt.Fprintf(writer, "Header[%q]=%q\n", k, v)
		if err != nil {
			return
		}
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Url.path=%q\n", request.URL.Path)
	if err != nil {
		return
	}
}

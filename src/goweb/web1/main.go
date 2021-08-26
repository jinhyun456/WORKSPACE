package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //"/" 절대경로지정 및 HandleFunc 핸들러등록
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar!")
	})

	http.Handle("/foo", &fooHandler{}) //인스턴스 형태로 구현

	// http.ListenAndServe(":3000", nil) // 웹서버 구동

}

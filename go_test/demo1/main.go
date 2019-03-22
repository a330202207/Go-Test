package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// 控制器
	var handler = func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "你大爷")
	}

	var f = http.HandlerFunc(handler)

	var middle = func(next http.Handler) http.Handler {

		// 中间件
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	log.Panicln(http.ListenAndServe("127.0.0.1:8080", middle(f)))
}

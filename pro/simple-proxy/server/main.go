package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	// 创建路由
	http.HandleFunc("/", hello)
	// 创建服务端
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		return
	}
}

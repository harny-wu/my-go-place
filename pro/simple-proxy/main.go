package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Proxy 多后端（target）代理
type Proxy struct {
}

// NewHttpReverseProxy 创建代理
func (p *Proxy) NewHttpReverseProxy(target *url.URL) *httputil.ReverseProxy {
	rp := httputil.NewSingleHostReverseProxy(target)
	return rp
}

// ServeHTTP 实现接口：ServeHTTP(w http.ResponseWriter, r *http.Request)
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var target *url.URL
	if r.URL.Path == "/test1" {
		target, _ = url.Parse("http://127.0.0.1:2003")
	} else if r.URL.Path == "/test2" {
		target, _ = url.Parse("http://127.0.0.1:2004")
	} else {
		target, _ = url.Parse("http://127.0.0.1:2005")
	}
	proxy := p.NewHttpReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func main() {
	addr := ":2002"
	p := &Proxy{}
	log.Println("Starting HttpServer at " + addr)
	err := http.ListenAndServe(addr, p)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

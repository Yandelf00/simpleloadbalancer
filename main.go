package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	Backend *[]Backend
	current uint64
}

func main() {
	u, _ := url.Parse("http://localhost:8000")
	rp := httputil.NewSingleHostReverseProxy(u)

	fmt.Println("hello world")
}

package srvnode

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

func NewSrvNode(url *url.URL, proxy *httputil.ReverseProxy) *SrvNode {
	return &SrvNode{
		URL:          url,
		Alive:        true,
		ReverseProxy: proxy,
		mux:          sync.RWMutex{},
	}
}

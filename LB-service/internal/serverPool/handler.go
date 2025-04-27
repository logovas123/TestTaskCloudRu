package serverpool

import (
	"net/http/httputil"
	"net/url"
)

// абстракция над сервером пула
type ServerNodeHandler interface {
	IsAlive() bool
	SetAlive(alive bool)
	GetURL() *url.URL
	GetReverseProxy() *httputil.ReverseProxy
}

type ServerPool struct {
	listOfSrvs []ServerNodeHandler
	current    uint64
}

func NewServerPool() *ServerPool {
	return &ServerPool{
		listOfSrvs: []ServerNodeHandler{},
		current:    0,
	}
}

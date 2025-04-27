package srvnode

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type SrvNode struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

// устанавливает новый статус сервера
func (b *SrvNode) SetAlive(alive bool) {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.Alive = alive
}

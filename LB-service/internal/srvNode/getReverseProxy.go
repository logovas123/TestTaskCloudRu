package srvnode

import "net/http/httputil"

func (b *SrvNode) GetReverseProxy() *httputil.ReverseProxy {
	return b.ReverseProxy
}

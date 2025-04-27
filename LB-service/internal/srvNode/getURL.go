package srvnode

import "net/url"

func (b *SrvNode) GetURL() *url.URL {
	return b.URL
}

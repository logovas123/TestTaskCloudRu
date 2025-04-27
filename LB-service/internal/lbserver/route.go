package lbserver

import (
	"net/http"

	"lb-service/internal/middleware"
	serverpool "lb-service/internal/serverPool"
)

func NewRouter(pool *serverpool.ServerPool) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/", pool.LBHandler)

	mux := middleware.AccessLog(r)
	mux = middleware.Panic(mux)

	return mux
}

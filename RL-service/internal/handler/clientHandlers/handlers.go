package clientHandlers

import (
	"net/http"

	"rl-service/internal/middleware"
)

type ClientHandlers struct{}

type Handlers interface {
	GetReq(w http.ResponseWriter, r *http.Request)
}

func NewClientHandlers() *ClientHandlers {
	return &ClientHandlers{}
}

func MapClientHandlers(h Handlers, mw *middleware.MDWManager) http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", h.GetReq)

	mux := mw.RateLimit(r)
	mux = middleware.AccessLog(mux)
	mux = middleware.Panic(mux)

	return mux
}

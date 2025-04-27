package clientHandlers

import (
	"fmt"
	"net/http"
)

func (h *ClientHandlers) GetReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "request received success")
}

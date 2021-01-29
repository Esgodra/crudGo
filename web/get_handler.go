package web

import (
	"crudGo/gateway"
	"crudGo/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *GetBookHandler) GetBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := h.Get()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in get book"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

type GetBookHandler struct {
	gateway.BookGetGateway
}

func NewGetBookHandler(client *database.DB2Client) *GetBookHandler {
	fmt.Printf("Handler")
	return &GetBookHandler{BookGetGateway: gateway.NewBookGetGateway(client)}
}

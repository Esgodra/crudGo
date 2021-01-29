package web

import (
	"crudGo/gateway"
	"crudGo/internal/database"
	"crudGo/models"
	"encoding/json"
	"net/http"
)

func (h *CreateBookHandler) SaveBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := parserRequest(r)

	res, err := h.Create(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create book"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

type CreateBookHandler struct {
	gateway.BookCreateGateway
}

func NewCreateBookHandler(client *database.DB2Client) *CreateBookHandler {
	return &CreateBookHandler{BookCreateGateway: gateway.NewBookCreateGateway(client)}
}

func parserRequest(r *http.Request) *models.CreateBookCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateBookCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}

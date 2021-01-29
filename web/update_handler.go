package web

import (
	"crudGo/gateway"
	"crudGo/internal/database"
	"crudGo/models"
	"encoding/json"
	"net/http"
)

func (h *UpdateBookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := updateParserRequest(r)

	res, err := h.Update(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in update book"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

type UpdateBookHandler struct {
	gateway.BookUpdateGateway
}

func NewUpdateBookHandler(client *database.DB2Client) *UpdateBookHandler {
	return &UpdateBookHandler{BookUpdateGateway: gateway.NewBookUpdateGateway(client)}
}

func updateParserRequest(r *http.Request) *models.UpdateBookCMD {
	body := r.Body
	defer body.Close()
	var cmd models.UpdateBookCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}

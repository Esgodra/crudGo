package web

import (
	"crudGo/gateway"
	"crudGo/internal/database"
	"crudGo/models"
	"encoding/json"
	"net/http"
)

func (h *DeleteBookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := deleteParserRequest(r)

	res, err := h.Delete(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in delete book"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

type DeleteBookHandler struct {
	gateway.BookDeleteGateway
}

func NewDeleteBookHandler(client *database.DB2Client) *DeleteBookHandler {
	return &DeleteBookHandler{BookDeleteGateway: gateway.NewBookDeleteGateway(client)}
}

func deleteParserRequest(r *http.Request) *models.DeleteBookCMD {
	body := r.Body
	defer body.Close()
	var cmd models.DeleteBookCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}

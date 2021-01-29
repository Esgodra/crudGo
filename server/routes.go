package server

import (
	"crudGo/web"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//Routes method
func Routes(c *web.CreateBookHandler, u *web.UpdateBookHandler, d *web.DeleteBookHandler, g *web.GetBookHandler) *chi.Mux {
	fmt.Println("Route")
	mux := chi.NewMux()

	//globals middlewares
	mux.Use(
		middleware.Logger,    //log every request
		middleware.Recoverer, //recover if a panic occurs
	)

	mux.Post("/insertBook", c.SaveBookHandler)
	mux.Put("/updateBook", u.UpdateBookHandler)
	mux.Delete("/deleteBook", d.DeleteBookHandler)
	mux.Get("/books", g.GetBookHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "test")

	res := map[string]interface{}{"message": "Hello world"}

	_ = json.NewEncoder(w).Encode(res)
}

package api

import (
	"net/http"

	"github.com/Kianelc/ama-room.git/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
)

//estamos adicionando o tipo concreto e adicionando aqui o ideal seria colocar como uma interface
type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NweHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler {
		q: q,
	}

	r := chi.NewRouter()

	a.r = r
	return a
}
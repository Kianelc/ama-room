package api

import (
	"net/http"

	"github.com/Kianelc/ama-room.git/internal/store/pgstore"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	//midleware CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//websocket
	r.Get("/subscribe/{room_id}", a.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)

			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", a.handleCreateRoomMessage)
				r.Get("/", a.handleGetRoomMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", a.handleGetRoomMessage)
					r.Patch("/react", a.handleReactToMessage)
					r.Delete("/react", a.handleRemoveReactFromMessage)
					r.Patch("/answer", a.handleMarkMessageAsAnswered)
				})
			})
		})
	})

	a.r = r
	return a
}

func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleReactToMessage(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleRemoveReactFromMessage(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request){}
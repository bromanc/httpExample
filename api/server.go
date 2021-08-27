package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type Item struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router
	items []Item
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		items:  []Item{},
	}
	s.route()
	return s
}

func (s *Server) route() {
	s.HandleFunc("/items", s.listItems()).Methods("GET")
	s.HandleFunc("/items", s.createItem()).Methods("POST")
	s.HandleFunc("/items", s.removeItem()).Methods("DELETE")
}

func (s *Server) createItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.Id = uuid.New()
		s.items = append(s.items, i)
		
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.items {
			if item.Id == id {
				s.items = append(s.items[:i], s.items[i+1:]...)
				break
			}
		}
	}
}

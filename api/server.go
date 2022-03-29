package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Items struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Status struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

type Server struct {
	*mux.Router
	shoppingList []Items
}

func NewServer() *Server {
	s := &Server{
		Router:       mux.NewRouter(),
		shoppingList: []Items{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/shopping-list", s.getShopingItem()).Methods("GET")
	s.HandleFunc("/shopping-list", s.createShopingItem()).Methods("POST")
	s.HandleFunc("/shopping-list/{id}", s.getShopingItemByID()).Methods("GET")
	s.HandleFunc("/shopping-list/{id}", s.removeShoppingItem()).Methods("DELETE")
	s.HandleFunc("/status", s.ServerWorking()).Methods("GET")
}

func (s *Server) ServerWorking() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dt := time.Now()
		w.Header().Set("Content-Type", "application/json")
		data := Status{Message: "Server is working", Time: dt.String()}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s *Server) createShopingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item Items
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item.ID = uuid.New()
		s.shoppingList = append(s.shoppingList, item)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s *Server) getShopingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingList); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s *Server) getShopingItemByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, item := range s.shoppingList {
			if item.ID.String() == params["id"] {
				if err := json.NewEncoder(w).Encode(item); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}
		}
	}
}

func (s *Server) removeShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range s.shoppingList {
			if item.ID.String() == params["id"] {
				s.shoppingList = append(s.shoppingList[:index], s.shoppingList[index+1:]...)
				break
			}
		}
		if err := json.NewEncoder(w).Encode(s.shoppingList); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Siddhant6674/ECOM/service/user"

	"github.com/gorilla/mux"
)

type APIserver struct {
	Address string
	db      *sql.DB
}

func NewAPIserver(Address string, db *sql.DB) *APIserver {
	return &APIserver{
		Address: Address,
		db:      db,
	}
}

func (s *APIserver) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Server running on port", s.Address)
	return http.ListenAndServe(s.Address, router)

}
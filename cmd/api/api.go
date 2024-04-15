// Package api bertanggung jawab untuk menangani logika utama aplikasi web API.
package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/arif-rizal1122/golang-jwt/service/users"
	"github.com/gorilla/mux"
)

// APIServer adalah struktur yang menyimpan konfigurasi server API.
type APIServer struct {
	addr string
	db   *sql.DB
}

// NewAPIServer mengembalikan instance baru dari APIServer.
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run menjalankan server API dengan konfigurasi yang telah ditentukan.
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	// Membuat handler user dan mendaftarkannya ke subrouter.
	userHandler := user.NewHandler(userStore)
	
	userHandler.RegisterRoutes(subrouter)
	// Memulai server dengan alamat dan router yang telah disiapkan.
	log.Println("listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

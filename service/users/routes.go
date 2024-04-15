// Package user berisi logika terkait user seperti registrasi dan login.
package user

import (
	"net/http"

	"github.com/arif-rizal1122/golang-jwt/types"
	"github.com/arif-rizal1122/golang-jwt/utils"
	"github.com/gorilla/mux"
)

// Handler adalah struktur yang menyimpan logika terkait user.
type Handler struct {
   store *types.UserStore
}

// NewHandler mengembalikan instance baru dari Handler.
func NewHandler(store *types.UserStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes mendaftarkan rute-rute yang terkait user ke router yang diberikan.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

// handleLogin adalah fungsi yang menangani permintaan login.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Implementasi dari handleLogin akan ditambahkan di sini.
}

// handleRegister adalah fungsi yang menangani permintaan registrasi.
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// ambil json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// cek jika user sudah ada
	h.store.GetUserByEmail()
	// jika belum ada kita buat user baru

    
}

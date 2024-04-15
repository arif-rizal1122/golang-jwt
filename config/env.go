package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config merupakan struktur data yang menyimpan konfigurasi aplikasi.
// Konfigurasi ini biasanya berasal dari variabel lingkungan (environment variables).
type Config struct {
	PublicHost string // Alamat publik aplikasi, contoh: "http://localhost"
	Port       string // Port aplikasi, contoh: "8080"
	DBUser     string // Pengguna database, contoh: "root"
	DBPassword string // Kata sandi database, contoh: ""
	DBAddress  string // Alamat dan port database, contoh: "127.0.0.1:3306"
	DBName     string // Nama database, contoh: "go-rest-ecommerce"
}

// Envs adalah variabel global yang menyimpan konfigurasi aplikasi.
var Envs = initConfig()

// initConfig adalah fungsi yang akan dijalankan saat package diinisialisasi.
// Fungsi ini membaca konfigurasi dari variabel lingkungan dan mengembalikan
// struktur Config yang sesuai.
func initConfig() Config {
	// Memuat file .env ke variabel lingkungan
	godotenv.Load()

	// Membuat dan mengembalikan struktur Config dengan konfigurasi yang sesuai
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),         // Mengambil alamat publik atau default "http://localhost"
		Port:       getEnv("PORT", "8080"),                             // Mengambil port atau default "8080"
		DBUser:     getEnv("DB_USER", "root"),                          // Mengambil pengguna database atau default "root"
		DBPassword: getEnv("DB_PASSWORD", ""),                          // Mengambil kata sandi database atau default kosong
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), // Menggabungkan alamat dan port database
			getEnv("DB_PORT", "3306")),                                  // Mengambil port database atau default "3306"
		DBName: getEnv("DB_NAME", "go-rest-ecommerce"),                 // Mengambil nama database atau default "go-rest-ecommerce"
	}
}

// getEnv adalah fungsi utilitas untuk mendapatkan nilai dari variabel lingkungan.
// Jika variabel lingkungan tidak ditemukan, akan mengembalikan nilai fallback.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok { // Mencari nilai variabel lingkungan dengan kunci tertentu
		return value
	}

	return fallback // Mengembalikan nilai fallback jika variabel lingkungan tidak ditemukan
}

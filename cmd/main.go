// Package main adalah entry point dari aplikasi.
package main

import (
	"database/sql"
	"log"

	"github.com/arif-rizal1122/golang-jwt/cmd/api"
	"github.com/arif-rizal1122/golang-jwt/config"
	"github.com/arif-rizal1122/golang-jwt/db"
	"github.com/arif-rizal1122/golang-jwt/helper"
	"github.com/go-sql-driver/mysql"
)

// main adalah fungsi utama yang akan dijalankan pertama kali saat aplikasi dimulai.
func main() {

	db, err := db.NewMysqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	helper.IfLogError(err)

	initStorage(db)

	// Membuat instance baru dari APIServer dengan konfigurasi tertentu.
	server := api.NewAPIServer(":8080", db)

	// Menjalankan server dan menangani kesalahan jika ada.
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}



func initStorage(db *sql.DB)  {
	err := db.Ping()
	helper.IfLogError(err)

	log.Println("DB: succesfully connected")
}

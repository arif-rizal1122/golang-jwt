package db

import (
	"database/sql"

	"github.com/arif-rizal1122/golang-jwt/helper"
	"github.com/go-sql-driver/mysql"
)

func NewMysqlStorage(cfg mysql.Config) (*sql.DB, error) {
   db, err := sql.Open("mysql", cfg.FormatDSN())
   helper.IfLogError(err)

   return db, nil
}
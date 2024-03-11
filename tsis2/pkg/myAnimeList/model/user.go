package model

import (
	"database/sql"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	// CreatedAt string `json:"createdAt"` // DD-MM-YYYY
}

type UserModel struct {
	DB       *sql.DB
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

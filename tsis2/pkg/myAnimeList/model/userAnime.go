package model

import (
	"database/sql"
	"log"
)

type UserAnime struct {
	UserID  int    `json:"userId"`  // Foreign key referencing User table
	AnimeID int    `json:"animeId"` // Foreign key referencing Anime table
	Status  string `json:"status"`  // "watching", "completed", "on hold"
	// Rating  float64 `json:"rating"`  // (User rating on that anime /10)
}

type UserAnimeModel struct {
	DB       *sql.DB
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

// Methods for managing user-anime relationships (e.g., AddUserAnime, GetUserAnimes)

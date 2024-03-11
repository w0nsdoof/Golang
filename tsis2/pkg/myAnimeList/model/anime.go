package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Anime struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"releaseDate"` // DD-MM-YYYY
	Status      string  `json:"status"`      // Finished | Airing | Announced
	Type        string  `json:"type"`        // film | TV serial | etc...
	Genre       string  `json:"genre"`
	Studio      string  `json:"studio"`
	Cover       string  `json:"cover"`
	Episodes    int     `json:"episodes"` // Number of episodes
	Duration    int     `json:"duration"` // avg duration of series
	Rating      float64 `json:"rating"`   // MAL score

	// SumRating float64 `json:"sumRating"` // for userAnime model
	// CntRating float64 `json:"cntRating"`	// for userAnime model
}

type AnimeModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m AnimeModel) Insert(anime *Anime) error {
	query := `INSERT INTO animes(title,releaseDate,status,type,genre,studio,cover,episodes,duration)
			  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
			  RETURNING id`

	args := []interface{}{anime.Title, anime.ReleaseDate, anime.Status, anime.Type, anime.Genre, anime.Studio, anime.Cover, anime.Episodes, anime.Duration}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&anime.Id)
}

func (m AnimeModel) Get(id int) (*Anime, error) {
	query := `SELECT * FROM animes
			  WHERE id = $1
			  `

	var anime Anime
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&anime.Id, &anime.Title, &anime.ReleaseDate, &anime.Status, &anime.Type, &anime.Genre, &anime.Studio, &anime.Cover, &anime.Episodes, &anime.Duration)

	if err != nil {
		return nil, err
	}
	return &anime, nil
}

func (m AnimeModel) Update(anime *Anime) error {
	query := `
	  UPDATE animes
	  SET title = $1, release_date = $2, status = $3, type = $4, genre = $5, studio = $6, cover = $7, episodes = $8, duration = $9
	  WHERE id = $10
	  RETURNING updated_at
	`
	args := []interface{}{anime.Title, anime.ReleaseDate, anime.Status, anime.Type, anime.Genre, anime.Studio, anime.Cover, anime.Episodes, anime.Duration, anime.Id} // Order matters!
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

func (m AnimeModel) Delete(id int) error {
	query := `
		DELETE FROM animes
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}

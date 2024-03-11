package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tsis2/pkg/myAnimeList/model"

	"github.com/gorilla/mux"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createAnimeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string  `json:"title"`
		ReleaseDate string  `json:"releaseDate"`
		Status      string  `json:"status"`
		Type        string  `json:"type"`
		Genre       string  `json:"genre"`
		Studio      string  `json:"studio"`
		Cover       string  `json:"cover"`
		Episodes    int     `json:"episodes"`
		Duration    int     `json:"duration"`
		Rating      float64 `json:"rating"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	anime := &model.Anime{
		Title:       input.Title,
		ReleaseDate: input.ReleaseDate,
		Status:      input.Status,
		Type:        input.Type,
		Genre:       input.Genre,
		Studio:      input.Studio,
		Cover:       input.Cover,
		Episodes:    input.Episodes,
		Duration:    input.Duration,
		Rating:      input.Rating,
	}

	err = app.models.Animes.Insert(anime)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}
	app.respondWithJSON(w, http.StatusCreated, anime)
}

func (app *application) getAnimeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid anime id")
		return
	}

	anime, err := app.models.Animes.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, anime)
}

func (app *application) updateAnimeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid anime id")
		return
	}

	anime, err := app.models.Animes.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Title       *string  `json:"title"`
		ReleaseDate *string  `json:"releaseDate"`
		Status      *string  `json:"status"`
		Type        *string  `json:"type"`
		Genre       *string  `json:"genre"`
		Studio      *string  `json:"studio"`
		Cover       *string  `json:"cover"`
		Episodes    *int     `json:"episodes"`
		Duration    *int     `json:"duration"`
		Rating      *float64 `json:"rating"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.Title != nil {
		anime.Title = *input.Title
	}
	if input.ReleaseDate != nil {
		anime.ReleaseDate = *input.ReleaseDate
	}
	if input.Status != nil {
		anime.Status = *input.Status
	}
	if input.Type != nil {
		anime.Type = *input.Type
	}
	if input.Genre != nil {
		anime.Genre = *input.Genre
	}
	if input.Studio != nil {
		anime.Studio = *input.Studio
	}
	if input.Cover != nil {
		anime.Cover = *input.Cover
	}
	if input.Episodes != nil {
		anime.Episodes = *input.Episodes
	}
	if input.Duration != nil {
		anime.Duration = *input.Duration
	}
	if input.Rating != nil {
		anime.Rating = *input.Rating
	}

	err = app.models.Animes.Update(anime)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, anime)
}

func (app *application) deleteAnimeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid anime ID")
		return
	}

	err = app.models.Animes.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return err
	}

	return nil
}

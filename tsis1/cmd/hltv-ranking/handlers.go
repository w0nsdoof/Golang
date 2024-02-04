package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	hltv_ranking "awesomeProject/pkg/hltv-ranking"
	"github.com/gorilla/mux"
)

type Response struct {
	Players []hltv_ranking.Player `json:"players"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": hltv_ranking.Info()})
}

func Players(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, hltv_ranking.GetPlayers())
}

func Player(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["player"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "400 Bad request")
		return
	}
	player, err := hltv_ranking.GetPlayer(i)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 Not found")
		return
	}
	respondWithJSON(w, http.StatusOK, player)
}

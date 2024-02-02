package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Player struct {
	Rank     int     `json:"rank"`
	Nickname string  `json:"nickname"`
	Rating   float32 `json:"rating"`
	MVPs     int     `json:"mvps"`
	EVPs     int     `json:"evps"`
	Image    string  `json:"image"`
}

var players = []Player{
	{1, "ZywOo", 1.31, 5, 5, "https://img-cdn.hltv.org/playerbodyshot/RoTpFBevmVDbeCBJLSUV8o.png?ixlib=java-2.1.0&w=400&s=6d65cb68c0cbfbfe0f14d16971615c4b"},
	{2, "NiKo", 1.21, 0, 3, "https://img-cdn.hltv.org/playerbodyshot/eefPmsiMIo4dAiiPUmZE6-.png?ixlib=java-2.1.0&w=400&s=8d9765f9d2c40c13d7bd6c96c45a2849"},
	{3, "ropz", 1.17, 3, 4, "https://img-cdn.hltv.org/teamlogo/m-wCSia4m35BeWhGzB9yv1.png?ixlib=java-2.1.0&w=400&s=1872dd13f27d231bed497445d734363d"},
	{4, "m0NESY", 1.19, 0, 4, "https://img-cdn.hltv.org/playerbodyshot/F8-kMfRKbV5UDnMo4elY5J.png?ixlib=java-2.1.0&w=400&s=8b740a6d3e55871a8d773b298b486e8c"},
	{5, "Spinx", 1.17, 0, 6, "https://img-cdn.hltv.org/playerbodyshot/tLkyuJMkGunXQD7-CAKsFE.png?ixlib=java-2.1.0&w=400&s=a605fe434f8a458fa8c1d20e82e49eaf"},
}

func ListPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	params := mux.Vars(r)
	playerRank := params["rank"]

	rank, err := strconv.Atoi(playerRank)
	if err != nil {
		http.Error(w, "Invalid player rank", http.StatusBadRequest)
		return
	}

	for _, player := range players {
		if player.Rank == rank {
			json.NewEncoder(w).Encode(player)
			return
		}
	}

	http.NotFound(w, r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This app lists top-5 of CSGO players 2023 (HLTV)"))
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/players", ListPlayers).Methods("GET")
	r.HandleFunc("/players/{rank}", GetPlayer).Methods("GET")
	r.HandleFunc("/health", HealthCheck).Methods("GET")
}

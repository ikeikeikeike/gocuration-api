package main

import (
	"encoding/json"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {
	goji.Get("/divas", Divas)
	goji.Get("/animes", Animes)
	goji.Get("/characters", Characters)
	goji.NotFound(NotFound)
	goji.Serve()
}

func Divas(c web.C, w http.ResponseWriter, r *http.Request) {
	var list []Diva
	DB.Order("id", true).Find(&list)

	encoder := json.NewEncoder(w)
	encoder.Encode(list)
}

func Animes(c web.C, w http.ResponseWriter, r *http.Request) {
	var list []Anime
	DB.Order("id", true).Find(&list)

	encoder := json.NewEncoder(w)
	encoder.Encode(list)
}

func Characters(c web.C, w http.ResponseWriter, r *http.Request) {
	var list []Character
	DB.Order("id", true).Find(&list)

	encoder := json.NewEncoder(w)
	encoder.Encode(list)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", 404)
}

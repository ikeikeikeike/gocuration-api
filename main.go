package main

import (
	"encoding/json"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {
	goji.Use(SuperSecure)

	goji.Get("/divas", Divas)
	goji.Get("/animes", Animes)
	goji.Get("/characters", Characters)

	goji.NotFound(NotFound)
	goji.Serve()
}

func Divas(c web.C, w http.ResponseWriter, r *http.Request) {
	var list []Diva
	DB.Order("id", true).
		Preload("Icon").
		Where("height > ?", 0).
		Where("bracup != ''").
		Limit(10000).
		Find(&list)

	encoder := json.NewEncoder(w)
	encoder.Encode(list)
}

func Animes(c web.C, w http.ResponseWriter, r *http.Request) {
	var list []Anime
	DB.Order("id", true).
		Preload("Characters").Preload("Icon").
		Limit(100).
		Find(&list)

	encoder := json.NewEncoder(w)
	encoder.Encode(list)
}

func Characters(c web.C, w http.ResponseWriter, r *http.Request) {
	var list []Character
	DB.Order("id", true).
		Preload("Anime").Preload("Icon").
		Limit(100).
		Find(&list)

	encoder := json.NewEncoder(w)
	encoder.Encode(list)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", 404)
}

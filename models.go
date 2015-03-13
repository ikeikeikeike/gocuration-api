package main

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Diva struct {
	Id          int64
	Name        string
	Kana        string
	Romaji      string
	Gyou        string
	Birthday    time.Time
	Blood       string
	Height      int
	Weight      int
	Bust        int
	Waste       int
	Hip         int
	Bracup      string
	Outline     string
	Created     time.Time
	Updated     time.Time
	Icon        int64
	VideosCount int
}

type Anime struct {
	Id            int64
	Name          string
	Alias         string
	Kana          string
	Romaji        string
	Gyou          string
	Url           string
	Author        string
	Works         string
	ReleaseDate   time.Time
	Outline       string
	Created       time.Time
	Updated       time.Time
	Icon          int64
	PicturesCount int
}

type Character struct {
	Id            int64
	Name          string
	Kana          string
	Romaji        string
	Gyou          string
	Birthday      time.Time
	Blood         string
	Height        int
	Weight        int
	Bust          int
	Waste         int
	Hip           int
	Bracup        string
	Outline       string
	Created       time.Time
	Updated       time.Time
	Icon          int64
	Product       string
	Anime         int64
	PicturesCount int
}

var DB gorm.DB

func init() {
	var err error

	DB, err = gorm.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}
	DB.DB()
	DB.LogMode(true)

	DB.DB().Ping()
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.SingularTable(true)
}

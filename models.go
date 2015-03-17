package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Image struct {
	ID      int64
	Name    string
	Src     string
	Ext     string
	Mime    string
	Width   int
	Height  int
	Created time.Time
	Updated time.Time
}

type Diva struct {
	ID          int64
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
	Icon        Image
	IconID      sql.NullInt64
	VideosCount int
}

type Anime struct {
	ID            int64
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
	Icon          Image
	IconID        sql.NullInt64
	Characters    []Character
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
	Icon          Image
	IconID        sql.NullInt64
	Product       string
	Anime         Anime
	AnimeID       sql.NullInt64
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

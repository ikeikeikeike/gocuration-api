package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Entry struct {
	Id          int64
	Url         string
	Title       string
	Content     string
	SeoTitle    string
	SeoContent  string
	Encoded     string
	Creator     string
	Publisher   string
	PublishedAt time.Time

	Q string

	Created time.Time
	Updated time.Time

	Blog   *Blog
	BlogId sql.NullInt64

	// Video   *Video
	Picture *Picture
	// Summary *Summary

	// Tags   []Tag   `gorm:"many2many:entry_tag;"`
	// Images []Image `gorm:"many2many:entry_image;"`

	// Scores []Score
}

type Blog struct {
	Id          int64
	Rss         string
	Url         string
	Name        string
	Mediatype   string
	Adsensetype string

	VerifyLink  int
	VerifyRss   int
	VerifyParts int

	IsPenalty bool

	LastModified time.Time

	Created time.Time
	Updated time.Time

	// User   User
	// UserId sql.NullInt64

	Icon   Image
	IconId sql.NullInt64

	// Scores  []Score
	Entries []Entry
}

type Picture struct {
	Id int64

	Created time.Time
	Updated time.Time

	Entry   Entry
	EntryId sql.NullInt64

	Anime   Anime
	AnimeId sql.NullInt64

	Images     []Image
	Characters []Character `gorm:"many2many:picture_character;"`
}

type Diva struct {
	Id int64

	Name   string
	Kana   string
	Romaji string
	Gyou   string

	Birthday time.Time

	Blood string

	Height int
	Weight int

	Bust   int
	Waste  int
	Hip    int
	Bracup string

	Outline string

	Created time.Time
	Updated time.Time

	Icon   Image
	IconId sql.NullInt64

	VideosCount int
	// Videos      []Video
}

type Anime struct {
	Id          int64
	Name        string
	Alias       string
	Kana        string
	Romaji      string
	Gyou        string
	Url         string
	Author      string
	Works       string
	ReleaseDate time.Time
	Outline     string
	Created     time.Time
	Updated     time.Time

	Icon   Image
	IconId sql.NullInt64

	Characters []Character

	Pictures      []Picture
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
	IconId        sql.NullInt64
	Product       string
	Anime         Anime
	AnimeId       sql.NullInt64
	PicturesCount int
}

type Image struct {
	Id int64

	Name string
	Src  string

	Ext    string
	Mime   string
	Width  int
	Height int

	Created time.Time
	Updated time.Time

	Picture   *Picture
	PictureId sql.NullInt64
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

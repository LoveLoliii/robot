package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
type Song struct {
	Title  string
	Pic    string
	Singer string
	Score  string
	gorm.Model
}

func getSong(t string) Song {
	db, err := gorm.Open("sqlite3", "song.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Song{})
	var ss Song
	//db.First(&product, 1)
	db.First(&ss, "title = ?", t)
	return ss
}
func addSong(s Song) {
	db, err := gorm.Open("sqlite3", "song.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Song{})

	db.Create(&Song{Title: s.Title, Pic: s.Pic,
		Singer: s.Singer, Score: s.Score})

	var ss Song
	//db.First(&product, 1)
	db.First(&ss, "title = ?", "支え")
	fmt.Printf(ss.Singer)
	//db.Model(&product).Update("Price", 2000)

	//db.Delete(&product)

}

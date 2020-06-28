package api

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
	title  string
	pic    string
	singer string
	score  string
}

func addSong(s Song) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Song{})

	db.Create(&Song{title: s.title, pic: s.pic,
		singer: s.singer, score: s.score})

	var ss Song
	//db.First(&product, 1)
	db.First(&ss, "title = ?", "支え")
	fmt.Printf(ss.singer)
	//db.Model(&product).Update("Price", 2000)

	//db.Delete(&product)

}

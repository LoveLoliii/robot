package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"

	//gorm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	//redis
	"context"

	"github.com/go-redis/redis/v8"
)

const maxSize = 5 << 20 //5MB
// test new ssh key
// type Song struct {
// 	title  string
// 	pic    string
// 	singer string
// 	score  string
// }

func main() {
	//initData()
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	//app.Use(myMiddleware)
	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hi Sayari !")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})
	app.Post("/addSong", func(ctx iris.Context) {
		title := ctx.PostValue("title")
		pic := ctx.PostValue("pic")
		singer := ctx.PostValue("singer")
		score := ctx.PostValue("score")
		issue := ctx.PostValue("issue")
		fmt.Printf("成功取得参数\ntitle:%s\npic:%s\nsinger:%s\nscore:%s\nissue:%s",
			title, pic, singer, score, issue)

		addSong(Song{Title: title, Pic: pic, Singer: singer, Score: score, Issue: issue})
	})
	// use song name
	app.Get("/getSong/{title:string}", func(ctx iris.Context) {
		title := ctx.Params().GetString("title")
		// gorm get data
		var s Song
		s = getSong(title)
		ctx.Writef("title:%s\npic:%s\nsinger:%s\nscore:%s\nissue:%s", s.Title, s.Pic, s.Singer, s.Score, s.Issue)
	})
	// query by song name
	app.Post("/querySong/{searchKey:string}", func(ctx iris.Context) {
		searchKey := ctx.Params().GetString("searchKey")
		fmt.Printf("%s", searchKey)
		// query titles from \redis server\ sqlite
		//rdb.Get(ctxx,"songs")
		var ss []Song
		ss = querySong(searchKey)
		for k, v := range ss {
			fmt.Printf("K:%d,v:%s", k, v.Title)
			ctx.Writef("title:%s\npic:%s\nsinger:%s\nscore:%s\nissue:%s", v.Title, v.Pic, v.Singer, v.Score, v.Issue)
		}
	})

	// query on web
	app.Get("/querySong", func(ctx iris.Context) {
		ctx.View("query_song.html")

	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from the SECURE server on path /mypath")
	})
	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
	app.HandleDir("/static", "./views")
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})
	app.Get("/search/{search:string}", func(ctx iris.Context) {
		searchKey := ctx.Params().GetString("search")
		ctx.Writef("not ready to return result by key :%s", searchKey)
	})

	// Serve the upload_form.html to the client.
	app.Get("/upload", func(ctx iris.Context) {
		// create a token (optionally).

		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		// render the form with the token for any use you'd like.
		// ctx.ViewData("", token)
		// or add second argument to the `View` method.
		// Token will be passed as {{.}} in the template.
		ctx.View("upload_form.html", token)
	})

	app.Post("/upload", iris.LimitRequestBodySize(maxSize), func(ctx iris.Context) {
		fmt.Printf("upload")
		// Get the file from the request.
		file, info, err := ctx.FormFile("uploadfile")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}

		defer file.Close()
		fname := info.Filename

		// Create a file with the same name
		// assuming that you have a folder named 'uploads'
		out, err := os.OpenFile("./views/upload/"+fname,
			os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer out.Close()

		io.Copy(out, file)
		path := "https://sayarii.art/static/upload/" + fname
		ctx.Writef("file path is -->:\n%s", path)
	})
	// 使用nginx监听80/443 再转发到8888
	app.Listen(":8888")
}
func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	// ip := ctx.RemoteAddr()
	// ip = strings.Replace(ip, ".", "_", -1)
	// ip = strings.Replace(ip, ":", "_", -1)
	//file.Filename = ip + "-" + file.Filename
	//fmt.Printf("fileNmae:%s", file.Filename)
}
func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

//gorm
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
	Issue  string
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
func querySong(t string) []Song {
	db, err := gorm.Open("sqlite3", "song.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Song{})
	var ss []Song
	db.Where("title LIKE ?", "%"+t+"%").Find(&ss)
	return ss
}
func addSong(s Song) {
	db, err := gorm.Open("sqlite3", "song.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Song{})
	var sr Song
	db.Where("issue =?", s.Issue).First(&sr)
	fmt.Printf(sr.Issue)
	if sr.Issue == "" {
		db.Create(&Song{Title: s.Title, Pic: s.Pic,
			Singer: s.Singer, Score: s.Score, Issue: s.Issue})
		fmt.Printf("add success")
	} else {
		fmt.Printf("add failed")
	}

	// var ss Song
	//db.First(&product, 1)
	// db.First(&ss, "title = ?", "支え")
	// fmt.Printf(ss.Singer)
	//db.Model(&product).Update("Price", 2000)

	//db.Delete(&product)

}

// do init
func initData() {
	// try load tilte to redis server

	//exampleNewRedisClient()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no pwd set
		DB:       0,        //use default db
	})

	pong, err := rdb.Ping(ctxx).Result()
	fmt.Println(pong, err)
	val, err1 := rdb.Get(ctxx, "songs").Result()
	if err != nil {
		panic(err1)
	}
	fmt.Println("key", val)
	db, err := gorm.Open("sqlite3", "song.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Song{})
	var sss []Song
	//db.First(&product, 1)
	rs := db.Find(&sss).Error
	fmt.Println(rs)
	fmt.Println(sss[1].Title)
	// write to redis server
	t1 := time.Now().UnixNano() / 1e6
	for idx, song := range sss {

		fmt.Println(song.Title)
		// ZSet 无法存入 未知原因 改用list list会重复 使用set
		// val, err1 := rdb.ZInterStore(ctxx, "songs", &redis.ZStore{
		// 	Keys:    []string{song.Title},
		// 	Weights: []float64{1.0}}).Result()
		// if err != nil {
		// 	panic(err1)
		// }
		// fmt.Println("key", val)
		//rdb.Command()
		//sort := int64(idx)
		// t := rdb.Get(ctxx, "songs")

		// var flag bool
		// flag = true
		// for s, i := range t {
		// 	if s == song.Title {
		// 		flag = false
		// 	}
		// }
		// if !flag {
		// 	rs := rdb.LPush(ctxx, "songs", song.Title)
		// 	fmt.Println("key", rs)
		// } else {
		// 	fmt.Println("已存在")
		// }

		idx = idx + 1

	}
	t2 := time.Now().UnixNano() / 1e6
	fmt.Printf("%v毫秒", (t2 - t1))

	//readDataFromSqlite3()
}

var ctxx = context.Background()

//var rdb redis.Client

// func exampleNewRedisClient() {

// }
// func readDataFromSqlite3() {

// }

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
)

const maxSize = 5 << 20 //5MB
// test new ssh key
type Song struct {
	title  string
	pic    string
	singer string
	score  string
}

func main() {
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
		fmt.Printf("成功取得参数\ntitle:%s\npic:%s\nsinger:%s\nscore:%s",
			title, pic, singer, score)

		module.api.addSong(&Song{title: title, pic: pic, singer: singer, score: score})
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

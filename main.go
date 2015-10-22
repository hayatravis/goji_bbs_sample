package main

import (
	"fmt"
	"io"
	"net/http"
// "encoding/json"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/elgs/gosqljson"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

//	"github.com/goji/param"
	"reflect"
)


func main() {
	/* -- Router -- */
	// Top
	goji.Get("/", Root)
	// User list
	goji.Get("/users/:offset/:limit", Users)

	// Comment List
	goji.Get("/comment_list/:offset/:limit", CommentList)

	// Comment
	goji.Post("/coment", Comment)

	goji.NotFound(NotFound)
	goji.Serve()
}

func Root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Golang Goji.")
}

func Users(c web.C, w http.ResponseWriter, r *http.Request) {
	/* -- DB -- */
	db, err := sql.Open("mysql", "root:@/test_bbs")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/* -- Params -- */
	limit := c.URLParams["limit"]
	offset := c.URLParams["offset"]

	/* -- SQL -- */
	sql := "SELECT id, name FROM user LIMIT ?, ?;"
	theCase := "lower"
	data, _ := gosqljson.QueryDbToMapJson(db, theCase, sql, offset, limit)

	/* -- View -- */
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, data)
}

func CommentList(c web.C, w http.ResponseWriter, r *http.Request) {
	/* -- DB -- */
	db, err := sql.Open("mysql", "root:@/test_bbs")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/* -- Params -- */
	limit := c.URLParams["limit"]
	offset := c.URLParams["offset"]

	/* -- SQL -- */
	sql := "SELECT * FROM comment LIMIT ?, ?;"
	theCase := "lower"
	data, _ := gosqljson.QueryDbToMapJson(db, theCase, sql, offset, limit)

	/* -- View -- */
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, data)
}

func Comment(w http.ResponseWriter, r *http.Request) {
	/* -- DB -- */
	db, err := sql.Open("mysql", "root:@/test_bbs")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/* -- Params -- */
	r.ParseForm()
	post := r.Form
	name := string(post["name"])
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(post["comment"])
	fmt.Println(post["tags"])
	fmt.Println(post["area"])


	/* -- SQL -- */
	sql := "INSERT INTO comment (name, comment, tags, area, create, modiefied)  VALUES (?, ?, ?, ?, NOW(), NOW());"
	result, err := db.Exec(sql, post["name"], post["comment"], post["tags"], post["area"])
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println(result)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found...", 404)
}

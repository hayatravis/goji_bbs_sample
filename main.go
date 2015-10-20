package main

import (
	"fmt"
	"io"
	"net/http"
	"database/sql"
//	"encoding/json"
	"github.com/elgs/gosqljson"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	_ "github.com/go-sql-driver/mysql"

//	"github.com/goji/param"
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

	/* view */
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

	/* view */
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, data)
}

func Comment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, "aaaaaaaaaaaaaa")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found...", 404)
}

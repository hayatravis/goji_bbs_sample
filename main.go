package main

import (
	"fmt"
	"io"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/elgs/gosqljson"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"strings"
)


func main() {
	/* -- Router -- */
	// Top
	goji.Get("/", Root)

	// User list
	goji.Get("/user_list/:offset/:limit", UserList)

	// Comment List
	goji.Get("/comment_list/:offset/:limit", CommentList)

	// Comment
	goji.Post("/coment", Comment)

	goji.NotFound(NotFound)
	goji.Serve()
}

/**
 * root page
 */
func Root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Golang Goji.")
}


/**
 * user list
 */
func UserList(c web.C, w http.ResponseWriter, r *http.Request) {
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

/**
 * comment list
 */
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

/**
 * post comment
 */
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
	var name, comment, tags, area string
	name = strings.Join(post["name"], "")
	comment = strings.Join(post["comment"], "")
	tags = strings.Join(post["tags"], "")
	area = strings.Join(post["area"], "")


	/* -- SQL -- */
	sql := "INSERT INTO comment (name, comment, tags, area, created, modified)  VALUES (?, ?, ?, ?, NOW(), NOW());"
	result, err := db.Exec(sql, name, comment, tags, area)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println(result)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found...", 404)
}

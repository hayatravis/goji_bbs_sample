package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)


func main() {
	// Top
	goji.Get("/", Root)
	// Thread list
	goji.Get("/threads/:page", Threads)
	// Thread
	goji.Get("/thread/:id/:page", Thread)

	// Create Thread
	goji.Post("/create_thread", CreateThread)
	// Delete Thread(Admin)
//	goji.Post("/delete_thread/:id", DeleteThread)
	// Post Comment
//	goji.Post("/post_comment/:id", PostComment)
	// Delete Comment(Admin)
//	goji.Post("/delete_comment/:id", DeleteComment)

	goji.NotFound(NotFound)
	goji.Serve()
}

func Root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Golang Goji.")
}

func Threads(c web.C, w http.ResponseWriter, r *http.Request) {
	pid := c.URLParams["page"]
	fmt.Fprintf(w, "Threads: This page is %s!", pid)
}

func Thread(c web.C, w http.ResponseWriter, r *http.Request) {
	pid := c.URLParams["page"]
	tid := c.URLParams["id"]
	fmt.Fprintf(w, "Thread: Thread id is %s and page is %s", tid, pid)
}

func CreateThread(c web.C, w http.ResponseWriter, r *http.Request) {

}
//
//func DeleteThread() {
//
//}
//
//func PostComment() {
//
//}
//
//func DeleteComment() {
//
//}

/**
 *
 */
func User(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	fmt.Printf("%+v", name)
	io.WriteString(w, name)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found...", 404)
}

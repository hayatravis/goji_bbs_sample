package main

import (
	"log"
	"io"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)


func main() {
	goji.Get("/", Root)
	goji.Get("/user/:name", User)
	goji.NotFound(NotFound)
	goji.Serve()
}

func Root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Golang Goji.")
}

func User(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	log.Printf("%+v", name)
	io.WriteString(w, name)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found...", 404)
}

package controllers

import (
    "net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./blog/index.html");
}

func ServeBlog(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./blog/blog.html");
}

func ServeAbout(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./blog/about.html");
}

func ServeMmblog(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./blog/mmblog.html");
}

func ServeNewTest(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./blog/new-test.html");
}
package controllers

import (
    "net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./blog/index.html");
}

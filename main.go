package main

import (
	"net/http"
	"log"
	"fmt"
	"os"

	"github.com/jobutterfly/makeblog/controllers"
	"github.com/jobutterfly/makeblog/new"
)

func main() {
	if len(os.Args) < 2 {
	    log.Fatal("Expected at least two arguments");
	}

	arg := os.Args[1];

	switch (arg){

	case "serve":
	fs := http.FileServer(http.Dir("./blog/static"));
	http.Handle("/static/", http.StripPrefix("/static/", fs));

	http.HandleFunc("/", controllers.ServeIndex);
	http.HandleFunc("/blog", controllers.ServeBlog);
	http.HandleFunc("/about", controllers.ServeAbout);
	http.HandleFunc("/mmblog", controllers.ServeMmblog);

	log.Print("Listening on port :3000");
	err := http.ListenAndServe(":3000", nil);
	if err != nil {
	    log.Fatal(err);
	}

	case "new":
	if len(os.Args) < 3 {
	    log.Fatal("Not enough arguments, expected input file")
	}

	if err := new.New(os.Args[2]); err != nil {
	    log.Fatal(err);
	}
	case "help":
	fmt.Printf(`
Welcome to makeblog! These are the possible commands:

    makeblog serve

This command serves the website on port 3000.

    makeblog new -i <input.md> -o <output.html>

This command is for creating new blog posts. It takes the input markdown file,
parses it into html, places it inside the layout template, and writes into the
output file. If the file does not exist, it creates a new one. The path to the
must be after the -i or --input flag and the same for the output file.
`)

	default:
	fmt.Printf("Not valid argument, try makeblog help");
	}
}






package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// struct type that is responsible for loading,
// compiling, and delivering the template.
// Compile the template once (using the sync.Once type),
// keep the reference to the compiled template, and then respond to HTTP requests.

type templateHandler struct {
	// The sync.Once type guarantees that the function we pass as an argument will
	// only be executed once, regardless of how many goroutines are calling ServeHTTP.
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
/*
The templateHandler type has a single method called ServeHTTP whose signature looks
suspiciously like the method we passed to http.HandleFunc earlier.
This method will load the source , compile the template and execute it, and write
the output to the  http.ResponseWriter object.
Because the ServeHTTP method satisfies the http.Handler interface,
we can actually pass it directly to http.Handle.
*/

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Compiling the template inside the ServeHTTP method also ensures that our
	// code does not waste time doing work before it is needed. lazy initialization
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	//  create and then run a room for everybody to connect to:
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// get the room going
	go r.run()
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

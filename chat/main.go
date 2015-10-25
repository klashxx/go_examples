package main

import (
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc function maps the path pattern "/"
	// to the function we pass as the second argument,
	// so when the user hits http://localhost:8080/,
	// the function will be executed.
	// The function signature of func(w http. ResponseWriter, r *http.Request)
	// is a common way of handling HTTP requests throughout the Go standard library.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
         <html>
           <head>
             <title>Chat</title>
         </head>
<body>
Let's chat!
</body>
         </html>
`))
	})
	// Listen to the root path using the net/http package
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

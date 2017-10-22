package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Logger is a cool type
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

func serveIndex(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(w, req, "static/index.html")
}

// HelloWorld is a cool function
func HelloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Simply write some test data for now
	fmt.Fprintln(w, "Hello World!")
}

// App is a cool function
func App() http.Handler {
	router := httprouter.New()

	// Add a handler on /hello
	router.GET("/hello", HelloWorld)

	// Serve static assets
	router.GET("/", serveIndex)
	router.ServeFiles("/stylesheets/*filepath", http.Dir("static/stylesheets"))

	return &Logger{router}
}

func main() {
	// Fire up the server
	fmt.Println("Server listening on port 3000")
	http.ListenAndServe("localhost:3000", App())
}

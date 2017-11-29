package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// Logger is a cool type
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()

	l.handler.ServeHTTP(w, req)

	log.Printf("- %s %s (%s)", req.Method, req.URL.Path, time.Since(start).String())
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
	db := NewDB()

	// Add a handler on /hello
	router.GET("/hello", HelloWorld)

	router.GET("/campaigns", CampaignsIndex(db))

	// Serve static assets
	router.GET("/", serveIndex)
	router.ServeFiles("/stylesheets/*filepath", http.Dir("static/stylesheets"))

	return &Logger{router}
}

func CampaignsIndex(db *sql.DB) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var title, author string
		err := db.QueryRow("SELECT title, author FROM campaigns").Scan(&title, &author)
		if err != nil && err == sql.ErrNoRows {
			fmt.Fprintf(rw, "No campaigns found")
			return
		}
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(rw, "The first campaign is '%s' by '%s'", title, author)
	}
}

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://gcappellotto@localhost/issuehunter?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Fire up the server
	fmt.Println("Server listening on port " + port)
	http.ListenAndServe(":"+port, App())
}

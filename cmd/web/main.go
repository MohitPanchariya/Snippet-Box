package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	static_dir := flag.String("static-dir", "./ui/static/", "Static assests path relative to project root")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Handler to serve static files
	fileServer := http.FileServer(http.Dir(*static_dir))

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s\n", srv.Addr)

	err := srv.ListenAndServe()

	if err != nil {
		errorLog.Fatal(err)
	}
}

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog      *log.Logger
	errorLog     *log.Logger
	staticAssets string // Path to static assests
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	static_dir := flag.String("static-dir", "./ui/static/", "Static assests path relative to project root")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:      infoLog,
		errorLog:     errorLog,
		staticAssets: *static_dir,
	}

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s\n", srv.Addr)

	err := srv.ListenAndServe()

	if err != nil {
		errorLog.Fatal(err)
	}
}

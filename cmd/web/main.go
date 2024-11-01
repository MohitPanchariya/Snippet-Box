package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MohitPanchariya/Snippet-Box/internal/models"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog        *log.Logger
	errorLog       *log.Logger
	staticAssets   string // Path to static assests
	snippetModel   *models.SnippetModel
	usersModel     *models.UserModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	static_dir := flag.String("static-dir", "./ui/static/", "Static assests path relative to project root")
	dsn := flag.String("dsn", "", "MySQL Data Source Name")
	certfile := flag.String("certfile", "", "TLS certifile path")
	keyfile := flag.String("keyfile", "", "TLS keyfile path")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if *dsn == "" {
		errorLog.Fatal("dsn required")
	}

	if *certfile == "" {
		errorLog.Fatal("certfile required")
	}

	if *keyfile == "" {
		errorLog.Fatal("keyfile required")
	}

	// Databse connection pool
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	app := &application{
		infoLog:        infoLog,
		errorLog:       errorLog,
		staticAssets:   *static_dir,
		snippetModel:   &models.SnippetModel{DB: db},
		usersModel:     &models.UserModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    form.NewDecoder(),
		sessionManager: sessionManager,
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     errorLog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s\n", srv.Addr)

	err = srv.ListenAndServeTLS(*certfile, *keyfile)
	if err != nil {
		errorLog.Fatal(err)
	}
}

// Return a database connection pool. Returns an error if connection
// pool can't be initialized or a ping check fails.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

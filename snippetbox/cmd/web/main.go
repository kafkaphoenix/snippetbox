package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	//addr := os.Getenv("SNIPPETBOX_ADDR")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}

	// web server
	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on %s", cfg.addr)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}

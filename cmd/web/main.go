package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct{
	errorlog *log.Logger
	infolog *log.Logger 
}

func main() {

	addr := flag.String("addr",":4000","HTTP Network Address")
	flag.Parse()


	errorlog := log.New(os.Stdout,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)
	infolog := log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	app := application{
		errorlog: errorlog,
		infolog: infolog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorlog,
		Handler: mux,
	}

	mux.Handle("/static/",http.StripPrefix("/static",fileServer))

	infolog.Printf("Starting server on %s\n",*addr)
	err := srv.ListenAndServe()
	
	errorlog.Fatal(err)
}
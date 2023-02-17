package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct{
	errorlog *log.Logger
	infolog *log.Logger 
}

func main() {

	addr := flag.String("addr",":4000","HTTP Network Address")
	dsn := flag.String("dsn","web:12345678@/snippetbox?parseTime=true","Mysql Data Source Name")
	flag.Parse()


	errorlog := log.New(os.Stdout,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)
	infolog := log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	app := application{
		errorlog: errorlog,
		infolog: infolog,
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorlog,
		Handler: app.routes(),
	}
	db, err := openDB(*dsn)
	if err != nil{
		errorlog.Fatal(err)
	}
	defer db.Close()

	
	infolog.Printf("Starting server on %s\n",*addr)
	err = srv.ListenAndServe()
	
	errorlog.Fatal(err)
}

func openDB(dsn string) (*sql.DB,error){
	db,err := sql.Open("mysql",dsn) 
	if err != nil{
		return nil,err
	}
	if err = db.Ping(); err != nil{
		return nil,err
	}
	return db,nil
}
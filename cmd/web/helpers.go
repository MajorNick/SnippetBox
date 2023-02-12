package main

import (
	"net/http"
	"runtime/debug"
	"fmt"
)

func (app * application) serverError(w http.ResponseWriter,err error){
	trace := fmt.Sprintf("%s\n%s",err,debug.Stack)
	app.errorlog.Println(trace)
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
	}
package main

import (
	"net/http"
	"runtime/debug"
	"fmt"
	
)

func (app *application) render(w http.ResponseWriter,r *http.Request, name string,data *templateData){
	tp , ok := app.templateCache[name]
	if !ok{
		app.serverError(w,fmt.Errorf("template %s doesn't exist",name))
		return 
	} 
	
	err  := tp.Execute(w,data)
	if err != nil{
		app.serverError(w,err)
	}
}

func (app * application) serverError(w http.ResponseWriter,err error){
	
	trace := fmt.Sprintf("%s\n%s",err,debug.Stack)
	app.errorlog.Println(trace,)
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
	}
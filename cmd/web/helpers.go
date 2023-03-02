package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)
func (app *application) addDefaultData(data *templateData,r *http.Request)*templateData{
	if data == nil{
		return &templateData{}
	}
	data.CurrentYear = time.Now().Year()
	return data
}

func (app *application) render(w http.ResponseWriter,r *http.Request, name string,data *templateData){
	tp , ok := app.templateCache[name]
	if !ok{
		app.serverError(w,fmt.Errorf("template %s doesn't exist",name))
		return 
	} 
	buf := new(bytes.Buffer)

	err  := tp.Execute(buf,app.addDefaultData(data,r))
	if err != nil{
		app.serverError(w,err)
	}
	buf.WriteTo(w)
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
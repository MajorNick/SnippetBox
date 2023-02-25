package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/MajorNick/snippetbox/pkg/models"
)
func (app *application)home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		}
	s, err := app.snippets.Latest()

	if err!= nil{
		app.serverError(w,err)
	}

	data := &templateData{Snippets: s}
	

	ts,err := template.ParseFiles(files...)
	if err!= nil{
		app.serverError(w,err)
		http.Error(w, "Internal Server Error", 500)
	}
	err = ts.Execute(w,data)
	if err!= nil{
		app.serverError(w,err)
		http.Error(w, "Internal Server Error", 500)
	}
	
}
func (app *application)showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
	}else{
		if err != nil{
			app.serverError(w,err)
			return
		}
	}
	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	
	data := &templateData{Snippet: s}
	ts, err := template.ParseFiles(files...)
	if err != nil{
		app.serverError(w,err)
		return
	}
	err = ts.Execute(w,data)
	if err != nil{
		app.serverError(w,err)
		return
	}

	
}
func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w,http.StatusMethodNotAllowed)
		return
	}
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"
	id, err  := app.snippets.Insert(title,content,expires)
	if err != nil{
		app.serverError(w,err)
		return
	}
	
	http.Redirect(w,r,fmt.Sprintf("/snippet?id=%d",id),http.StatusSeeOther)
	
}
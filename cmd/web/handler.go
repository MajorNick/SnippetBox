package main
import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	
)
func (app *application)home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		}


	ts,err := template.ParseFiles(files...)
	if err!= nil{
		app.errorlog.Println(ts)
		http.Error(w, "Internal Server Error", 500)
	}
	err = ts.Execute(w,nil)
	if err!= nil{
		app.errorlog.Println(ts)
		http.Error(w, "Internal Server Error", 500)
	}
	
}
func (app *application)showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
package main

import (
	"html/template"
	"log"
	"path/filepath"

	"github.com/MajorNick/snippetbox/pkg/models"
)

type templateData struct{
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template,error){
	
	cache := make(map[string]*template.Template)
	
	pages,err := filepath.Glob(filepath.Join(dir,"*.page.tmpl"))
	if err != nil{
		return nil,err
	}
	//  _,err = template.ParseFiles("ui/html/home.page.tmpl")
	
	// if err != nil{
	// 	log.Println("fsasfsa")
	// 	return nil,err
	// }
	
	for _,page := range pages{
		name := filepath.Base(page)
		ts,err := template.ParseFiles(name)
		
		if err != nil{
			log.Println(name)
			return nil,err
		}

		ts, err  = ts.ParseGlob(filepath.Join(dir,"*.layout.tmpl"))
		if err != nil{
			return nil, err
		}
		ts, err  = ts.ParseGlob(filepath.Join(dir,"*.partial.tmpl"))
		if err != nil{
			return nil, err
		}
		cache[name] = ts

	}
	return cache,nil
}
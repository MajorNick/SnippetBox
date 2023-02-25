package main

import (
	"html/template"

	"github.com/MajorNick/snippetbox/pkg/models"
)

type templateData struct{
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template){
	cache := map[string]*template.Template{}

}
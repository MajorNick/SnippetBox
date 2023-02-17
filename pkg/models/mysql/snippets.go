package mysql

import (
	"database/sql"
	"github.com/MajorNick/snippetbox/pkg/models"
)

type SnippetModel struct{
	Db *sql.DB
}

// insert snippet

func (m * SnippetModel) Insert(title,content,expires string) (int,error){
	return 0, nil
}
// get snippet by id
func (m * SnippetModel) Get(id int) (*models.Snippets,error){
	return nil,nil
}
//return 10 latest added snippets
func (m * SnippetModel) Latest()(*[]models.Snippets,error){
	return nil,nil
}
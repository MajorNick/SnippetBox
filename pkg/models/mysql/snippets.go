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

	stmt := `INSERT INTO snippets (title,content,created,expires) VALUES(
		?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(),INTERVAL ? DAY));
	`

	result, err := m.Db.Exec(stmt,title,content,expires)
	if err != nil{
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil{
		return 0, err 
	}
	return int(id), nil
}
// get snippet by id
func (m * SnippetModel) Get(id int) (*models.Snippets,error){
	return nil,nil
}
//return 10 latest added snippets
func (m * SnippetModel) Latest()(*[]models.Snippets,error){
	return nil,nil
}
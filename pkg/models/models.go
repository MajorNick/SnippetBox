package models

import (
	"errors"
	"time"
)

var errNoRecord error = errors.New("models: No Matching record found")

type Snippets struct{
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}
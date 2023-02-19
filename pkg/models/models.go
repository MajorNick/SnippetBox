package models

import (
	"errors"
	"time"
)

var ErrNoRecord error = errors.New("models: No Matching record found")

type Snippet struct{
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}
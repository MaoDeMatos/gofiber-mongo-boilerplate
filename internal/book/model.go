package book

import "github.com/kamva/mgm/v3"

type Book struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Author           string `json:"author" bson:"author"`
	ISBN             string `json:"isbn" bson:"isbn"`
}

// func NewBook(title, author, isbn string) *Book {
// 	return &Book{
// 		Title:  title,
// 		Author: author,
// 		ISBN:   isbn,
// 	}
// }

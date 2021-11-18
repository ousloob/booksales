// Package bookstore implements functions ton manage a book store.
package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book.
type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

// Buy diminush the number of copies of a book if it's available.
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies--
	return b, nil
}

// GetAllBooks shows us the list of all the books of a specific catalog.
func GetAllBooks(catalog map[int]Book) []Book {
	var books []Book
	for _, book := range catalog {
		books = append(books, book)
	}
	return books
}

// GetBook gets the book based on a given id and title.
func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}

	return b, nil
}

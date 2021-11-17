// Package bookstore implements functions ton manage a book store.
package bookstore

import "errors"

// Book represents information about a book.
type Book struct {
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

// GetAllBooks show us the list of all the books of a specific catalog.
func GetAllBooks(catalog []Book) []Book {
	return catalog
}

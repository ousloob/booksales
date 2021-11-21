// Package bookstore implements functions ton manage a book store.
package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book.
type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

// Catalog is a slice of Book.
type Catalog map[int]Book

// GetAllBooks shows us the list of all the books of a specific catalog.
func (c Catalog) GetAllBooks() []Book {
	var books []Book
	for _, book := range c {
		books = append(books, book)
	}

	return books
}

// GetBook gets the book based on a given id and title.
func (c Catalog) GetBook(ID int) (Book, error) {
	b, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}

	return b, nil
}

// Buy diminush the number of copies of a book if it's available.
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies--
	return b, nil
}

// NetPriceCents calculs the price with the discount in cents.
func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

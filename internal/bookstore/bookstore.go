// Package bookstore implements functions ton manage a book store.
package bookstore

import (
	"errors"
	"fmt"
)

const (
	CategoryFantasy Category = iota
	CategoryManga
	CategoryDevelopment
)

type Category int

var validCategory = map[Category]bool{
	CategoryFantasy:     true,
	CategoryManga:       true,
	CategoryDevelopment: true,
}

// Book represents information about a book.
type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
	category        Category
}

// SetPriceCents updates the price of a book.
func (b *Book) SetPriceCents(newPrice int) error {
	if newPrice < 0 {
		return fmt.Errorf("invalid price %d want a price superior of zero", newPrice)
	}

	b.PriceCents = newPrice
	return nil
}

// SetCategory helps us to set a category.
func (b *Book) SetCategory(ctg Category) error {
	if !validCategory[ctg] {
		return fmt.Errorf("unknown category %v", ctg)
	}

	b.category = ctg
	return nil
}

// Category return the category of a book.
func (b Book) Category() Category {
	return b.category
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

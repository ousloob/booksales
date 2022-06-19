// Package bookery provides a functionality of a core business API for book.
package bookery

// Book represents information about a book.
type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	Price           float64
	DiscountPercent int
	Category        string
}

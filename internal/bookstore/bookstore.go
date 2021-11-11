// Package bookstore implements functions ton manage a book store.
package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

// Buy diminush the number of copies of a book if it's available.
func Buy(b Book) Book {
	if b.Copies == 0 {
		return b
	}

	b.Copies--
	return b
}

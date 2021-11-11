package bookstore_test

import (
	"booklib/internal/bookstore"
	"testing"
)

// type testCases struct {
// 	book bookstore.Book
// 	got  int
// 	want int
// }

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "The Witcher: The Last Wish",
		Author: "Andrzej Sapkowski",
		Copies: 3,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	// tcs := []testCases{
	// 	{bookstore.Book{Title: "The Lord of the Rings", Author: "J. R. R. Tolkien", Copies: 3}, 3, 2},
	// 	{bookstore.Book{Title: "Dragon Age: Asunder", Author: "David Gaider", Copies: 0}, 0, 0},
	// }

	b := bookstore.Book{
		Title:  "The Lord of the Rings",
		Author: "J. R. R. Tolkien",
		Copies: 2,
	}

	want := 1
	result := bookstore.Buy(b)
	got := result.Copies

	if want != got {
		t.Errorf("want %d copies after buying one copy from a stock of 2, got %d", want, got)
	}

}

package bookstore_test

import (
	"booklib/internal/bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// type testCases struct {
// 	book bookstore.Book
// 	got  int
// 	want int
// }

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		ID:     1,
		Title:  "The Witcher: The Last Wish",
		Author: "Andrzej Sapkowski",
		Copies: 3,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	// TODO: manage a slice of test cases.
	// tcs := []testCases{
	// 	{bookstore.Book{Title: "The Lord of the Rings", Author: "J. R. R. Tolkien", Copies: 3}, 3, 2},
	// 	{bookstore.Book{Title: "Dragon Age: Asunder", Author: "David Gaider", Copies: 0}, 0, 0},
	// }

	b := bookstore.Book{
		ID:     9,
		Title:  "The Lord of the Rings",
		Author: "J. R. R. Tolkien",
		Copies: 2,
	}

	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies

	if want != got {
		t.Errorf("want %d copies after buying one copy from a stock of 2, got %d", want, got)
	}

}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		ID:     13,
		Title:  "Dragon Age: Asunder",
		Author: "David Gaider",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)
	if err == nil {
		t.Errorf("want error buying book when zero copies left, but got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	catalog := []bookstore.Book{
		{Title: "One Piece"},
		{Title: "The Promises Neverland"},
		{Title: "Kimetsu No Yaiba"},
	}

	want := []bookstore.Book{
		{Title: "One Piece"},
		{Title: "The Promises Neverland"},
		{Title: "Kimetsu No Yaiba"},
	}

	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error((cmp.Diff(want, got)))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	catalog := []bookstore.Book{
		{ID: 2, Title: "The Witcher: Sword of Destiny"},
		{ID: 3, Title: "The Witcher: Season of Storms"},
	}

	want := bookstore.Book{
		ID:    3,
		Title: "The Witcher: Season of Storms",
	}

	got := bookstore.GetBook(catalog, 3)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

package bookstore_test

import (
	"booklib/internal/bookstore"
	"sort"
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

	catalog := map[int]bookstore.Book{
		33: {ID: 33, Title: "One Piece"},
		34: {ID: 34, Title: "The Promises Neverland"},
		35: {ID: 35, Title: "Kimetsu No Yaiba"},
	}

	want := []bookstore.Book{
		{ID: 33, Title: "One Piece"},
		{ID: 34, Title: "The Promises Neverland"},
		{ID: 35, Title: "Kimetsu No Yaiba"},
	}

	got := bookstore.GetAllBooks(catalog)
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got) {
		t.Error((cmp.Diff(want, got)))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	catalog := map[int]bookstore.Book{
		2: {ID: 2, Title: "The Witcher: Sword of Destiny"},
		3: {ID: 3, Title: "The Witcher: Season of Storms"},
	}

	want := bookstore.Book{
		ID:    3,
		Title: "The Witcher: Season of Storms",
	}

	got, err := bookstore.GetBook(catalog, 3)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	catalog := map[int]bookstore.Book{}

	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Fatal("want error for non-esistent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	book := bookstore.Book{
		Title:           "The Witcher: Blood of Elves",
		PriceCents:      5000,
		DiscountPercent: 25,
	}

	want := 3750
	got := book.NetPriceCents()

	if want != got {
		t.Errorf("with price %d, after %d%% discount want net %d, got %d", book.PriceCents, book.DiscountPercent, want, got)
	}
}

package bookstore_test

import (
	"booklib/internal/bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

	c := bookstore.Catalog{
		33: {ID: 33, Title: "Kimetsu No Yaiba"},
		34: {ID: 34, Title: "One Piece"},
		35: {ID: 35, Title: "The Promises Neverland"},
	}

	want := []bookstore.Book{
		{ID: 33, Title: "Kimetsu No Yaiba"},
		{ID: 34, Title: "One Piece"},
		{ID: 35, Title: "The Promises Neverland"},
	}

	got := c.GetAllBooks()
	sort.Slice(got, func(i, j int) bool { return got[i].ID < got[j].ID })

	if !cmp.Equal(want, got, cmp.AllowUnexported(bookstore.Book{})) {
		t.Error((cmp.Diff(want, got, cmp.AllowUnexported(bookstore.Book{}))))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	c := bookstore.Catalog{
		2: {ID: 2, Title: "The Witcher: Sword of Destiny"},
		3: {ID: 3, Title: "The Witcher: Season of Storms"},
	}

	want := bookstore.Book{
		ID:    3,
		Title: "The Witcher: Season of Storms",
	}

	got, err := c.GetBook(3)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got, cmp.AllowUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got, cmp.AllowUnexported(bookstore.Book{})))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	c := bookstore.Catalog{}

	_, err := c.GetBook(999)
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

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "The Witcher: Time of Contempt",
		PriceCents: 325,
	}
	want := 5000
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}

	got := b.PriceCents

	if want != got {
		t.Errorf("want to update the price, want %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "The Witcher: Baptism of Fire",
		PriceCents: 425,
	}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Errorf("want error setting invalid price -1, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "The Witcher: The Tower of the Swallow",
	}

	for cat := range bookstore.ValidCategory {
		err := b.SetCategory(cat)
		if err != nil {
			t.Fatal(err)
		}
		got := b.Category
		if cat != got {
			t.Errorf("want to update category, want %q, got %q", cat, got)
		}
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "The Witcher: The Tower of the Swallow",
	}
	err := b.SetCategory("random")
	if err == nil {
		t.Fatal("want error setting invalid category 'game', got nil")
	}
}

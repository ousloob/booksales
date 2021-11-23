package creditcard

import (
	"errors"
)

type card struct {
	number string
}

func New(n string) (card, error) {
	if n == "" {
		return card{}, errors.New("number of card empty")
	}

	return card{number: n}, nil
}

func (c *card) Number() string {
	return c.number
}

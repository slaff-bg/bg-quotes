package domain

import (
	"github.com/google/uuid"
)

type Author struct {
	AuthorID                             uuid.UUID
	FirstName, SecondName, AKA, ImageURL string
}

func CreateAuthor(fn, sn, aka, img string) Author {
	return Author{
		AuthorID:   uuid.New(),
		FirstName:  fn,
		SecondName: sn,
		AKA:        aka,
		ImageURL:   img,
	}
}

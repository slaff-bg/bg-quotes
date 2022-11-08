package domain

import (
	"github.com/google/uuid"
)

type Author struct {
	AuthorId                             uuid.UUID
	FirstName, SecondName, AKA, ImageURL string
}

func CreateAuthor(fn, sn, aka, img string) Author {
	return Author{
		AuthorId:   uuid.New(),
		FirstName:  fn,
		SecondName: sn,
		AKA:        aka,
		ImageURL:   img,
	}
}

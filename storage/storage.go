package storage

import (
	"bg-quotes/domain"
	"sync"

	"github.com/google/uuid"
)

var aData = sync.Map{}

func AuthorAdd(a domain.Author) domain.Author {
	aData.Store(a.AuthorId, a)
	return a
}

func AuthorGet(id uuid.UUID) (domain.Author, bool) {
	author, found := aData.Load(id)
	return author.(domain.Author), found
}

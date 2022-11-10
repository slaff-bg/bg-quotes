package storage

import (
	"bg-quotes/domain"
	"sync"

	"github.com/google/uuid"
)

var aData = sync.Map{}
var qData = sync.Map{}

// Author functions

func AuthorAdd(a domain.Author) domain.Author {
	aData.Store(a.AuthorID, a)
	return a
}

func AuthorGet(id uuid.UUID) (domain.Author, bool) {
	author, found := aData.Load(id)
	if found {
		return author.(domain.Author), found
	}

	return domain.Author{}, found
}

// Quote functions

func QuoteAdd(q domain.Quote) domain.Quote {
	qData.Store(q.QuoteID, q)
	return q
}

func QuoteGet(id uuid.UUID) (domain.Quote, bool) {
	quote, found := qData.Load(id)
	if found {
		return quote.(domain.Quote), found
	}

	return domain.Quote{}, found
}

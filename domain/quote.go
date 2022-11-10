package domain

import "github.com/google/uuid"

type Quote struct {
	QuoteID     uuid.UUID
	Quote       string
	SmokingRoom bool   // Indicates whether the quote has content requiring a minimum age (18+).
	AuthorID    string // To be changed to uuid.UUID later.
}

func CreateQuote(quote string, smokingRoom bool, author string) Quote {
	return Quote{
		QuoteID:     uuid.New(),
		Quote:       quote,
		SmokingRoom: smokingRoom,
		AuthorID:    author,
	}
}

// https://theculturetrip.com/europe/bulgaria/articles/11-sayings-and-proverbs-to-help-you-understand-bulgarian-culture/

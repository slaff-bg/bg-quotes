package api

import (
	"bg-quotes/domain"

	"github.com/google/uuid"
)

type QuoteDTO struct {
	QuoteID     uuid.UUID `json:"quote_id"`
	Quote       string    `json:"quote"`
	SmokingRoom bool      `json:"smoking_room"`
	Author      any       `json:"author"`
	// Author      AuthorDTO
}

// I set the third argument to be of type any. This is because in case the author
// is unknown in the JSON response, I will return an empty JSON object instead of
// the JSON of the AuthorDTO structure. This way, I'll keep the structure of the
// JSON response without having to add a new argument, whether the author is known or not.
// So in both cases, we'll have a JSON object, no matter if it's empty or not.
func createQuoteDTO(q domain.Quote, a any) QuoteDTO {
	return QuoteDTO{
		QuoteID:     q.QuoteID,
		Quote:       q.Quote,
		SmokingRoom: q.SmokingRoom,
		Author:      a,
	}
}

package domain

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuote(t *testing.T) {
	// Validation of an empty or shorter-than-necessary length Quote is done in the foreground at the HTTP/handler level.
	// This is the reason why I have left out some scenarios here.
	// That is, we assume that we have validated input data. This is just a unit test.
	tests := []struct {
		Quote       string
		SmokingRoom bool
		AuthorID    string
	}{
		{"If I feel like working, I sit down and wait for it to pass.", false, ""},
		{"If I feel like working, I sit down and wait for it to pass.", false, "25937d34-16ad-4083-98cb-ddd7fe3d5fab"},
		{"If I feel like working, I sit down and wait for it to pass.", true, ""},
		{"If I feel like working, I sit down and wait for it to pass.", true, "25937d34-16ad-4083-98cb-ddd7fe3d5fab"},
		{"Ако ми се работи, сядам и чекам да ми мине.", false, ""},
		{"Ако ми се работи, сядам и чекам да ми мине.", false, "25937d34-16ad-4083-98cb-ddd7fe3d5fab"},
		{"Ако ми се работи, сядам и чекам да ми мине.", true, ""},
		{"Ако ми се работи, сядам и чекам да ми мине.", true, "25937d34-16ad-4083-98cb-ddd7fe3d5fab"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {
			quote := CreateQuote(tc.Quote, tc.SmokingRoom, tc.AuthorID)

			// Tests the type of the response.
			var rt interface{} = quote
			if _, rtok := rt.(Quote); !rtok {
				t.Fatalf("CreateQuote: got %T want domain.Quote", rt)
			}

			// Tests generated Quote UUID.
			var x interface{} = quote.QuoteID
			_, xok := x.(uuid.UUID)
			assert.True(t, xok)

			// Tests that the transmitted data has the correct setup.
			assert.Equal(t, tc.Quote, quote.Quote)
			assert.Equal(t, tc.SmokingRoom, quote.SmokingRoom)
			assert.Equal(t, tc.AuthorID, quote.AuthorID)
		})
	}
}

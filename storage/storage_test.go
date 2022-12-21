package storage

import (
	"bg-quotes/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorCreate(t *testing.T) {
	tests := []struct {
		FirstName, SecondName, AKA, ImageURL string
	}{
		{"Йордан", "Радичков", "Йордан Радичков", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg"},
		{"Yordan", "Radichkov", "Yordan Radichkov", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {
			da := domain.CreateAuthor(tc.FirstName, tc.SecondName, tc.AKA, tc.ImageURL)
			sa := AuthorCreate(da)

			assert.Equal(t, sa.FirstName, da.FirstName)
			assert.Equal(t, sa.AKA, da.AKA)
		})
	}
}

func TestAuthorRead(t *testing.T) {
	tests := make(map[string]domain.Author, 2)
	tests["tc1"] = domain.CreateAuthor("Йордан", "Радичков", "Йордан Радичков", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg")
	tests["tc2"] = domain.CreateAuthor("Yordan", "Radichkov", "Yordan Radichkov", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg")

	for k, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", k), func(t *testing.T) {
			sa := AuthorCreate(tc)

			got, ok := AuthorRead(sa.AuthorID)
			assert.True(t, ok)
			assert.Equal(t, sa.FirstName, got.FirstName)
			assert.Equal(t, sa.AKA, got.AKA)
		})
	}
}

func TestQuoteCreate(t *testing.T) {
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
			dq := domain.CreateQuote(tc.Quote, tc.SmokingRoom, tc.AuthorID)
			sq := QuoteCreate(dq)

			assert.Equal(t, sq.Quote, dq.Quote)
			assert.Equal(t, sq.SmokingRoom, dq.SmokingRoom)
			assert.Equal(t, sq.AuthorID, dq.AuthorID)
		})
	}
}

func TestQuoteRead(t *testing.T) {
	tests := make(map[string]domain.Quote, 2)
	tests["tc1"] = domain.CreateQuote("If I feel like working, I sit down and wait for it to pass.", false, "")
	tests["tc2"] = domain.CreateQuote("Ако ми се работи, сядам и чекам да ми мине.", true, "25937d34-16ad-4083-98cb-ddd7fe3d5fab")

	for k, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", k), func(t *testing.T) {
			sq := QuoteCreate(tc)

			got, ok := QuoteRead(sq.QuoteID)
			assert.True(t, ok)
			assert.Equal(t, sq.Quote, got.Quote)
			assert.Equal(t, sq.SmokingRoom, got.SmokingRoom)
			assert.Equal(t, sq.AuthorID, got.AuthorID)
		})
	}
}

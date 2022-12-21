package domain

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor(t *testing.T) {
	tests := []struct {
		FirstName  string
		SecondName string
		AKA        string
		ImageURL   string
	}{
		{"Yordan", "Radichkov", "Yordan Radichkov", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg"},
		{"Йордан", "Радичков", "Йордан Радичков", "https://upload.wikimedia.org/wikipedia/en/thumb/1/1a/Yordan_Radichkov.jpg/200px-Yordan_Radichkov.jpg"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {
			author := CreateAuthor(tc.FirstName, tc.SecondName, tc.AKA, tc.ImageURL)

			// Tests the type of the response.
			var rt interface{} = author
			if _, rtok := rt.(Author); !rtok {
				t.Fatalf("CreateAuthor: got %T want domain.Author", rt)
			}

			// Tests generated Author UUID.
			var x interface{} = author.AuthorID
			_, xok := x.(uuid.UUID)
			assert.True(t, xok)

			// Tests that the transmitted data has the correct setup.
			assert.Equal(t, tc.FirstName, author.FirstName)
			assert.Equal(t, tc.SecondName, author.SecondName)
			assert.Equal(t, tc.AKA, author.AKA)
			assert.Equal(t, tc.ImageURL, author.ImageURL)
		})
	}
}

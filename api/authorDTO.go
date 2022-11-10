package api

import (
	"bg-quotes/domain"

	"github.com/google/uuid"
)

type AuthorDTO struct {
	AuthorID   uuid.UUID `json:"author_id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	AKA        string    `json:"aka"`
	ImageURL   string    `json:"img_url"`
}

// Wrap Author (struct) to Data Transfer Object
func createAuthorDTO(a domain.Author) AuthorDTO {
	return AuthorDTO{
		AuthorID:   a.AuthorID,
		FirstName:  a.FirstName,
		SecondName: a.SecondName,
		AKA:        a.AKA,
		ImageURL:   a.ImageURL,
	}
}

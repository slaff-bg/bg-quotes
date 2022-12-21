package api

import (
	"bg-quotes/storage"

	"github.com/google/uuid"
)

// Checks whether:
//   - the passed UUID string can be parsed
//   - there is an author with such a UUID
//
// ... and returns an empty interface or AuthorDTO.
func GetAuthorDTO(uid string) (interface{}, error) {
	var author interface{}

	if len(uid) > 0 {
		aid, err := uuid.Parse(uid)
		if err != nil {
			return author, err
		}

		if a, found := storage.AuthorRead(aid); found {
			author = createAuthorDTO(a)
		}
	}

	// I keep the Quotes JSON structure by adding an empty JSON object to the Author if its current state is nil.
	// Thus, the frontend part can always rely on Author as a JSON object instead of doing a JSON/nil check.
	// In other words, if the response status is 200, an idempotent payload can be relied upon.
	if author == nil {
		author = make(map[string]string)
	}

	return author, nil
}

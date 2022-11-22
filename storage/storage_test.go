package storage

import (
	"bg-quotes/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorAdd(t *testing.T) {
	da := domain.CreateAuthor("Henry", "Bukowski", "Charles Bukowski", "https://upload.wikimedia.org/wikipedia/en/e/e2/Charles_Bukowski_smoking.jpg")
	sa := AuthorAdd(da)

	assert.Equal(t, sa.FirstName, "Henry")
	assert.Equal(t, sa.AKA, "Charles Bukowski")
}

func TestAuthorGet(t *testing.T) {
	da := domain.CreateAuthor("Henry", "Bukowski", "Charles Bukowski", "https://upload.wikimedia.org/wikipedia/en/e/e2/Charles_Bukowski_smoking.jpg")
	sa := AuthorAdd(da)

	assert.Equal(t, sa.FirstName, "Henry")
	assert.Equal(t, sa.AKA, "Charles Bukowski")
}

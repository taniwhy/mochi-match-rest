package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
)

func TestContainsRecord(t *testing.T) {
	record := []input.FavoriteGameRecord{{GameTitle: "foo"}}
	assert.Equal(t, containsRecord(record, "foo"), true)

	record = []input.FavoriteGameRecord{{GameTitle: "foo"}}
	assert.Equal(t, containsRecord(record, "bar"), false)
}

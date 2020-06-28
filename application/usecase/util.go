package usecase

import "github.com/taniwhy/mochi-match-rest/domain/models/input"

func containsRecord(record []input.FavoriteGameRecord, id string) bool {
	for _, r := range record {
		if id == r.GameTitle {
			return true
		}
	}
	return false
}

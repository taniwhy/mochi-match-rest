package usecase

import "github.com/taniwhy/mochi-match-rest/domain/models/input"

func contains(gr []input.FavoriteGameRecord, id string) bool {
	for _, r := range gr {
		if id == r.GameTitle {
			return true
		}
	}
	return false
}

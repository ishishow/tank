package cache

import (
	rm "20dojo-online/pkg/domain/model/ranking"
	um "20dojo-online/pkg/domain/model/user"
)

// RankingRepository is
type RankingRepository interface {
	GetRankingList(start int64) ([]*rm.RankInfo, error)
	SaveRankingList(users []*um.User) error
	SaveRanking(user *um.User) error
}

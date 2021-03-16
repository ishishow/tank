package interactor

import (
	"log"

	"20dojo-online/pkg/domain/repository/cache"

	rm "20dojo-online/pkg/domain/model/ranking"
	ur "20dojo-online/pkg/domain/repository/db"
)

type rankingUseCase struct {
	userRepository    ur.UserRepository
	rankingRepository cache.RankingRepository
}

// RankingUseCase is
type RankingUseCase interface {
	GetRankingList(start int64) ([]*rm.RankInfo, error)
	SaveRankingList() error
}

// NewRankingUseCase Userデータに関するユースケースを生成
func NewRankingUseCase(userRepo ur.UserRepository,
	rankRepo cache.RankingRepository) RankingUseCase {
	rankingUseCase := &rankingUseCase{
		userRepository:    userRepo,
		rankingRepository: rankRepo,
	}
	rankingUseCase.SaveRankingList()
	return rankingUseCase
}

func (ru rankingUseCase) GetRankingList(start int64) ([]*rm.RankInfo, error) {
	return ru.rankingRepository.GetRankingList(start)
}

// SetRankingData ハイスコアランキングデータをソート済みマップにセット
func (ru rankingUseCase) SaveRankingList() error {
	// ランキングデータ取ってくる
	users, err := ru.userRepository.SelectAllPlayingUsers()
	if err != nil {
		log.Printf("getting users from DB is failed: %v", err)
	}

	return ru.rankingRepository.SaveRankingList(users)
}

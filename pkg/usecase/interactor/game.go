package interactor

import (
	"fmt"

	"20dojo-online/pkg/domain/repository/cache"
	ur "20dojo-online/pkg/domain/repository/db"
)

type gameUseCase struct {
	userRepository ur.UserRepository
	rankingRepository cache.RankingRepository
}

// GameUseCase is
type GameUseCase interface {
	FinishGame(userID string, score int32) (coin int32, err error)
}

// NewGameUseCase Userデータに関するユースケースを生成
func NewGameUseCase(userRepo ur.UserRepository,
	rankRepo cache.RankingRepository) GameUseCase {
	return &gameUseCase{
		userRepository: userRepo,
		rankingRepository: rankRepo,
	}
}

func (uu gameUseCase) FinishGame(userID string, score int32) (int32, error) {
	user, err := uu.userRepository.SelectByPrimaryKey(userID)
	if err != nil {
		return 0, err
	}
	// 得点が最高値かどうか判別
	if user.HighScore < score {
		fmt.Println(user.Coin, score, user.HighScore)
		user.HighScore = score
	}
	// 実行userにコインを付与
	user.Coin += score
	_ = uu.userRepository.Update(user)

	//rediscacheを更新
	if err := uu.rankingRepository.SaveRanking(user); err != nil {
		return 0, err
	}
	return user.Coin, nil
}

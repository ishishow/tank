package handler

import "20dojo-online/pkg/usecase/interactor"

// Handler 全体ハンドラ
type Handler struct {
	UserUseCase           interactor.UserUseCase
	GachaUseCase          interactor.GachaUseCase
	GameUseCase           interactor.GameUseCase
	CollectionItemUseCase interactor.CollectionItemUseCase
	RankingUseCase        interactor.RankingUseCase
}

// NewHandler Userデータに関するHandlerを生成
func NewHandler(userUseCase interactor.UserUseCase,
	gachaUseCase interactor.GachaUseCase,
	gameUseCase interactor.GameUseCase,
	collectionItemUseCase interactor.CollectionItemUseCase,
	rankingUseCase interactor.RankingUseCase) Handler {
	return Handler{
		UserUseCase:           userUseCase,
		GachaUseCase:          gachaUseCase,
		GameUseCase:           gameUseCase,
		CollectionItemUseCase: collectionItemUseCase,
		RankingUseCase:        rankingUseCase,
	}
}

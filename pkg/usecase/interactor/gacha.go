package interactor

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"20dojo-online/pkg/constant"
	gm "20dojo-online/pkg/domain/model/gacha_probability"
	ucm "20dojo-online/pkg/domain/model/user_collection_item"
	ur "20dojo-online/pkg/domain/repository/db"
)

// GachaUseCase is
type GachaUseCase interface {
	GachaDraw(ctx context.Context, userID string, times int32) ([]*GachaResult, error)
}

type gachaUseCase struct {
	userRepository               ur.UserRepository
	collectionItemRepository     ur.CollectionItemRepository
	userCollectionItemRepository ur.UserCollectionItemRepository
	gachaProbabilityRepository   ur.GachaProbabilityRepository
}

// NewGachaUseCase Userデータに関するユースケースを生成
func NewGachaUseCase(userRepo ur.UserRepository,
	collectionRepo ur.CollectionItemRepository,
	userCollectionRepo ur.UserCollectionItemRepository,
	gachaProbabilityRepo ur.GachaProbabilityRepository,
) GachaUseCase {
	return &gachaUseCase{
		userRepository:               userRepo,
		collectionItemRepository:     collectionRepo,
		userCollectionItemRepository: userCollectionRepo,
		gachaProbabilityRepository:   gachaProbabilityRepo,
	}
}

func (uu gachaUseCase) GachaDraw(ctx context.Context, userID string, times int32) ([]*GachaResult, error) {
	user, err := uu.userRepository.SelectByPrimaryKey(userID)
	if err != nil {
		return nil, err
	}

	// ユーザーが十分なコインを持っていなかった時のバリデーション
	if user.Coin < (constant.GachaCoinConsumption * times) {
		return nil, errors.New("this user don't have enough coins")
	}
	// ガチャ実行してコレクションアイテムを取得
	gachaProbabilities, err := uu.gachaProbabilityRepository.SelectAllGachaProbabilities()
	collectionItemIDs, err := getCollectionItems(times, gachaProbabilities)
	if err != nil {
		return nil, err
	}
	collectionItems, err := uu.collectionItemRepository.SelectCollectionItemsByPrimaryKeys(collectionItemIDs)
	// 現在ユーザーが所持しているコレクションアイテムを取得
	userCollectionItems, err := uu.userCollectionItemRepository.SelectUserCollectionItemsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// CollectionItemIDをkeyとしたMAPの作成
	userCollectionItemMapByCollectionItemID := make(map[string]struct{}, len(userCollectionItems))
	for _, userCollectionItem := range userCollectionItems {
		userCollectionItemMapByCollectionItemID[userCollectionItem.CollectionItemID] = struct{}{}
	}

	// collectionItemがすでに持っていないか判別、gachaDrawResponseの作成
	gachaResults := make([]*GachaResult, 0, times)
	newUserCollectionItems := []*ucm.UserCollectionItem{}
	for _, collectionItem := range collectionItems {
		_, hasItem := userCollectionItemMapByCollectionItemID[collectionItem.ID]
		gachaResult := &GachaResult{
			CollectionID: collectionItem.ID,
			Name:         collectionItem.Name,
			Rarity:       collectionItem.Rarity,
			IsNew:        !hasItem,
		}

		if gachaResult.IsNew {
			newUserCollectionItems = append(newUserCollectionItems, &ucm.UserCollectionItem{
				UserID:           userID,
				CollectionItemID: collectionItem.ID,
			})
		}
		gachaResults = append(gachaResults, gachaResult)
	}

	// 新規CollectionItemを全てDBに保存する
	user.Coin -= constant.GachaCoinConsumption * times
	if err := uu.gachaProbabilityRepository.GachaSave(ctx, newUserCollectionItems, user); err != nil {
		log.Println("gachaSave is failed")
		return nil, err
	}

	return gachaResults, nil
}

// getCollectionItems times回数分ガチャプログラムを実行してコレクションアイテムを取得する
func getCollectionItems(times int32, gachaProbabilities []*gm.GachaProbability) ([]string, error) {
	var sumRatio int32
	for _, gachaProbability := range gachaProbabilities {
		sumRatio += gachaProbability.Ratio
	}

	rand.Seed(time.Now().UnixNano())
	collectionItemIDs := make([]string, 0, times)
	for i := 0; i < int(times); i++ {
		randomRatio := rand.Int31n(sumRatio)
		var ratio int32
		for _, gachaProbability := range gachaProbabilities {
			ratio += gachaProbability.Ratio
			if randomRatio <= ratio {
				collectionItemIDs = append(collectionItemIDs, gachaProbability.CollectionItemID)
				break
			}
		}
	}
	return collectionItemIDs, nil
}

// GachaResult is
type GachaResult struct {
	CollectionID string `json:"collectionID"`
	Name         string `json:"name"`
	Rarity       int32  `json:"rarity"`
	IsNew        bool   `json:"isNew"`
}

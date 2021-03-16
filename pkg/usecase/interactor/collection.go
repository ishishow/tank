package interactor

import (
	ur "20dojo-online/pkg/domain/repository/db"
)

// CollectionItemUseCase is
type CollectionItemUseCase interface {
	SelectALLUserColectionItems(userID string) (*collectionListResponse, error)
}

type collectionItemUseCase struct {
	collectionRepository    ur.CollectionItemRepository
	userColectionRepository ur.UserCollectionItemRepository
}

// NewCollectionItemUseCase Userデータに関するユースケースを生成
func NewCollectionItemUseCase(collectionRepo ur.CollectionItemRepository,
	userCollectionRepo ur.UserCollectionItemRepository) CollectionItemUseCase {
	return &collectionItemUseCase{
		collectionRepository:    collectionRepo,
		userColectionRepository: userCollectionRepo,
	}
}

func (uu collectionItemUseCase) SelectALLUserColectionItems(userID string) (*collectionListResponse, error) {

	collectionItems, err := uu.collectionRepository.SelectAllCollectionItems()
	if err != nil {
		return nil, err
	}

	// userIDに紐づくUserCollectionItemを全件取得
	userCollectionItems, err := uu.userColectionRepository.SelectUserCollectionItemsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// CollectionItem_IDでkey判別するためのMAPを作成
	userCollectionItemMapByCollectionItemID := make(map[string]struct{}, len(collectionItems))
	for _, userCollectionItem := range userCollectionItems {
		userCollectionItemMapByCollectionItemID[userCollectionItem.CollectionItemID] = struct{}{}
	}

	// response作成
	collections := make([]*collection, 0, len(collectionItems))
	for _, collectionItem := range collectionItems {
		if _, hasItem := userCollectionItemMapByCollectionItemID[collectionItem.ID]; hasItem {
			collections = append(collections, &collection{
				CollectionID: collectionItem.ID,
				Name:         collectionItem.Name,
				Rarity:       collectionItem.Rarity,
				HasItem:      hasItem,
			})
		}
	}

	weaponList := make([]*weapon, 0)
	skinList := make([]*skin, 0)
	for _, collection := range collections {
		weaponmodel, _ := uu.userColectionRepository.GetWeaponByItemID(collection.CollectionID)
		if weaponmodel != nil {
			weaponList = append(weaponList, &weapon{
				WeaponID:	 weaponmodel.ID,
				Name:        collection.Name,
				Ballet: 	 weaponmodel.Ballet,
				Attack:      weaponmodel.Attack,
				Reload: 	 weaponmodel.Reload,
				Speed:       weaponmodel.Speed,
				Rarity:      collection.Rarity,
			})
		}

		skinmodel ,_ := uu.userColectionRepository.GetSkinByItemID(collection.CollectionID)
		if skinmodel != nil {
			skinList = append(skinList, &skin{
				SkinID: 	  skinmodel.ID,
				Name:         collection.Name,
				Speed:        skinmodel.Speed,
				HitPoint:     skinmodel.HitPoint,
				Rarity:       collection.Rarity,
			})
		}
	}


	return &collectionListResponse{
		WeaponList: weaponList,
		SkinList: skinList,
	}, nil
}

// collectionListResponse ユーザーコレクションリストレスポンス
type collectionListResponse struct {
	WeaponList []*weapon `json:"weaponList"`
	SkinList   []*skin `json:"skinList"`
}

// Collection コレクションアイテム情報詳細
type collection struct {
	CollectionID string `json:"collectionID"`
	Name         string `json:"name"`
	Rarity       int32  `json:"rarity"`
	HasItem      bool   `json:"hasItem"`
}

// weapon コレクションアイテム情報詳細
type weapon struct {
	WeaponID 	 string `json:"WeaponID"`
	Name         string `json:"name"`
	Ballet 	 	 int `json:"Ballet"`
	Attack       int `json:"Attack"`
	Reload 	 	 float32 `json:"Reload"`
	Speed        float32 `json:"Speed"`
	Rarity       int32  `json:"Rarity"`
}

// skin コレクションアイテム情報詳細
type skin struct {
	SkinID		string `json:"skinID"`
	Name        string `json:"name"`
	HitPoint    int `json:"hitpoint"`
	Speed       float32 `json:"Speed"`
	Rarity      int32  `json:"rarity"`
}



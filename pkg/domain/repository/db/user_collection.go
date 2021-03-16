package db

import (
	cm "20dojo-online/pkg/domain/model/collection_item"
	sm "20dojo-online/pkg/domain/model/skin"
	ucm "20dojo-online/pkg/domain/model/user_collection_item"
	wm "20dojo-online/pkg/domain/model/weapon"
)

// UserCollectionItemRepository データ永続化のために抽象化したUserCollectionItemデータ更新周りの処理
type UserCollectionItemRepository interface {
	SelectUserCollectionItemsByUserID(userID string) ([]*ucm.UserCollectionItem, error)
	GetWeaponByItemID(itemID string) (*wm.Weapon, error)
	GetSkinByItemID(itemID string) (*sm.Skin, error)
	GetWeaponByID(itemID string) (*wm.Weapon, error)
	GetSkinByID(itemID string) (*sm.Skin, error)
	GetItemByID(itemID string) (*cm.CollectionItem, error)
}

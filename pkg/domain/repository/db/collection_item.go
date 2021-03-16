package db

import (
	cm "20dojo-online/pkg/domain/model/collection_item"
)

// CollectionItemRepository データ永続化のために抽象化したCollectionItemデータ更新周りの処理
type CollectionItemRepository interface {
	SelectAllCollectionItems() ([]*cm.CollectionItem, error)
	SelectCollectionItemsByPrimaryKeys(collectionItemIDs []string) ([]*cm.CollectionItem, error)
}

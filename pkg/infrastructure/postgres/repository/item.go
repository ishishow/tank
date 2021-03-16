package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"20dojo-online/pkg/domain/repository/db"
	"20dojo-online/pkg/infrastructure/postgres"

	cm "20dojo-online/pkg/domain/model/collection_item"
)

type collectionItemRepositoryImpl struct {
	postgres.SQLHandler
}

// NewCollectionItemRepositoryImpl Userに関するDB更新処理を生成
func NewCollectionItemRepositoryImpl(sqlHandler postgres.SQLHandler) db.CollectionItemRepository {
	return &collectionItemRepositoryImpl{
		sqlHandler,
	}
}

// SelectAllCollectionItems コレクションアイテム全件取得
func (uri collectionItemRepositoryImpl) SelectAllCollectionItems() ([]*cm.CollectionItem, error) {
	rows, err := uri.SQLHandler.Conn.Query("SELECT * FROM item")
	if err != nil {
		return nil, err
	}
	return convertToCollectionItems(rows)
}

// SelectCollectionItemsByPrimaryKeys 複数の主キーに応じたcollectionItemを取得する
func (uri collectionItemRepositoryImpl) SelectCollectionItemsByPrimaryKeys(collectionItemIDs []string) ([]*cm.CollectionItem, error) {
	ids := strings.Join(collectionItemIDs, "','")
	sqlRaw := fmt.Sprintf(`SELECT * FROM item WHERE id IN ('%s')`, ids)
	rows, err := uri.SQLHandler.Conn.Query(sqlRaw)
	if err != nil {
		return nil, err
	}
	return convertToCollectionItems(rows)
}

// convertToCollectionItems rowsデータをCollectionItemsデータへ変換する
func convertToCollectionItems(rows *sql.Rows) ([]*cm.CollectionItem, error) {
	var collectionItems []*cm.CollectionItem
	for rows.Next() {
		collectionItem := cm.CollectionItem{}
		if err := rows.Scan(&collectionItem.ID, &collectionItem.Name, &collectionItem.Rarity); err != nil {
			return nil, err
		}
		collectionItems = append(collectionItems, &collectionItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return collectionItems, nil
}

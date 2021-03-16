package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"20dojo-online/pkg/domain/repository/db"
	"20dojo-online/pkg/infrastructure/postgres"

	cm "20dojo-online/pkg/domain/model/collection_item"
	sm "20dojo-online/pkg/domain/model/skin"
	ucm "20dojo-online/pkg/domain/model/user_collection_item"
	wm "20dojo-online/pkg/domain/model/weapon"
)

type userCollectionItemRepositoryImpl struct {
	postgres.SQLHandler
}

// NewUserCollectionItemRepositoryImpl Userに関するDB更新処理を生成
func NewUserCollectionItemRepositoryImpl(sqlHandler postgres.SQLHandler) db.UserCollectionItemRepository {
	return &userCollectionItemRepositoryImpl{
		sqlHandler,
	}
}

// SelectUserCollectionItemsByUserID user_idと一致する行全て取得
func (uri userCollectionItemRepositoryImpl) SelectUserCollectionItemsByUserID(userID string) ([]*ucm.UserCollectionItem, error) {
	rows, err := uri.SQLHandler.Conn.Query("SELECT * FROM user_item WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return convertToUserCollectionItems(rows)
}

// InsertUserCollectionItemsTx ctx, txを受け取り、user_collection_itemに新規データを全て挿入する
func insertUserCollectionItemsTx(ctx context.Context, tx *sql.Tx, records []*ucm.UserCollectionItem) error {
	insertValue := ""
	for _, record := range records {
		insertValue += fmt.Sprintf("('%s', '%s'),", record.UserID, record.CollectionItemID)
	}
	sqlRaw := fmt.Sprintf(`INSERT INTO user_item(user_id, item_id) VALUES%s`, insertValue)
	sqlRaw = strings.TrimRight(sqlRaw, ",")

	_, err := tx.ExecContext(ctx, sqlRaw)
	return err
}

// SelectUserCollectionItemsByUserID user_idと一致する行全て取得
func (uri userCollectionItemRepositoryImpl) GetWeaponByItemID(itemID string) (*wm.Weapon, error) {
	row := uri.SQLHandler.Conn.QueryRow("SELECT id, ballet, attack, reload, speed FROM weapon WHERE item_id = ?", itemID)
	return convertToWeapon(row)
}

// SelectUserCollectionItemsByUserID user_idと一致する行全て取得
func (uri userCollectionItemRepositoryImpl) GetWeaponByID(itemID string) (*wm.Weapon, error) {
	row := uri.SQLHandler.Conn.QueryRow("SELECT id, ballet, attack, reload, speed FROM weapon WHERE id = ?", itemID)
	return convertToWeapon(row)
}

// SelectUserCollectionItemsByUserID user_idと一致する行全て取得
func (uri userCollectionItemRepositoryImpl) GetItemByID(itemID string) (*cm.CollectionItem, error) {
	row := uri.SQLHandler.Conn.QueryRow("select * from item where id =?", itemID)
	return convertToItem(row)
}

// SelectUserCollectionItemsByUserID user_idと一致する行全て取得
func (uri userCollectionItemRepositoryImpl) GetSkinByItemID(itemID string) (*sm.Skin, error) {
	row := uri.SQLHandler.Conn.QueryRow("SELECT id, hit_point, speed FROM skin WHERE item_id = ?", itemID)
	return convertToSkin(row)
}

// SelectUserCollectionItemsByUserID user_idと一致する行全て取得
func (uri userCollectionItemRepositoryImpl) GetSkinByID(itemID string) (*sm.Skin, error) {
	row := uri.SQLHandler.Conn.QueryRow("SELECT id, hit_point, speed FROM skin WHERE id = ?", itemID)
	return convertToSkin(row)
}

// convertToUserCollectionItems rowsデータをUserCollectionItemsデータへ変換する
func convertToUserCollectionItems(rows *sql.Rows) ([]*ucm.UserCollectionItem, error) {
	var userCollectionItems []*ucm.UserCollectionItem
	for rows.Next() {
		userCollectionItem := ucm.UserCollectionItem{}
		if err := rows.Scan(&userCollectionItem.UserID, &userCollectionItem.CollectionItemID); err != nil {
			return nil, err
		}
		userCollectionItems = append(userCollectionItems, &userCollectionItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userCollectionItems, nil
}

// convertToUserCollectionItems rowsデータをUserCollectionItemsデータへ変換する
func convertToWeapon(row *sql.Row) (*wm.Weapon, error) {
	weapon := wm.Weapon{}
	if err := row.Scan(&weapon.ID, &weapon.Ballet, &weapon.Attack, &weapon.Reload, &weapon.Speed); err != nil {
		return nil, err
	}
	return &weapon, nil
}

// convertToUserCollectionItems rowsデータをUserCollectionItemsデータへ変換する
func convertToWeapons(rows *sql.Rows) ([]*wm.Weapon, error) {
	var weapons []*wm.Weapon
	for rows.Next() {
		weapon := wm.Weapon{}
		if err := rows.Scan(&weapon.ID, &weapon.Ballet, &weapon.Attack, &weapon.Reload, &weapon.Speed); err != nil {
			return nil, err
		}
		weapons = append(weapons, &weapon)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return weapons, nil
}

// convertToUserCollectionItems rowsデータをUserCollectionItemsデータへ変換する
func convertToSkin(row *sql.Row) (*sm.Skin, error) {
	skin := sm.Skin{}
	if err := row.Scan(&skin.ID, &skin.HitPoint, &skin.Speed); err != nil {
		return nil, err
	}
	return &skin, nil
}

// convertToUserCollectionItems rowsデータをUserCollectionItemsデータへ変換する
func convertToSkins(rows *sql.Rows) ([]*ucm.UserCollectionItem, error) {
	var userCollectionItems []*ucm.UserCollectionItem
	for rows.Next() {
		userCollectionItem := ucm.UserCollectionItem{}
		if err := rows.Scan(&userCollectionItem.UserID, &userCollectionItem.CollectionItemID); err != nil {
			return nil, err
		}
		userCollectionItems = append(userCollectionItems, &userCollectionItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userCollectionItems, nil
}

// convertToCollectionItems rowsデータをCollectionItemsデータへ変換する
func convertToItem(row *sql.Row) (*cm.CollectionItem, error) {
	collectionItem := cm.CollectionItem{}
	if err := row.Scan(&collectionItem.ID, &collectionItem.Name, &collectionItem.Rarity); err != nil {
		return nil, err
	}
	return &collectionItem, nil
}

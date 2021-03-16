package repository

import (
	"context"
	"database/sql"
	"log"

	"20dojo-online/pkg/domain/repository/db"
	"20dojo-online/pkg/infrastructure/postgres"

	gm "20dojo-online/pkg/domain/model/gacha_probability"
	um "20dojo-online/pkg/domain/model/user"
	ucm "20dojo-online/pkg/domain/model/user_collection_item"
)

type gachaProbabilityRepositoryImpl struct {
	postgres.SQLHandler
}

// NewGachaProbabilityRepositoryImpl Userに関するDB更新処理を生成
func NewGachaProbabilityRepositoryImpl(sqlHandler postgres.SQLHandler) db.GachaProbabilityRepository {
	return &gachaProbabilityRepositoryImpl{
		sqlHandler,
	}
}

// SelectAllGachaProbabilities gachaProbabilityを全件取得する
func (gri gachaProbabilityRepositoryImpl) SelectAllGachaProbabilities() ([]*gm.GachaProbability, error) {
	rows, err := gri.SQLHandler.Conn.Query("SELECT * FROM gacha_probability")
	if err != nil {
		return nil, err
	}
	return convertToGachaProbabilities(rows)
}

// GachaSave ユーザーアイテムの保存処理,コイン消費処理
func (gri gachaProbabilityRepositoryImpl) GachaSave(ctx context.Context, records []*ucm.UserCollectionItem, user *um.User) error {
	if err := postgres.Transact(ctx, gri.SQLHandler.Conn, func(tx *sql.Tx) error {
		if err := updateUserCoinByPriaryKeyTx(ctx, tx, user); err != nil {
			log.Println(err)
			return err
		}
		if err := insertUserCollectionItemsTx(ctx, tx, records); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// convertToGachaProbabilities rowsデータをGachaProbabilitiesデータへ変換する
func convertToGachaProbabilities(rows *sql.Rows) ([]*gm.GachaProbability, error) {
	var gachaProbabilities []*gm.GachaProbability
	for rows.Next() {
		gachaProbability := gm.GachaProbability{}
		if err := rows.Scan(&gachaProbability.CollectionItemID, &gachaProbability.Ratio); err != nil {
			return nil, err
		}
		gachaProbabilities = append(gachaProbabilities, &gachaProbability)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return gachaProbabilities, nil
}

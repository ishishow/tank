package db

import (
	"context"

	gm "20dojo-online/pkg/domain/model/gacha_probability"
	um "20dojo-online/pkg/domain/model/user"
	ucm "20dojo-online/pkg/domain/model/user_collection_item"
)

// GachaProbabilityRepository データ永続化のために抽象化したGachaProbabilityデータ更新周りの処理
type GachaProbabilityRepository interface {
	SelectAllGachaProbabilities() ([]*gm.GachaProbability, error)
	GachaSave(ctx context.Context, records []*ucm.UserCollectionItem, user *um.User) error
}

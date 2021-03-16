package repository

import (
	"database/sql"

	usm "20dojo-online/pkg/domain/model/user_status"
)

// Create ユーザ登録処理
func (uri userRepositoryImpl) GetUserStatus(userID string) (*usm.UserStatus ,error) {
	row := uri.SQLHandler.Conn.QueryRow("SELECT * FROM user_status WHERE user_id = ?", userID)
	return convertToUserStatus(row)
}

// SelectByAuthToken auth_tokenを条件にUserを取得する
func (uri userRepositoryImpl) UpdateUserStatus(userID string, weaponID string, skinID string) error {
	stmt, err := uri.SQLHandler.Conn.Prepare("UPDATE user_status SET weapon_id=?, skin_id=? WHERE user_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(weaponID, skinID, userID)
	return err
}

// SelectByAuthToken auth_tokenを条件にUserを取得する
func (uri userRepositoryImpl) InitUserStatus(userID string) error {
	stmt, err := uri.SQLHandler.Conn.Prepare("INSERT INTO user_status (user_id, weapon_id, skin_id) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, "w_green", "s_green")

	stmt, err = uri.SQLHandler.Conn.Prepare("INSERT INTO user_item(user_id, item_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, "1")

	stmt, err = uri.SQLHandler.Conn.Prepare("INSERT INTO user_item(user_id, item_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, "6")
	return err
}

// convertToUserCollectionItems rowsデータをUserCollectionItemsデータへ変換する
func convertToUserStatus(row *sql.Row) (*usm.UserStatus, error) {
	userStatus := usm.UserStatus{}
	if err := row.Scan(&userStatus.UserID, &userStatus.WeaponID, &userStatus.SkinID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &userStatus, nil
}

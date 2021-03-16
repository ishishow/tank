package repository

import (
	"context"
	"database/sql"

	"20dojo-online/pkg/domain/repository/db"
	"20dojo-online/pkg/infrastructure/postgres"

	um "20dojo-online/pkg/domain/model/user"
)

type userRepositoryImpl struct {
	postgres.SQLHandler
}

// NewUserRepositoryImpl Userに関するDB更新処理を生成
func NewUserRepositoryImpl(sqlHandler postgres.SQLHandler) db.UserRepository {
	return &userRepositoryImpl{
		sqlHandler,
	}
}

// Create ユーザ登録処理
func (uri userRepositoryImpl) Create(ID, authToken, name string) error {
	stmt, err := uri.SQLHandler.Conn.Prepare("INSERT INTO user (id, auth_token, name, high_score, coin) VALUES (?, ?, ?, 0, 0)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ID, authToken, name)
	return err
}

// SelectByAuthToken auth_tokenを条件にUserを取得する
func (uri userRepositoryImpl) SelectByAuthToken(authToken string) (*um.User, error) {
	row := uri.SQLHandler.Conn.QueryRow("SELECT * FROM user WHERE auth_token = ?", authToken)
	return convertToUser(row)
}

// SelectByAuthToken auth_tokenを条件にUserを取得する
func (uri userRepositoryImpl) SelectByPrimaryKey(userID string) (*um.User, error) {
	row := uri.SQLHandler.Conn.QueryRow("SELECT * FROM user WHERE id = ?", userID)
	return convertToUser(row)
}

// UpdateName ユーザーの名前を更新する
func (uri userRepositoryImpl) Update(record *um.User) error {
	stmt, err := uri.SQLHandler.Conn.Prepare("UPDATE user SET name=?, high_score=?, coin=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.Name, record.HighScore, record.Coin, record.ID)
	return err
}

func (uri userRepositoryImpl) SelectAllPlayingUsers() ([]*um.User, error) {
	rows, err := uri.SQLHandler.Conn.Query("SELECT * FROM user WHERE high_score > 0")
	if err != nil {
		return nil, err
	}
	return convertToUsers(rows)
}

// UpdateUserCoinByPriaryKeyTx ctx, txを受け取り，主キーを条件にcoinを更新する
func updateUserCoinByPriaryKeyTx(ctx context.Context, tx *sql.Tx, record *um.User) error {
	stmt, err := tx.PrepareContext(ctx, "UPDATE user SET coin=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, record.Coin, record.ID)
	return err
}

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*um.User, error) {
	user := um.User{}
	if err := row.Scan(&user.ID, &user.AuthToken, &user.Name, &user.HighScore, &user.Coin); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// convertToUsers rowsデータをUsersデータへ変換する
func convertToUsers(rows *sql.Rows) ([]*um.User, error) {
	var users []*um.User
	for rows.Next() {
		user := um.User{}
		if err := rows.Scan(&user.ID, &user.AuthToken, &user.Name, &user.HighScore, &user.Coin); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

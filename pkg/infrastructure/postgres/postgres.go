package postgres

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Driver名
const driverName = "postgres"

// SQLHandler DB接続ハンドラ
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler DB接続ハンドラ生成
func NewSQLHandler() SQLHandler {

	conn, err := sql.Open(driverName, os.Getenv("DATABASE_URl"))
	if err != nil {
		log.Fatal(err)
	}

	return SQLHandler{
		Conn: conn,
	}
}

// Transact パニックの検出、エラー時にロールバック
func Transact(ctx context.Context, db *sql.DB, txFunc func(tx *sql.Tx) error) (err error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		log.Printf("BeginTx is failed :%v", err)
	}
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("rollback is failed: %v", err)
			}
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("rollback is failed: %v", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				log.Printf("commit is failed: %v", err)
			}
		}
	}()
	return txFunc(tx)
}

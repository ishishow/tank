package redis

import (
	"github.com/go-redis/redis"
)

// CacheHandler キャッシュ情報ハンドラ
type CacheHandler struct {
	Client *redis.Client
}

// NewCacheHandler キャッシュハンドラ生成
func NewCacheHandler() CacheHandler {
	// TODO: 設定ファイル読み込むようにするか.env作る
	// addr := os.Getenv("REDIS_PORT")
	// password := os.Getenv("REDIS_PASSWORD")
	// db := os.Getenv("DB_NUMBER")

	RClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return CacheHandler{
		Client: RClient,
	}
}

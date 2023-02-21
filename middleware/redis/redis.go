package redis

import (
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis"
)

var (
	USER_DB  = 0
	VIDEO_DB = 1
)

var rdb *redis.Client

func InitRDB() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RDB_HOST") + ":6379",
		Password: "",
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("[FATAL] Can not connect to redis.")
		return err
	}
	return nil
}

type RDB_Dao struct {
	db *redis.Client
}

var rdbDao *RDB_Dao
var getOnce sync.Once

func NewInstance() *RDB_Dao {
	getOnce.Do(func() {
		rdbDao = &RDB_Dao{
			db: rdb,
		}

	})
	return rdbDao
}

func (r *RDB_Dao) Select(databse int64) *RDB_Dao {
	r.db.Do("select", databse)
	return r
}

func (r *RDB_Dao) Get(key string) (string, error) {
	return r.db.Get(key).Result()
}

func (r *RDB_Dao) GetRange(key string, start string, end string) ([]string, error) {
	return r.db.ZRangeByScore(key, redis.ZRangeBy{Min: start, Max: end}).Result()

}

func (r *RDB_Dao) GetRevRange(key string, start string, end string) ([]string, error) {
	return r.db.ZRevRangeByScore(key, redis.ZRangeBy{Min: start, Max: end}).Result()
}

func (r *RDB_Dao) RemRange(key string, start string, end string) (int64, error) {
	return r.db.ZRemRangeByScore(key, start, end).Result()
}

func (r *RDB_Dao) Rem(key string, member string) (int64, error) {
	return r.db.ZRem(key, member).Result()
}
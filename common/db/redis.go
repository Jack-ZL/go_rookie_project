package db

import (
	"fmt"
	"github.com/Jack-ZL/go_rookie/config"
	"github.com/go-redis/redis"
	"strconv"

	"time"
)

var Redis *redis.Client
var RedisPrefix string

func init() {
	conf := config.Conf.Redis
	RedisPrefix = fmt.Sprintf("%v", conf["redis_prefix"])
	dbNo, _ := strconv.Atoi(conf["redis_database"].(string))
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v", conf["redis_host"]),
		Password: fmt.Sprintf("%v", conf["redis_auth"]),
		DB:       dbNo,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		fmt.Printf("===Redis==Ping==%v", err.Error())
		panic("Redis==Ping==failed")
	}
	fmt.Printf("====redis=conect=success==%v \n", pong)
}

func RedisSet(key string, value interface{}, expire int) (string, error) {
	res := Redis.Set(RedisPrefix+key, value, time.Duration(expire)*time.Second)
	return res.Result()
}

func RedisGet(key string) (string, error) {
	res, err := Redis.Get(RedisPrefix + key).Result()
	return res, err
}

func RedisTtl(key string) (time.Duration, error) {
	res := Redis.TTL(RedisPrefix+key).Val() / 1000000000
	return res, nil
}

func RedisDel(key string) (int64, error) {
	res := Redis.Del(RedisPrefix + key)
	return res.Result()
}

func RedisExists(key string) (int64, error) {
	res := Redis.Exists(RedisPrefix + key)
	return res.Result()
}

func RedisZRevRange(key string, start, stop int64) ([]string, error) {
	res := Redis.ZRevRange(RedisPrefix+key, start, stop)
	return res.Result()
}

func RedisZAdd(key string, members redis.Z) (int64, error) {
	res := Redis.ZAdd(RedisPrefix+key, members)
	return res.Result()
}

func RedisZIncrBy(key string, increment float64, member string) {
	res := Redis.ZIncrBy(RedisPrefix+key, increment, member)
	fmt.Println(res)
}

// 返回 key 所储存的值的类型。
func RedisType(key string) (string, error) {
	return Redis.Type(RedisPrefix + key).Result()
}

// 设置 key 的过期时间
func RedisPexpireAt(key string, timeT time.Time) (bool, error) {
	return Redis.PExpireAt(RedisPrefix+key, timeT).Result()
}

// 修改 key 的名称
func RedisReName(oldKeyName, newKeyName string) (string, error) {
	return Redis.Rename(RedisPrefix+oldKeyName, RedisPrefix+newKeyName).Result()
}

// 移除给定 key 的过期时间，使得 key 永不过期
func RedisPersist(key string) (bool, error) {
	return Redis.Persist(RedisPrefix + key).Result()
}

// 将当前数据库的 key 移动到给定的数据库 db 当中
func RedisMove(key string, destinationDb int64) (bool, error) {
	return Redis.Move(RedisPrefix+key, destinationDb).Result()
}

// 设置 key 的过期时间
func RedisExpire(key string, expiration time.Duration) (bool, error) {
	return Redis.Expire(RedisPrefix+key, expiration).Result()
}

/**
 * 以毫秒为单位返回 key 的剩余过期时间
 * key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1 。 否则，以毫秒为单位，返回 key 的剩余生存时间。
 */
func RedisPttl(key string) (time.Duration, error) {
	return Redis.PTTL(RedisPrefix + key).Result()
}

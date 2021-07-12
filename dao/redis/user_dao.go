package redis

import (
	"encoding/json"
	"gin_web_demo/models"
	"time"

	"strconv"

	"github.com/go-redis/redis"
)

func SetUserByUserId(user *models.User) (err error) {
	key := getRedisKey(UserKey + strconv.Itoa(int(user.UserId)))
	//json序列化
	u, _ := json.Marshal(user)
	err1 := client.Set(key, u, time.Second*60).Err()
	if err1 != nil {
		return
	}
	return
}

func GetUserByUserId(userId int64) (user *models.User, err error) {
	key := getRedisKey(UserKey + strconv.Itoa(int(userId)))

	val, err := client.Get(key).Result()
	if err == redis.Nil {
		// key does not exists
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	//json反序列化
	json.Unmarshal([]byte(val), &user)
	return
}

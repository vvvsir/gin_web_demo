package redis

// redis key
const (
	Prefix  = "gin_web_demo:" //redis key 前缀
	UserKey = "user:"         //单个用户信息 eg: user:1111
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}

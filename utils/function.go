package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

//字符串类型转化为int类型
func StringToInt(str string) (id int) {
	id, _ = strconv.Atoi(str)
	return id
}

//字符串类型转化为int64类型
func StringToInt64(str string) (id int64) {
	id, _ = strconv.ParseInt(str, 10, 64)
	return id
}

//int64类型转化为字符串类型
func Int64ToString(num int64) (str string) {
	str = strconv.FormatInt(num, 10)
	return str
}

//md5方法
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

const secret = "ginxxxxxxxx"

// encryptPassword 密码加密
func EncryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Hs256
func Hs256(message string, secret string) string {
	hasher := hmac.New(sha256.New, []byte(secret))
	hasher.Write([]byte(message))
	sign := strings.TrimRight(base64.URLEncoding.EncodeToString(hasher.Sum(nil)), "=")
	return sign
}

//字串截取
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

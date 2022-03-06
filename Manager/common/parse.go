package common

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/segmentio/ksuid"
	"math/big"
)

// Sha256哈希
func GetSha256Hash(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	bytes := hash.Sum(nil)

	return hex.EncodeToString(bytes)
}

// 生成uuid
func GenUUID() string {
	return ksuid.New().String()
}

// map型结构体转为json字符串
func MapToString(m map[string]interface{}) (string, error) {
	res, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// 生成随机字符串
func genRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// 生成盐
var SaltLen int = 10

func GenSalt() string {
	return genRandomString(SaltLen)
}

package utils

import (
	"encoding/json"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

func Md5(data interface{}) string {
	return gmd5.MustEncrypt(data)
}

func UUID() string {
	return uuid.New().String()
}

func GetString(data interface{}) string {
	return gconv.String(data)
}

func GetInt(data interface{}) int {
	return gconv.Int(data)
}

func GetInt64(data interface{}) int64 {
	return gconv.Int64(data)
}

func JsonToString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(b)
}

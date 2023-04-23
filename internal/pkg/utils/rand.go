package utils

import (
	"github.com/gogf/gf/v2/util/grand"
)

// RandInt 返回min到max之间的随机整数，支持负数，包含边界，即：[min, max]
func RandInt(min, max int) int {
	return grand.N(min, max)
}

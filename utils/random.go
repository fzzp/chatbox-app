package utils

import (
	"math/rand"
	"strings"
	"time"
)

var myRand *rand.Rand

const characters = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"

func init() {
	myRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt 随机数, 闭区间[min, max]
func RandomInt(min, max int) int {
	return min + myRand.Intn(max-min+1)
}

// RandomString 随机n个字符串
func RandomString(n int) string {
	var sb strings.Builder
	max := len(characters)
	for i := 0; i < n; i++ {
		s := characters[myRand.Intn(max)]
		sb.WriteByte(s)
	}

	return sb.String()
}

package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano()) // Mục đích: mỗi lần chạy chương trình các giá trị random sẽ đều khác nhau
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // Random 1 so ngau nhien trong khoang min - max
}

func RandomString(n int) string { // RandomString sinh ra chuỗi ngẫu nhiên độ dài n
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] //
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "VND"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

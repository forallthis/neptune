package randutils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// RandomNumGen 随机数生成(纯数字)，不能超过9位
func RandomNumGen(digit int32) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	max := int32(math.Pow(10, float64(digit)))
	rNum := fmt.Sprintf("%0"+strconv.Itoa(int(digit))+"v", rnd.Int31n(max))
	return rNum
}

// RandomStringGen 随机字符串生成
func RandomStringGen(n int) string {
	letters := []rune("0123456789ABCDEFGHJKLMNPQRSTUVWXYZ")
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	lettersLen := len(letters)
	for i := range b {
		b[i] = letters[r.Intn(lettersLen)]
	}
	return string(b)
}

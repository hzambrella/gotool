// 产生随机字符串的方法(可以改成接口)

package randStr

import (
	//	"crypto/sha1"
	"fmt"
	//	"io"
	"math/rand"
	"time"
)
/*
func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(Krand(16))
	}
}
*/
// 随机字符串
func Krand(size int) string {
	var ikind int
	kinds, result := [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		ikind = rand.Intn(3)
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

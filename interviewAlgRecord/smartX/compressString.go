package smartx

import "fmt"

// 字符串压缩
// aabbbccddaa_____>a2b3c2d2a2
func compressString(s string) string {

	ch := []byte(s)
	fmt.Println(ch)
	return s
}

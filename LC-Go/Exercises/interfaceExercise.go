package exercises

import (
	"fmt"
	"io"
)

type User struct {
	s        string // 已经读取的数据流
	offerset int    // 读写位置
}

func NewUser(s string) *User {
	return &User{s: s}

}
func (s *User) Len() int {
	return len(s.s) - s.offerset
}
func (s *User) Read(p []byte) (n int, err error) {
	for ; s.offerset < len(s.s) && n < len(p); s.offerset++ {
		c := s.s[s.offerset]
		//    将小写字母转化为大写字母
		if 'a' <= c && c <= 'z' {
			p[n] = c + 'A' - 'a'
		} else {
			p[n] = c
		}
		n++
	}
	if n == 0 {
		return n, io.EOF
	}
	return n, nil
}
func main() {
	s := NewUser("woshigezhu nizhidaoma? hahaha\n")
	buf := make([]byte, s.Len())
	n, err := io.ReadFull(s, buf)
	fmt.Printf("%s\n",buf)
	fmt.Println(n, err)
}

package exercises

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 300
	ptr = &a
   pptr=&ptr

   fmt.Println(a)
   fmt.Println(*ptr)
   fmt.Println(**pptr)
}

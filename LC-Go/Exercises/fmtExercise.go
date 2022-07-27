package exercises

import "fmt"

func main() {
	a,b,c:="",0,false
	//fmt.Scan(&a,&b,&c)
	fmt.Println(a,b,c)

	//fmt.Scanln(&a,&b,&c)
	fmt.Println(a, b, c)

	fmt.Scanf("%4s%d%t",&a,&b,&c)
	fmt.Println(a, b,c)
}

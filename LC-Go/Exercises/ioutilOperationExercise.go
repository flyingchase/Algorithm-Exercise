package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	//reader:=strings.NewReader("woshigezhu hhh\n")
	//b,_:=ioutil.ReadAll(reader)
	//fmt.Println(string(b))
	//
	//
	//fileInfos,err:=ioutil.ReadDir("/Users/qlzhou/Documents/github/Algorithm-Exercise/LC-Go")
	//if err != nil {
	//	errors.New("path is wrong")
	//}
	//for _,info:=range fileInfos{
	//	if info.IsDir() {
	//		fmt.Println(info.Name(),"\\")
	//	}else {
	//		fmt.Println(info.Name(),"\\")
	//	}
	//}
	//
	//const in = "This is used to test the Golang Standard Library \n This is the end of str"
	//scanner:=bufio.NewScanner(strings.NewReader(in))
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	a := "hello"
	b := "hello world"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Contains(b, a))
	fmt.Println(strings.ContainsAny(b, a))

	fmt.Println(strings.Count("woshiwoshi", "wo"))
	fmt.Printf("Fields are : %q\n", strings.Fields(" foo bar baz    "))

	fmt.Println(strings.HasPrefix("Gopher", "Go"))

	fmt.Println(strings.HasSuffix("Java", "a"))

	println(strings.Index("woaiaini", "a"))

	c := []string{"aa", "bb", "cc"}
	wait:=strings.Join(c, ",")
	println(strings.Join(c, ","))
	fmt.Println(reflect.TypeOf(wait))	// string

	sort.Strings(c)
}

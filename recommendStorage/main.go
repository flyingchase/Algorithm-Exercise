package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var errorText []string
	source := "[\"错误\",\"null\",\"exception\",\"time out\",\"bad gateway\",\"系统异常\",\"系统正忙\"]"
	json.Unmarshal([]byte(source), &errorText)
	fmt.Println(errorText)
	// fmt.Println(reflect.TypeOf(errorText))
	fmt.Println(len(errorText))
	fmt.Println(errorText[0])
	fmt.Println(errorText[len(errorText)-1])

}

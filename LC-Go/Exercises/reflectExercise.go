package exercises
import (
	"fmt"
	"reflect"
)
type Student struct {
	Name string
	Sex string
	Age int
}
func main() {
	var s Student= Student {
		Name : "1",
		Sex:"man",
		Age: 100,
	}

	// v 是 s 的值类型
	v:=reflect.ValueOf(s)
	// t 是 结构体内字段类型 即 s 结构体的字段类型
	t:=v.Type()
	kind:=t.Kind()
	if kind==reflect.Struct {
		for i:=0;i<v.NumField();i++ {
			field:=v.Field(i)
			fmt.Printf("name : %s type : %v value : %v \n",t.Field(i).Name,field.Type().Kind(),
				field.Interface())
		}
	}
}

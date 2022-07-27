package exercises

import "fmt"

func main() {

	monster := make([]map[string]string, 2)

	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["name"] = "1"
		monster[0]["secondeName"] = "One"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string, 3)
		monster[1]["firstName1"] = "2_1"
		monster[1]["secondeName1"] = "2_2"
		monster[1]["ThirdName1"] = "2_3"
	}

	newMonster := map[string]string{"firstName2": "3_1",
		"secondName2": "3_2", "thirdName2": "3_3", "fourthName": "3_4"}
	monster=append(monster,newMonster)

	fmt.Println(monster)

}

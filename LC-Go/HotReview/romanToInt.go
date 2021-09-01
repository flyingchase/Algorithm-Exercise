package HotReview
var roman = map[string]int{
	"I":1,
	"V":5,
	"x":10,
	"L":50,
	"C":100,
	"D":500,
	"M":1000,
}

func romanTOInt(s string)int  {
	if s == "" {
		return 0
	}
	num,lastint,total:=0,0,0
	for i := 0; i < len(s);i++ {
		char:=s[len(s)-(i+1):len(s)-i]
		num=roman[char]
		if num < lastint {
			total-=num
		}else {
			total+=num
		}
		lastint=num
	}
	return total
	
}

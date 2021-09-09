package HotReview
/*
二分查找


*/

func mySqrt(x int) int {
	if x==1 {
		return 1
	}
	low,high:=0,x
	var mid,sqr int
	for {
		mid=low+(high-low)>>1
		if low ==mid{
			return mid
		}
		sqr=mid*mid
		if sqr==x {
			return mid
		}
		if sqr>x {
			high=mid
		}
		if sqr<x {
			low=mid
		}
	}
}
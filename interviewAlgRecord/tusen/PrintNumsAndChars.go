package tusen

// // 交替打印数字和字母
// // 使用不带缓存的 channel
// func main() {
//
// 	number, letter := make(chan bool), make(chan bool)
// 	wait := sync.WaitGroup{}
// 	go func() {
// 		i := 1
// 		for {
// 			<-number
// 			fmt.Printf("%d%d", i, i+1)
// 			i += 2
// 			letter <- true
// 		}
// 	}()
// 	wait.Add(1)
// 	go func(wait *sync.WaitGroup) {
// 		str := "abcdefghijklmnopqrstuvwxyz"
// 		i := 0
// 		for {
// 			<-letter
// 			if i >= utf8.RuneCountInString(str) {
// 				wait.Done()
// 				return
// 			}
// 			fmt.Print(str[i : i+2])
// 			i += 2
// 			number <- true
// 		}
// 	}(&wait)
// 	number <- true
// 	wait.Wait()
// }
//

// var (
// 	num  = make(chan int, 1)
// 	char = make(chan int, 1)
// 	wg   sync.WaitGroup
// )
//
// func PrintNums() {
// 	defer wg.Done()
// 	for i := 0; i <= 12; i++ {
// 		// 双循环实现一次打印两个数字
// 		for j := 0; j < 2; j++ {
// 			fmt.Printf("%d", 2*i+j+1)
// 		}
// 		// 1 执行一次打印数字后，num channel 接收元素，阻塞
// 		num <- 1
// 		// 4 再执行一次打印后 char 取走元素，继续执行
// 		<-char
// 	}
// }
// func PrintChars() {
// 	defer wg.Done()
// 	for i := 0; i <= 12; i++ {
// 		for j := 0; j < 2; j++ {
// 			fmt.Printf("%c", 'A'+(2*i+j))
// 		}
// 		// 2 num 取走元素，继续执行
// 		<-num
//
// 		// 3 执行一次打印字符之后，char channel 接收元素，阻塞
// 		char <- 1
// 	}
// }
// func main() {
// 	wg.Add(2)
// 	go PrintNums()
// 	go PrintChars()
// 	wg.Wait()
// }

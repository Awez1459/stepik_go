package main
import "fmt"
func task2(ch chan string, txt string) {
    for i := 0; i < 5; i++ {
        ch <- txt + " "
    }
}
func main() {
	ch := make(chan int)
    go task2(ch, "A")
	for i := 0; i < 5; i++ {
		fmt.Print(<-ch)
	}
}
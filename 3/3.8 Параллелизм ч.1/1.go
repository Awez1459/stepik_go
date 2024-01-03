package main
import "fmt"
func task(ch chan int, N int) {
    ch <- N + 1
}
func main() {
	ch := make(chan int)
    go task(ch, 1)
	fmt.Println(<-ch)
}
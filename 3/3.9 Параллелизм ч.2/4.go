package main
import "fmt"
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
    result := make(chan int)
	go func() {
		defer close(result)
		sum := 0
		for {
			select {
			case arg, ok := <-arguments:
				if !ok {
					result <- sum
					return
				}
				sum += arg
			case <-done:
				result <- sum
				return
			}
		}
	}()
	return result
}
func main() {
	arguments := make(chan int)
    done := make(chan struct{})
    go func() {
        for i := 0; i < 10; i++ {
            arguments <- i
        }
        close(arguments)
    }()
    out := calculator(arguments, done)
    for i := 0; i < 10; i++ {
        fmt.Println(<-out)
    }
}
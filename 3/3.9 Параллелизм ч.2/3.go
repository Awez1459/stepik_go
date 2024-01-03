package main
import "fmt"
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        select {
            case num := <- firstChan:
                out <- num * num
            case num := <- secondChan:
                out <- num * 3
            case <- stopChan:
                return
        }
    }()
    return out
}
func main() {
	firstChan := make(chan int)
    secondChan := make(chan int)
    stopChan := make(chan struct{})
    go func() {
        for i := 0; i < 10; i++ {
            firstChan <- i
        }
        close(firstChan)
    }()
    go func() {
        for i := 0; i < 10; i++ {
            secondChan <- i
        }
        close(secondChan)
    }()
    out := calculator(firstChan, secondChan, stopChan)
    for i := 0; i < 10; i++ {
        fmt.Println(<-out)
    }
}
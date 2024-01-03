package main

import (
	"fmt"  // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"time" // пакет используется для проверки выполнения условия задачи, не удаляйте его
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
    idx := make(chan int)
    val := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			done := make(chan int)
			go func() {
				done <- fn(x1)
			}()
			go func() {
				done <- fn(x2)
			}()

			go func(i int) {
				idx <- i
				val <- (<-done) + (<-done)
			}(i)
		}
	}()
	go func() {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[<-idx] = <-val
		}
		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}

func main() {
	in1 := make(chan int)
    in2 := make(chan int)
    out := make(chan int)
    n := 10
    go merge2Channels(func(x int) int { return x * x }, in1, in2, out, n)
    for i := 0; i < n; i++ {
        fmt.Println(<-out)
    }
}
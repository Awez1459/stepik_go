package main
import "fmt"
func main() {
	done := make(chan bool)
	go func() {
		work()
		done <- true
	}()
	<-done
}
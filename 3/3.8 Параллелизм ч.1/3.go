package main
import "fmt"
func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
    previousValue := ""
    defer close(outputStream)
    for value := range inputStream {
        if value != previousValue {
            outputStream <- value
            previousValue = value
        }
    }
}
func main() {
	input := make(chan string)
    output := make(chan string)
    go removeDuplicates(input, output)
    input <- "aaabbbccc"
	fmt.Println(<-output)
}
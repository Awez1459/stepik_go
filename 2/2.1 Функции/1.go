package main
import "fmt"
func f(text string) {
    fmt.Print(text)
}
func main() {
    var text string
    fmt.Scan(&text)
    f(text)
}
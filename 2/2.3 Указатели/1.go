package main
import "fmt"
func test(x1 *int, x2 *int) {
    fmt.Print(*x1**x2)
}
func main() {
    var x1, x2 int
    _, _ = fmt.Scan(&x1, &x2)
    test(&x1, &x2)
}
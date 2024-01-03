package main
import "fmt"
func vote(x int, y int, z int) int {
    countOnes := x + y + z
    if countOnes >= 2 {
        return 1
    }
    return 0
}
func main() {
    var x, y, z int
    _, _ = fmt.Scan(&x, &y, &z)
    fmt.Println(vote(x, y, z))
}
package main
import "fmt"
func convert(num int64) uint16 {
	return uint16(num)
}
func main() {
    var num int64
    fmt.Scan(&num)
    fmt.Println(convert(num))
}
package main
import "fmt"
func main () {
    var s1, s2 string
    fmt.Scan(&s1)
    for i := 0; i < len(s1); i++{
        if i % 2 == 1{
            s2 += string(s1[i])
        }
    }
    fmt.Print(s2)
}
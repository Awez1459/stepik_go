package main
import (
	"fmt"
	"strings"
)
func main() {
    var s1, s2 string
    fmt.Scan(&s1)
    for i := 0; i < len(s1); i++{
        if strings.Count(s1, string(s1[i]))==1{
            s2 += string(s1[i])
        }
    }
    fmt.Print(s2)
}
package main
import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    var als = []rune(text)
	if unicode.IsUpper(als[0]) && als[len(als)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
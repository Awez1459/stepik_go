package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, ",", ".", -1)
	items := strings.Split(text, ";")
	if len(items) != 2 {
		panic("invalid format, must be two numbers!")
	}
	num1, err := strconv.ParseFloat(items[0], 64)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.ParseFloat(items[1], 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.4f\n", num1/num2)
}
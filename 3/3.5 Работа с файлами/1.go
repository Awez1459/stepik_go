package main
import (
	"bufio"
	"os"
	"strconv"
)
func main() {
	in := bufio.NewScanner(os.Stdin)
	num := 0
	for in.Scan() {
		if len(in.Text()) == 0 {
			break
		}
		n, err := strconv.Atoi(in.Text())
		if err != nil {
			panic(err)
		}

		num += n
	}
	_, _ = os.Stdout.WriteString(strconv.Itoa(num))
}
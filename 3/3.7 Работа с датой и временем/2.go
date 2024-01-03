package main
import (
	"bufio"
	"fmt"
	"os"
	"time"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	layout := "2006-01-02 15:04:05"
	datetime, _ := time.Parse(layout, input)
	if datetime.Hour() < 13 {
		fmt.Println(datetime.Format(layout))
	} else {
		nextDay := datetime.AddDate(0, 0, 1)
		fmt.Println(nextDay.Format(layout))
	}
}
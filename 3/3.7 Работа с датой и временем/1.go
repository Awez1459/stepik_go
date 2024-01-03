package main
import (
	"fmt"
	"time"
)
func main() {
	var yours string
	fmt.Scan(&yours)
	parsedTime, err := time.Parse(time.RFC3339, yours)
	if err != nil {
        fmt.Println(err)
        return
	}
	unixDate := parsedTime.Format("Mon Jan _2 15:04:05 MST 2006")
	fmt.Println(unixDate)
}
package main
import (
	"fmt"
	"time"
)
const now = 1589570165
func main() {
	var minutes, seconds int
	_, _ = fmt.Scanf("%d мин. %d сек.", &minutes, &seconds)
	duration := time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds)
	baseDate := time.Unix(now, 0).UTC()
	newDate := baseDate.Add(duration)
	fmt.Println(newDate.Format(time.UnixDate))
}
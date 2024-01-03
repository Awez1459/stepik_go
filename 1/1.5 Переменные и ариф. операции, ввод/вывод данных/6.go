package main
import ("fmt")
func main() {
	var d int
	fmt.Scan(&d)
	hours := d / 30
	minutes := (d % 30) * 2
	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
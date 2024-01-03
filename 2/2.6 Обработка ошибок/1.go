package main
import (
	"errors"
	"fmt"
)
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка")
	}
	return a / b, nil
}
func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	result, err := divide(num1, num2)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(result)
	}
}
package main
import (
	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	"os"            // пакет используется для проверки ответа, не удаляйте его
)
func main() {
	value1, value2, operation := os.readTask()
    if a, ok := value1.(float64); ok {
        if b, ok := value2.(float64); ok {
            switch operation {
                case "+": fmt.Printf("%.4f", a+b)
                case "-": fmt.Printf("%.4f", a-b)
                case "*": fmt.Printf("%.4f", a*b)
                case "/": fmt.Printf("%.4f", a-b)
                default: 
                fmt.Println("неизвестная операция")
                return
            }     
        } else {
        fmt.Printf("value=%v: %T", value2, value2)
        }
    } else {
    fmt.Printf("value=%v: %T", value1, value1)
    }                                       // все полученные значения имеют тип пустого интерфейса
}
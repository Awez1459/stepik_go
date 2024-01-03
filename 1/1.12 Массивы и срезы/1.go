
package main
import "fmt"
func main() {
  var workArray [10]uint8
  for i := 0; i < 10; i++ {
    fmt.Scan(&workArray[i])
  }
  a, b := 0, 0
  for i := 0; i < 3; i++ {
    fmt.Scan(&a, &b)
    workArray[a], workArray[b] = workArray[b], workArray[a]
  }
  for i := 0; i < len(workArray); i++{
    fmt.Printf("%d ", workArray[i])
  }
}
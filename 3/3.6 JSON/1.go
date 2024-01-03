package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}
type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}
type Result struct {
	Average float64 `json:"Average"`
}
func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var group Group
	if err := json.Unmarshal(data, &group); err != nil {
		panic(err)
	}
	var totalRatings int
	var totalStudents int
	for _, student := range group.Students {
		totalRatings += len(student.Rating)
		totalStudents++
	}
	averageRating := float64(totalRatings) / float64(totalStudents)
	result := Result{
		Average: averageRating,
	}
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resultJSON))
}
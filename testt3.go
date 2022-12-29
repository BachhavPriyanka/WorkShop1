/*Q7) Enhance the Power of 2 function using the math package of GoLang.

  Note: Power function of Math package :
  func Pow(x, y float64) float64


  Write a function which takes a series (slice) of numbers as an input argument
  and returns a series (slice) of numbers which are Power of 2 numbers of
  respective input numbers. This function uses the Power of 2 function developed

  Write a main function where a slice of integers are defined. Compute the power
  of this slice using the above function.

  For example:
  For Input: [10 43 0 8]

  Output will be :
  [1024 16 8 1 256]*/

package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// Q8)Enhance above function to return a map where key is input number and value is
// power of 2 of the input number
func powerInMap(nums []int) {
	powerMap := make(map[int]int)

	for _, value1 := range nums {
		powerStore := int(math.Pow(2, float64(value1)))
		powerMap[value1] = powerStore
	}
	fmt.Println(powerMap)

	//Q10 Convert this map in json
	jsonConversion, err := json.Marshal(powerMap)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(jsonConversion))
	}

}

func findPower(nums []int) []int {
	var powerResult []int
	for _, numb := range nums {
		power := int(math.Pow(2, float64(numb)))
		powerResult = append(powerResult, power)
	}

	return powerResult
}

func main() {
	numbers := []int{10, 43, 0, 8}
	result := findPower(numbers)
	fmt.Println(result)
	powerInMap(numbers)

}

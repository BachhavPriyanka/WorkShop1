/*Q6) Write a function which computes Power of 2 for a number. Function takes an
  integer as an argument and returns power of 2 for that integer.
  Develop this function without using the math package of Golang.

  Write a main function which takes user input and computes the power of 2 using
  the above function and prints the same on terminal in the following way :
  “Power of 9 is 512”

  e.g: 2*2*2*2*2*2*2*2*2  = Will return 512
       2^n


  Test this function with at least few integer numbers like 9, 0, 7, 1, 2*/

package main

import "fmt"

func PowerOf2(num int, pow int) {
	val := 1
	for i := 0; i < pow; i++ {
		val = val * num
	}

	fmt.Printf("Power of %d : %d\n", pow, val)
}

func main() {
	num := 2
	//pow := 4
	pow := []int{9, 0, 7, 1, 2}
	for _, val := range pow {
		PowerOf2(num, val)
	}

}

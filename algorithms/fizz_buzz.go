/*
############################################################

This is the solution for the exercices that algorithmswithgo.com has in the repo https://github.com/joncalhoun/algorithmswithgo.com

Please visit the course that is really useful 
#################################
*/

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.

package main

import (
	 	"fmt"
		"runtime"
)

func main(){
    var memStats runtime.MemStats
	FizzBuzz(5)
    runtime.ReadMemStats(&memStats)
    fmt.Printf("Allocated memory: %v bytes\n", memStats.Alloc)
    fmt.Printf("Total memory: %.5f MB\n", float64(memStats.Sys/1000000))
}


func FizzBuzz(n int)  {
	for i:=0 ; i <= n ; i++ {
		if i%3 == 0 {
		fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3==0 && i%5==0 {
			fmt.Println("Fizz Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
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
)

func main(){
	fmt.Println(dec_to_base(10,2))
}

func dec_to_base(dec,base int) string {
	var result string
    for dec > 0 {
		remand := dec % base
		result = result + fmt.Sprintf("%X",remand)
		dec = dec/base
    }
	return Reverse(result)

}
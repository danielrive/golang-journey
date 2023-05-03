/*
############################################################

This is the solution for the exercices that algorithmswithgo.com has in the repo https://github.com/joncalhoun/algorithmswithgo.com

Please visit the course that is really useful 
#################################
*/

package main

import (
	 	"fmt"
		"runtime"
)

func main(){
    var memStats runtime.MemStats
	test:= []int{4,788,99,5,99,23}	
	fmt.Println(SumRecursion(test))
	
    runtime.ReadMemStats(&memStats)
    fmt.Printf("Allocated memory: %v bytes\n", memStats.Alloc)
    fmt.Printf("Total memory: %v bytes\n", memStats.Sys)
}

func Sum(numbers []int) int {
	result:=0
	for _,value:= range numbers {
		result=result + value
	}
	return result
}

func SumRecursion(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + SumRecursion(numbers[1:])
}

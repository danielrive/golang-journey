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
	fmt.Println(Reverse("catdog"))
    runtime.ReadMemStats(&memStats)
    fmt.Printf("Allocated memory: %v bytes\n", memStats.Alloc)
    fmt.Printf("Total memory: %v bytes\n", memStats.Sys)
}


func Reverse(word string) string {
	var word_reversed string
	for i:= len(word) -1 ; i >= 0; i-- {
		word_reversed = word_reversed + string(word[i])
	}
	return word_reversed
}
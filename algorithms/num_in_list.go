/*
############################################################

This is the solution for the exercices that algorithmswithgo.com has in the repo https://github.com/joncalhoun/algorithmswithgo.com

Please visit the course that is really useful 
#################################
*/

package main

import "fmt"

func main(){
	test:= []int{4,788,99,5,9,6,1,7777,5559,669752,588}
	fmt.Println(NumInList(test,4))
}

func NumInList(list []int, num int) bool {
	for _, i:= range list {
		if num == i {
			return true
		}
	}
 	return false
}

package main


import "fmt"

/*
############################################################

This is the solution for the exercices that algorithmswithgo.com has in the repo https://github.com/joncalhoun/algorithmswithgo.com

Please visit the course that is really useful 
#################################
*/

// FindTwoThatSum will look for two numbers in the provided
// numbers slice that sum up to the sum argument. It will then
// return the indices of each of these numbers.
//
// If there are multiple correct answers, any of them may be
// used. If there are no correct answers, (-1, -1) is returned.
//
// Eg:
//
//   FindTwoThatSum([1,2,3,4], 7) => (2, 3)
//   FindTwoThatSum([0, -1, 1], 0) => (1, 2)
//   FindTwoThatSum([0, 1, 1], 0) => (-1, -1)
//
// Or if we have duplicate answers any of them are okay. All
// of the following are correct.
//
//   FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (5, 6)
//   FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (1, 3)
//   FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (1, 7)
//
// Each number will only be used once, and the original numbers
// list should not be rearranged or altered in any way.
//
// If you want to challenge
// yourself further, try adding different criteria for which
// numbers will be used when there are multiple correct answers.
// Eg:
// 1. Always return the lowest index possible for the first value.
// 2. Always return the index of lowest value possible for the
//    first return value.
// 3. Always return the index of the two values who have a minimal
//    difference. Eg prefer the values 2, 2 over 1, 3 over 0, 4
//    for the sum of 4.
//

func main() {
    fmt.Println("hello")
	slide_to_test := []int{0,1,1}
	fmt.Println(FindTwoThatSum(slide_to_test,0))

}


func FindTwoThatSum(numbers []int, sum int) (int,int) {
		for index:=0; index < len(numbers); index++ {
		temp_slide:= numbers[index+1:]
		for i:=0; i < len(temp_slide); i++ {
			if sum  == numbers[index] + temp_slide[i] {
				return index,i+1+index
			}
		}
	}
	return -1,-1

	

	
}

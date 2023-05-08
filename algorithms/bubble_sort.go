package main

import "fmt"

// {22,67,72,57,90,45,15}
// i = 0
//    j=0    
//      	{22,67}
  //    j=1

func main() {
   test_slide := []int{22,67,66,7,80,91,55,66,7,80,91}
   compare(test_slide)
}

func compare(numbers []int)  {
	for i:=0; i <=len(numbers); i++ {
		for j:=0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j],numbers[j+1] = numbers[j+1],numbers[j]
			}
			fmt.Println(numbers)
			}
	}
}
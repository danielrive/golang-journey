/*
############################################################

This is the solution for the exercices that algorithmswithgo.com has in the repo https://github.com/joncalhoun/algorithmswithgo.com

Please visit the course that is really useful 
#################################
*/


package main

import ("fmt"
		"runtime"
)

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//

func main() {

   var memStats runtime.MemStats
   runtime.ReadMemStats(&memStats)
   //fmt.Println(Fibonacci_recursion(44))
   fmt.Println(Fibonacci(60))
   fmt.Printf("Allocated memory: %v bytes\n", memStats.Alloc)
   fmt.Printf("Total memory: %v MB\n", float64(memStats.Sys)/1000000)
   fmt.Println(Fibonacci_recursion(60))
   fmt.Printf("Allocated memory: %v bytes\n", memStats.Alloc)
   fmt.Printf("Total memory: %v MB\n", float64(memStats.Sys)/1000000)

}

func Fibonacci(n int) int {
	fibonacci_result :=0
    if n <= 1 {
		return n
	}
	nMinus1:=1
	nMinus2:=0
	for i:=2; i <= n; i++ {
         fibonacci_result = nMinus1 + nMinus2
		 nMinus2 = nMinus1
		 nMinus1 = fibonacci_result
	}
	return fibonacci_result
}
func Fibonacci_recursion(n int) int {
	//Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
    if n==0 {
		return 0
	}
	if n==1 {
		return 1
	}
	return Fibonacci_recursion(n-1) + Fibonacci_recursion(n-2)
}
/*
f(0) = 0 
f(1) = 1
       F(N-1) + F(N-2)
f(2) = f(1)   + f(0)     = 0 + 1 = 1
f(3) = f(2)   + f(1)     = 1 + 1 = 2
f(4) = f(3)   + f(2)     = 2 + 1 = 3
f(5) = f(4)   + f(3)     = 3 + 2 = 5
f(6) = f(5)   + f(4)     = 5 + 3 = 8
*/




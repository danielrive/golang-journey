/*
############################################################

This is the solution for the exercices that algorithmswithgo.com has in the repo https://github.com/joncalhoun/algorithmswithgo.com

Please visit the course that is really useful
#################################
*/

package main

import (
	//"encoding/csv"
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var points int
	file, err := os.Open("quiz.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	full_file := csv.NewReader(file)

	records, err2 := full_file.ReadAll()
	if err2 != nil {
		log.Fatal(err2)
	}
	for i := 0; i <= len(records)-1; i++ {
		question := records[i]
		last_item := len(records[i])
		fmt.Printf("Problem # %d: %v = ", i, question[0:last_item-1])
		reader := bufio.NewReader(os.Stdin)
		user_answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			break
		}
		if strings.TrimSpace(user_answer) == strings.TrimSpace(question[last_item-1]) {
			points = points + 1
		}
	}
	fmt.Printf("%d ok of %d in total", points, len(records))

}

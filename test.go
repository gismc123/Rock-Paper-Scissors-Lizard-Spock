package main

import (
	"fmt"       // for formatted I/O like Println
	"strings"   // For string manipulation (trimming, lowercasing)
)

func main() {

	botAnswers := [5]string{"rock","paper","scissors","lizzard","spock"}

	fmt.Println("Choose your pick. Enter the number corresponding to your choice.")
				for i, a := range botAnswers {
					fmt.Print (i,": ", a, "\n")
				}
		answer, _ := reader.ReadString('\n')  // Read input until newline
		text := strings.TrimSpace(strings.ToLower(answer)) // Trim whitespace and convert to lowercase

}
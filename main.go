package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	const MIN int = 1
	const MAX int = 100

	// Init a new rand object
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Init the state of our program
	answer := r1.Intn(MAX+1) + MIN
	guess := -1

	scanner := bufio.NewScanner(os.Stdin)

	// Main game loop
	for guess != answer {
		// Receive user input
		for {
			isCorrect := false

			fmt.Print("Input a number between ", MIN, " and ", MAX, ": ")
			for scanner.Scan() {
				input := scanner.Text()

				n, err := strconv.ParseInt(input, 10, 0)
				num := int(n)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Invalid number: %s\n", input)
				} else if num < MIN || num > MAX {
					fmt.Fprintf(os.Stderr, "Number out of range: %d\n", num)
				} else {
					guess = num
					isCorrect = true
				}

				break
			}

			if isCorrect {
				break
			}
		}

		// Check our guess is correct
		if guess > answer {
			fmt.Println("Too large")
		} else if guess < answer {
			fmt.Println("Too small")
		} else {
			fmt.Println("You guess right")
			break
		}
	}
}

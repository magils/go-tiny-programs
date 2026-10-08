package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func ReadNumber() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Guess the number: ")
		scan, _ := reader.ReadString('\n')
		numberValue, err := strconv.Atoi(strings.TrimSpace(scan))

		if err != nil {
			fmt.Println("Error: Invalid value. Value should be numeric")
		} else {
			return numberValue
		}
	}
}

func GetRandomNumber(randSeed int) int {
	return rand.IntN(randSeed)
}

func ShowHint(userGuess, guessNumber int) {
	distance := userGuess - guessNumber
	absoluteDistance := abs(distance)

	if absoluteDistance <= 6 {
		fmt.Println("HINT: Close!")
	} else if absoluteDistance <= 15 {
		fmt.Println("HINT: Getting close!")
	} else if distance < 0 {
		fmt.Println("HINT: Too far, try higher.")
	} else {
		fmt.Println("HINT: Too far, try lower.")
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	numberToGuess := GetRandomNumber(100)

	for {
		userGuess := ReadNumber()

		if userGuess == numberToGuess {
			fmt.Println("Congratulations! You guessed the number.")
			break
		}

		ShowHint(userGuess, numberToGuess)
	}
}

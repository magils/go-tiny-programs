package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a number: ")
		scanInput, _ := reader.ReadString('\n')
		var err error
		input, err = strconv.Atoi(strings.TrimSpace(scanInput))

		if err != nil {
			fmt.Println("Error: invalid number value")
		} else {
			break
		}
	}

	for i := 1; i < 13; i++ {
		fmt.Printf("%d x %d = %d\n", input, i, (input * i))
	}
}

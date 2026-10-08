package main

import (
	"fmt"
	"os"
	"strconv"
)

func Calculate(initialValue float64, opFunc func(float64, float64) float64, numbers []string) float64 {
	result := initialValue

	for _, v := range numbers {
		n, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf("Error: Invalid value: %v\n", n)
			os.Exit(1)
		}

		result = opFunc(result, n)
	}

	return result
}

func ShowUsage() {
	fmt.Println("Usage: go-calc -op n1 n2 ...")
	fmt.Println()
	fmt.Println("Operations:")
	fmt.Println("  -a    Add numbers")
	fmt.Println("  -s    Subtract numbers")
	fmt.Println("  -m    Multiply numbers")
	fmt.Println("  -d    Divide numbers")
	fmt.Println("  -!    Calculate factorial")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go-calc -a 1 2 3")
	fmt.Println("  go-calc -s 10 2 3")
	fmt.Println("  go-calc -m 2 3 4")
	fmt.Println("  go-calc -d 100 2 5")
	fmt.Println("  go-calc -! 5")
}

func ParseValueToInt(value string) int {
	parsedValue, err := strconv.Atoi(value)

	if err != nil {
		fmt.Printf("Error: Invalid value '%s'\n", value)
		os.Exit(1)
	}

	return parsedValue
}

func ExtractFirstNumber(numbers []string) float64 {
	initialValue, err := strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		fmt.Printf("Error: Invalid value '%s'\n", numbers[0])
		os.Exit(1)
	}
	return initialValue
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: At least 2 arguments are needed.")
		ShowUsage()
		os.Exit(1)
	}

	operation := os.Args[1]
	numbers := os.Args[2:]

	switch operation {
	case "-a":
		opFunc := func(x, y float64) float64 {
			return x + y
		}
		res := Calculate(0, opFunc, numbers)
		fmt.Println(res)
	case "-s":
		opFunc := func(x, y float64) float64 {
			return x - y
		}
		initialValue := ExtractFirstNumber(numbers)
		res := Calculate(initialValue, opFunc, numbers[1:])
		fmt.Println(res)
	case "-m":
		opFunc := func(x, y float64) float64 {
			return x * y
		}
		res := Calculate(1, opFunc, numbers)
		fmt.Println(res)
	case "-d":
		opFunc := func(x, y float64) float64 {
			return x / y
		}
		initialValue := ExtractFirstNumber(numbers)
		res := Calculate(initialValue, opFunc, numbers[1:])
		fmt.Println(res)
	case "-!":
		rangeValue := ParseValueToInt(numbers[0])
		numbers := []string{}

		if rangeValue < 0 {
			fmt.Println("Error: Number cannot be negative")
			os.Exit(1)
		}

		for v := range rangeValue {
			numbers = append(numbers, strconv.Itoa(v+1))
		}

		opFunc := func(x, y float64) float64 {
			return x * y
		}

		res := Calculate(1, opFunc, numbers)
		fmt.Println(res)
	default:
		fmt.Println("Error: Operation not found.")
		ShowUsage()
	}
}

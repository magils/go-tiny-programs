package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func CalculateBirthdateByAge(age int, month time.Month, day int) time.Time {
	today := time.Now()
	birthyear := today.Year() - age

	if !(today.Month() > month || (today.Month() == month && today.Day() >= day)) {
		birthyear--
	}

	return time.Date(birthyear, month, day, 0, 0, 0, 0, time.UTC)
}

func CalculateAgeByBirthdate(birthdate string) (int, error) {
	today := time.Now()
	birthdateToDate, err := time.Parse("2006-01-02", birthdate)

	if err != nil {
		return 0, err
	}

	age := today.Year() - birthdateToDate.Year()

	// Check to subtract 1 year if the month and day has not passed yet
	if today.Month() < birthdateToDate.Month() || (today.Month() == birthdateToDate.Month() && today.Day() < birthdateToDate.Day()) {
		age--
	}

	return age, nil
}

func ShowUsage() {
	fmt.Println("Usage:")
	fmt.Println("   -a <age> <month> <day>   Calculate the birthdate from an age and birth month/day")
	fmt.Println("  age-calculator -b <birthdate>           Calculate the age from a birthdate (format: YYYY-MM-DD)")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("  <age>       Age in years (integer)")
	fmt.Println("  <month>     Birth month as a number, 1-12")
	fmt.Println("  <day>       Birth day of the month, 1-31")
	fmt.Println("  <birthdate> Full birthdate in YYYY-MM-DD format, e.g. 1993-10-22")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  age-calculator -a 33 10 22")
	fmt.Println("  age-calculator -b 1993-10-22")
}

func ParseArgToInt(value string) int {
	intValue, err := strconv.Atoi(value)

	if err != nil {
		fmt.Printf("Error: Invalid value '%s', it should be numeric.", value)
		os.Exit(1)
	}

	return intValue
}

func main() {
	cmdArgs := os.Args

	if len(cmdArgs) < 3 {
		ShowUsage()
		os.Exit(1)
	}

	switch cmdArgs[1] {
	case "-a":
		if len(cmdArgs) < 5 {
			fmt.Println("Error: Missing argurments. Values required (3): <age> <month> <day>")
			os.Exit(1)
		}

		age := ParseArgToInt(cmdArgs[2])
		month := ParseArgToInt(cmdArgs[3])
		day := ParseArgToInt(cmdArgs[4])
		birthdate := CalculateBirthdateByAge(age, time.Month(month), day)
		fmt.Println(birthdate.Format("2006-01-02"))
	case "-b":
		age, err := CalculateAgeByBirthdate(cmdArgs[2])

		if err != nil {
			fmt.Println("Error: Unable to calculate the age with the value provided.")
			fmt.Printf("==Error Detail==\n%s\n", err)
			os.Exit(1)
		}

		fmt.Println(age)
	default:
		ShowUsage()
	}

}

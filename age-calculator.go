package main

import (
	"fmt"
	"time"
)

func CalculateBirthdateByAge(age int) {

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

func main() {
	fmt.Println(CalculateAgeByBirthdate("1993-10-22"))
}

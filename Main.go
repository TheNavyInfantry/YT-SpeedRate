package main

import (
	"errors"
	"fmt"
)

func main() {

	var hours int
	var minutes int
	var speed float64

	fmt.Print("Hour(s): ")
	fmt.Scan(&hours)

	fmt.Print("Minute(s): ")
	fmt.Scan(&minutes)

	fmt.Print("Speed Rate (Ex: 1.25): ")
	fmt.Scan(&speed)

	result, err := calculateResult(hours, minutes, speed)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("You need", result, "minutes to finish this video.")
	}

}

func calculateResult(hour int, minute int, speed float64) (int, error) {

	if minute > 60 {
		return 0, errors.New("Minutes can not be bigger than 60!")
	}

	if hour != 0 {
		hour = 60 * hour

		result := (float64(hour) + float64(minute)) / speed

		return int(result), nil

	}
	result := (float64(hour) + float64(minute)) / speed

	return int(result), nil
}

package main

import "fmt"

func main() {

	fmt.Println("Hour(s): ")
	var hours float64
	fmt.Scanln(&hours)

	fmt.Println("Minute(s): ")
	var minutes float64
	fmt.Scanln(&minutes)

	fmt.Println("Speed Rate (Ex: 1.25): ")
	var speed float64
	fmt.Scanln(&speed)

	var result float64

	if hours == 0 {
		hours = 60 * 0
		result = (hours + minutes) / speed

		var result int = int(result)

		println("You will need ", result, " minutes to finish this video")

	} else {
		hours = 60 * hours
		result = (hours + minutes) / speed

		var result int = int(result)

		println("You will need ", result, " minutes to finish this video")

	}
}

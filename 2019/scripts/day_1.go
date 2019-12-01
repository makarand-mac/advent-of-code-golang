package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var dataPath = "./2019/data/day_1.txt"

func main() {
	firstQuiz()
	secondQuiz()
}

func firstQuiz() {
	totalRequiredFuel := 0
	file, err := os.Open(dataPath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataPoint, err := strconv.Atoi(scanner.Text())

		if err != nil {
			continue
		}

		fuelRequired := (dataPoint / 3) - 2
		totalRequiredFuel = totalRequiredFuel + fuelRequired
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(totalRequiredFuel)

}

func secondQuiz() {
	totalRequiredFuel := 0
	file, err := os.Open(dataPath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataPoint, err := strconv.Atoi(scanner.Text())

		if err != nil {
			continue
		}

		fuelRequiredForModule := 0

		for dataPoint > 0 {

			fuelRequired := (dataPoint / 3) - 2

			if fuelRequired < 1 {
				break
			}

			fuelRequiredForModule = fuelRequiredForModule + fuelRequired
			dataPoint = fuelRequired
		}

		totalRequiredFuel = totalRequiredFuel + fuelRequiredForModule
	}

	println(totalRequiredFuel)
}

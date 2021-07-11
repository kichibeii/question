package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// read file
	file, err := os.Open("docs_extract.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		highBenefit int
		start       int
		end         int
		counter     int
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		counter++

		value, _ := strconv.Atoi(scanner.Text())
		if counter == 1 {
			start = value
			continue
		}

		if value >= start {
			if (value - start) > highBenefit {
				end = value
				highBenefit = end - start
			}
		} else {
			start = value
		}
	}

	fmt.Println("really benefit:", highBenefit)

	return
}

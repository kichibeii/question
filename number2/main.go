package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var fullString string
	file, err := os.Open("string.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mapString := getMapString()
		newStringLeftToRight := ""
		newStringSmallestOder := ""

		fullString = scanner.Text()

		for _, byte := range fullString {
			mapString[string(byte)]++
			// if mapString[string(byte)] == 1 {
			// 	newStringLeftToRight = newStringLeftToRight + string(byte)
			// }
		}

		mapCounterString := make(map[string]int)
		for key, value := range mapString {
			if value > 0 {
				mapCounterString[key] = value
			}
		}

		for _, byte := range fullString {
			// fmt.Println(string(byte), " ", mapString[string(byte)], " ", mapCounterString[string(byte)])
			// assign byte that have value one
			if mapString[string(byte)] == 1 {
				newStringLeftToRight = newStringLeftToRight + string(byte)
				newStringSmallestOder = newStringSmallestOder + string(byte)
			}

			if mapString[string(byte)] > 1 && mapCounterString[string(byte)] == 1 {
				newStringSmallestOder = newStringSmallestOder + string(byte)
			}

			// find left to right
			if mapString[string(byte)] == mapCounterString[string(byte)] && mapCounterString[string(byte)] != 1 {
				newStringLeftToRight = newStringLeftToRight + string(byte)
				mapCounterString[string(byte)]--
			} else if mapString[string(byte)] > mapCounterString[string(byte)] {
				mapCounterString[string(byte)]--
			}
		}

		fmt.Println("key : ", fullString)
		fmt.Println("left to right:", newStringLeftToRight)
		fmt.Println("smallest order:", newStringSmallestOder)
		fmt.Println("===========================")
	}

}

func getMapString() map[string]int {
	return map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
}

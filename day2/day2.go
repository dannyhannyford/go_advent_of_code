package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result, err := read_file("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func read_file(file_name string) (int, error) {
	var sum int
	file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum, err = read_lines_pt_1(scanner, sum)
	
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}


func read_lines_pt_1(scanner *bufio.Scanner, sum int) (int, error){
	maxCubes := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	// sum of ID's of possible games
	game_id := 1
	for scanner.Scan() {
		line := scanner.Text()
		// start at :
		// loop
			// if two digit and >=12 check
			// go until ; or end of line
			// if not possible, continue

		// add possible game_id to sum
		

		game_id++
	}
	return sum, nil
}
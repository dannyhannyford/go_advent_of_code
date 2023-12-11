package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
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
	sum, err = read_lines_pt_2(scanner, sum)
	
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}


func read_lines_pt_1(scanner *bufio.Scanner, sum int) (int, error){
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	maxCubes := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	
	// sum of ID's of possible games
	game_id := 1
	for scanner.Scan() {
		line := scanner.Text()
		start_idx := strings.IndexByte(line, ':')
		// start at :
		cube_substring := line[start_idx+1:]
		cubeSlice := strings.Fields(cube_substring)
		if validBag, err := isValidBag(cubeSlice, maxCubes, logger); err == nil {
			// add possible game_id to sum
			if validBag {
				sum += game_id
				logger.Info("success, add to sum", "game_id", game_id, "curr_sum", sum)
			}
		}
		game_id++
	}
	logger.Info("sum", "sum", sum)
	return sum, nil
}

func read_lines_pt_2(scanner *bufio.Scanner, sum int) (int, error) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	maxCubes := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	minCubes := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}
	
	// sum of ID's of possible games
	game_id := 1
	for scanner.Scan() {
		line := scanner.Text()
		start_idx := strings.IndexByte(line, ':')
		// start at :
		cube_substring := line[start_idx+1:]
		cubeSlice := strings.Fields(cube_substring)
		if validBag, err := isValidBag(cubeSlice, maxCubes, logger); err == nil {
			// add possible game_id to sum
			if validBag {
				sum += game_id
				logger.Info("success, add to sum", "game_id", game_id, "curr_sum", sum)
			}
		}
		game_id++
	}
	logger.Info("sum", "sum", sum)
	return sum, nil
}

func isValidBag(cubeSlice []string, maxCubes map[string]int, logger *slog.Logger) (bool, error) {
	for i := 0; i < len(cubeSlice); i++ {
		if num, err := strconv.Atoi(cubeSlice[i]); err == nil {
			color := cubeSlice[i+1]
			if color[len(color)-1] == ';' || color[len(color)-1] == ',' {
				color = color[0:len(color)-1]
			}
			if num > maxCubes[color] {
				logger.Info("too many cubes", "color", color, "maxCubes", maxCubes[color], "claimed_cubes", num)
				return false, nil
			}
			i++
		}
	}
	return true, nil
}
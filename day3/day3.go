package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

// gear ratios

// not a part number if it's adjacent to anther number
// add up all the part numbers

// any number djacent to a symbol is a "part number

// islands problem

// I matrix
// O
// C
// E

/*
if isDigit(line[i]) {
	num := int(line[i] - '0')
	logger.Info("Is it a num?", "num", num)
}
*/

func build_matrix_pt_1(scanner *bufio.Scanner, sum int) (int, error) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{ReplaceAttr: logger_replace}))
	var matrix [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var lineSlice []string
		for i := 0; i < len(line); i++ {
			lineSlice = append(lineSlice, string(line[i]))
		}
		matrix = append(matrix, lineSlice)
	}
	fmt.Println("before traversal:")
	for i:=0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	matrix = searchOcean(matrix)
	fmt.Println("after traversal:")
	for i:=0; i < len(matrix); i++ {
		logger.Info("pt1", "sum", sum, "matrix", matrix[i])
	}
	return sum, nil
}

func searchOcean(matrix [][]string) [][]string {



	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == "." {
				continue
			}
			if !isDigit(matrix[row][col]) {
			//	slog.Info("searchOcean:", "symbol", matrix[row][col])
				continue
			}
			//slog.Info("explore the Island:", "num", matrix[row][col])
			exploreIsland(matrix, row, col)

		}
	}
	return matrix
}



// typing issues
func exploreIsland(matrix [][]string, row, col int) [][]int {
	islandSum := 0
	found_god := false
	unvisitedNeighbors := [][]int{{}}
	queue := [][]int{{row, col}}

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]
		row, col := currNode[0], currNode[1]
		if isDigit(matrix[row][col]) {
			islandSum *= 10
			newSum, err := strconv.Atoi(matrix[row][col])
			if err !=nil {
				fmt.Println(err)
			}
			fmt.Println(newSum)
			islandSum += newSum
			fmt.Println(islandSum)
		}
		unvisitedNeighbors, found_god = exploreNeighbors(row, col, matrix, found_god)
		fmt.Println(unvisitedNeighbors)
	}


	return queue
}
// typing issues
func exploreNeighbors(row, col int, matrix [][]string, found_god bool) ([][]int, bool){
	if !isValid(matrix, row, col) {
		fmt.Println("invalid")
	}
	if !isDigit(matrix[row][col]) {
		found_god = true
		slog.Info("symbol:", "found_god", found_god)
	}

	// down right
	// down
	if row+1 < len(matrix) {
		
	}
	// down left
	// left
	if col-1 >= 0 {

	}
	// up left
	// up
	if row-1 >= 0 {

	}
	// up right
	// right
	if col+1 < len(matrix[0]) {
		if isDigit(matrix[row][col]) {
			
		}
	}
	return matrix[row][col], found_god
}

func isValid(matrix [][]string, row int, col int) bool {
	// out of bounds
	if row < 0 || row > len(matrix)-1 || col < 0 || col > len(matrix[0])-1 || matrix[row][col] == "." {
		return false
	}
	return true
}

func isDigit(r string) bool {
	return r >= "0" && r <= "9"
}
// ------------------------------------------------------------------------------

func logger_replace(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}

		if a.Key == slog.LevelKey && len(groups) ==0 {
			return slog.Attr{}
		}
		return a
}

func read_lines_pt2(scanner * bufio.Scanner, sum int) (int, error) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("pt2", "sum", sum)
	return sum, nil
}

func read_file(file_name string) (int, error) {
	var sum int

	file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum, err = build_matrix_pt_1(scanner, sum)

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil

}

func main() {
	result, err := read_file("test.txt")
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println(result)
}
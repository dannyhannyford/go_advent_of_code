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

// PT2


func build_matrix_pt_2(scanner *bufio.Scanner, sum int) (int, error) {
	// go until you find a gear
	return sum, nil
}



// --------------------------------------------------------------
func build_matrix_pt_1(scanner *bufio.Scanner, sum int) (int, error) {
	// logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{ReplaceAttr: logger_replace}))
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

	matrix, sum = searchOcean(matrix, sum)
	return sum, nil
}

func searchOcean(matrix [][]string, sum int) ([][]string, int) {
	var nodeMatrix [][]bool
	rows := len(matrix)
	for i := 0; i < rows; i++ {
		cols := len(matrix[i])
		row := make([]bool, cols)
		nodeMatrix = append(nodeMatrix, row)
		fmt.Println(row)
	}
	visitedNodes := &nodeMatrix
	var islandSum int


	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if (*visitedNodes)[row][col] {
				continue
			}
			if matrix[row][col] == "." {
				continue
			}
			if !isDigit(matrix[row][col]) {
			//	slog.Info("searchOcean:", "symbol", matrix[row][col])
				continue
			}
			//slog.Info("explore the Island:", "num", matrix[row][col])
			
			islandSum = exploreIsland(matrix, row, col, visitedNodes)
			sum += islandSum
			if islandSum != 0 {
				slog.Info("Island explored:", "Sum", sum)
			}
		}
	}
	return matrix, sum
}



// typing issues
func exploreIsland(matrix [][]string, row, col int, visitedNodes *[][]bool) (int) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{ReplaceAttr: logger_replace}))
	islandSum := 0
	found_god := false
	var ok bool
	// unvisitedNeighbors := [][]int{{}}
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
			islandSum += newSum
			// logger.Info("Sums:", "islandSum", islandSum, "newSum",newSum )
		}
		row, col, found_god, ok = exploreNeighborsGoRight(row, col, matrix, visitedNodes, found_god)
		// logger.Info("explore neighbors:", "row", row, "col", col, "foundGod", found_god, "islandSum", islandSum)
		if ok {
			queue = append(queue, []int{row, col})
		}
		//logger.Info("postQueue", "new Queue", queue, "foundgod", found_god)
		
	}
	if found_god {
		logger.Info("end of island:", "island #", islandSum)
		return islandSum
	}

	return 0
}
// typing issues
func exploreNeighborsGoRight(row, col int, matrix [][]string, visitedNodes *[][]bool, found_god bool) (int, int, bool, bool){
	// if valid and not a number it's a symbol

	// down right
	if row+1 < len(matrix) && col+1 <len(matrix[0]){
		if !isDigit(matrix[row+1][col+1]) && isValid(matrix, row+1, col+1){
			found_god = true
		}
	}
	// down
	if row+1 < len(matrix) {
		if !isDigit(matrix[row+1][col]) && isValid(matrix, row+1, col){
			found_god = true
		}
	}
	// down left
	if row+1 < len(matrix) && col-1 >= 0 {
		if !isDigit(matrix[row+1][col-1]) && isValid(matrix, row+1, col-1){
			found_god = true
		}
	}
	// left
	if col-1 >= 0 {
		if !isDigit(matrix[row][col-1]) && isValid(matrix, row, col-1){
			found_god = true
		}

	}
	// up left
	if row-1 >= 0 && col-1 >= 0 {
		if !isDigit(matrix[row-1][col-1]) && isValid(matrix, row-1, col-1){
			found_god = true
		}
	}
	// up
	if row-1 >= 0 {
		if !isDigit(matrix[row-1][col]) && isValid(matrix, row-1, col){
			found_god = true
		}

	}
	// up right
	if row-1 >= 0 && col+1 < len(matrix[0]) {
		if !isDigit(matrix[row-1][col+1]) && isValid(matrix, row-1, col+1){
			found_god = true
		}
	}
	// right
	if col+1 < len(matrix[0]) {
		if !isDigit(matrix[row][col+1]) && isValid(matrix, row, col+1) {
			found_god = true
			// slog.Info("symbol:", "found_god", found_god)
		}
		if isDigit(matrix[row][col+1]) {
			(*visitedNodes)[row][col+1] = true
			return row, col+1, found_god, true
		}
	}
	return 0, 0, found_god, false
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
	result, err := read_file("input.txt")
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println("after traversal:", result)
}
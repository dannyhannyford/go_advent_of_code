package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
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

	matrix = islandTravel(matrix)
	fmt.Println("after traversal:")
	logger.Info("pt1", "sum", sum, "matrix", matrix)
	return sum, nil
}

func islandTravel(matrix [][]string) [][]string {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if isValid(matrix, row, col) {

			}
		}
	}
	return matrix
}

func getNeighbors(matrix [][]string, row, col int) [][]string {
	neighbors := make([][]string, 0)
	// down right
	if row-1 >= 0 {

	}
	// down
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

}

func isValid(matrix [][]string, row int, col int) bool {
	// out of bounds
	if row < 0 || row > len(matrix)-1 || col < 0 || col > len(matrix[0])-1 || matrix[row][col] == "." {
		return false
	}
	return true
}

func isDigit(matrix [][]string, row int, col int) bool {

}



func logger_replace(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}

		if a.Key == slog.LevelKey && len(groups) ==0 {
			return slog.Attr{}
		}
		return a
}




func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
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
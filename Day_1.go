package main

import (
	"advent_of_go/trie"
	"bufio"
	"fmt"
	"os"
)

func read_file(file_name string) (int, error) {
	var sum int

	file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}
	// defer is put on a defer stack, called before the function returns
		// defer can call recover to stop a panic and regain control
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum = read_lines_pt_2(scanner, sum)

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}

func read_lines_pt_2(scanner *bufio.Scanner, sum int) int {
	var first, last, double_digit int
	var found_first bool
	var queue []string
	
	trie_check := oneNineTrie()
	string_to_int := string_to_int()
	// runtime.Breakpoint()
	for scanner.Scan() {
		line := scanner.Text()
		// traverse through each char in the line
		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				if len(queue) > 0 {
					queue = queue[:0]
				}
				current := int(line[i] - '0')
				if !found_first {
					first = current
					found_first = true
				}
				last = current
			}
			if !isDigit(line[i]) {
				// put char on queue
				
				if len(queue) > 0 {
					for k := 0; k < len(queue); k++ {
						queue[k] = queue[k]+string(line[i])
						if trie_check.Search(queue[k]) {
							curr_str, _ := dequeue(queue)
							current := string_to_int[curr_str]
							if !found_first {
								first = current
								found_first = true
							}
							last = current
							k--
							continue
						}
						if trie_check.StartsWith(queue[k]) {
							continue
						}
						queue = append(queue[:k], queue[k+1:]...)
						k--
					}
				}
				if trie_check.StartsWith(string(line[i])) {
					queue = append(queue, string(line[i]))
				}
			}
		}
		// reset flag
		found_first = false

		// create double_digit
		first = first*10
		double_digit = first+last

		// add to sum
		sum += double_digit
	}
	return sum
}

func read_lines_pt_1(scanner *bufio.Scanner, sum int) int {
	var first, last, double_digit int
	var found_first bool

	for scanner.Scan() {
		line := scanner.Text()
		// traverse through each char in the line
		for i := 0; i < len(line); i++ {
			// if num
			if isDigit(line[i]) {
				current := int(line[i] - '0')
				if !found_first {
					first = current
					found_first = true
				}
				last = current
			}
		}
		// reset flag
		found_first = false

		// create double_digit
		first = first*10
		double_digit = first+last

		// add to sum
		sum += double_digit
	}
	return sum
}

func dequeue(arr []string) (string, bool) {
	if len(arr) == 0 {
		return "", false
	}
	value := arr[0]
	arr = arr[1:]
	return value, true
}

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func string_to_int() map[string]int{
	string_to_int := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
	return string_to_int
}

func oneNineTrie() *trie.Trie{
	trie := trie.NewTrie()
	trie.Insert("one")
	trie.Insert("two")
	trie.Insert("three")
	trie.Insert("four")
	trie.Insert("five")
	trie.Insert("six")
	trie.Insert("seven")
	trie.Insert("eight")
	trie.Insert("nine")
	return trie
}

func main() {
	result, err := read_file("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(result)
}

// Input
// Output
// Constraints
// Exceptions
	// one number
		// 7 = 77
	// zero numbers?

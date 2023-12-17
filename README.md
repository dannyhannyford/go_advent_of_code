# go_advent_of_code

# cli

`go get -u` = get latest package

# questions 
* what is _

# ...
... spreads out an slice

... as an argument means you can have many arguments
ex: `func myFunc(...int)`

# debug
use delve (dlv)
`dlv debug``
set a breakpoint `runtime.Breakpoint()`

# worker pools
https://gobyexample.com/worker-pools
Firstly, in order to make the most out of parallelization you should utilize worker pools. Here's your calibration method rewritten with worker pools:

package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// We want to spawn as many worker threads as we have logical CPUs.
//
//	If we spawn too many, the overhead of managing the threads will decrease
//	performance.
var WorkerCount = runtime.NumCPU()

func main() {
	now := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("couldn't open input.txt: %s", err)
	}
	defer file.Close()

	// Because we can only have WorkerCount number of line reads at a time,
	// there's no need to use more additional memory
	inputChan := make(chan string, WorkerCount)
	// This is arbitrarily bigger than the input file. It would make more 
	// sense to have this also be set to WorkerCount and have a goroutine
	// that's reading from the file.
	resultChan := make(chan int, 10000)

	var wg sync.WaitGroup
	for x := 0; x < WorkerCount; x++ {
		wg.Add(1)
		go calibrationWorker(&wg, inputChan, resultChan)
	}

	beforeRead := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputChan <- scanner.Text()
	}
	slog.Info("time to read", "time", time.Since(beforeRead))
	close(inputChan)
	wg.Wait()
	close(resultChan)

	sum := 0
	for v := range resultChan {
		sum += v
	}
	slog.Info("time to execute", "time", time.Since(now))
	fmt.Println(sum)
}

func calibrationWorker(wg *sync.WaitGroup, inputChan <-chan string, resultChan chan<- int) {
	defer wg.Done()
	for line := range inputChan {
		firstDigit, lastDigit := -1, -1
		for _, c := range line {
			if digit, err := strconv.Atoi(string(c)); err == nil {
				if firstDigit == -1 {
					firstDigit = digit
				}
				lastDigit = digit
			}
		}
		resultChan <- firstDigit*10 + lastDigit
	}
}

However, this isn't going to make much of a difference in performance. If you look at the output of this you'll notice that the majority of time (~85% in my local test) is spent reading the file while the computation is trivial and fast. You're not going to benefit from parallelization much here. In fact, the overhead of managing the threads might make this slower.

# range in go

# extend v implements(duck typing)
https://stackoverflow.com/questions/38834625/whats-the-difference-between-extends-and-implements-in-typescript
package asg1

import (
	"strings"
)

// var sum int
// Task 1
//-------
// This function should output 0^2 + 1^2 + 2^2 + ... + (|n|)^2
func getSumSquares(n int) int {
	// To Do	
	n = abs(n)
	sum:=0
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
//----------------------------------------------------------------------------------------
// Task 2
//-------
// This function extracts all words ending with endLetter from the string text
// Hints:
// - You may find the strings.Fields method useful.
// - Read about the difference between the types "rune" and "byte", and see how the function is tested.
// byte = 8bits(one Ascii) , Rune = 32bits(Unicode - non-Ascii)
func getWords(text string, endLetter rune) []string {
	// split text to words
	words := strings.Fields(text)
	var result []string

	for _, word := range words {
		if len(word) > 0 {
			// characters
			runes := []rune(word)
			// get last
			lastChar := runes[len(runes)-1]
			if lastChar == endLetter {
				result = append(result, word)
			}
		}
	}
	return result
}

//----------------------------------------------------------------------------------------
// Task 3
//--------
type RegRecord struct {
	studentId  int
	courseName string
}

// This method receives a list of student registration records and should return
// a map that shows the number of students registered per course.
// Note that if a duplicate record appears in the input list, it should not be considered in the count.
func getCourseInfo(records []RegRecord) map[string]int {
	courseCount := make(map[string]int)
	seenRecords := make(map[RegRecord]bool)

	for _, record := range records {
		if !seenRecords[record] {
			seenRecords[record] = true
			courseCount[record.courseName]++
		}
	}
	return courseCount
}

//----------------------------------------------------------------------------------------
// Task 4
//-------
// This method is required to count the occurrences of an input key in a list of integers.
// This should be done in parallel. Each invoked go routine should run the countWorker method on part of the list.
// The communication between the main thread and the workers should be done via channels.
// You can use any way to divide the input list across your workers.
// numThreads will not exceed the length of the array


// inputChan = buffer channel to send elements to worker goroutines
// outChan = channel to receive counts from worker goroutines
func count(list []int, key int, numThreads int) int {
	// ensure number of threads is valid
	if numThreads > len(list) {
		numThreads = len(list)
	}

	inputChan := make(chan int, len(list))
	outChan := make(chan int, numThreads)

	for i := 0; i < numThreads; i++ {
		// but pause till fill input channel
		go countWorker(key, inputChan, outChan)
	}

	// go lightweight thread 
	// fill input channel
	go func() {
		for _, num := range list {
			inputChan <- num
		}
		close(inputChan)
	}()

	totalCount := 0
	for i := 0; i < numThreads; i++ {
		totalCount += <-outChan
	}

	return totalCount
}

// This worker method receives inputs via inputChan, and outputs the number of occurrences to outChan
// Note: The worker does not have any information about the number of inputs it will process, i.e.,
// the method should keep working as long as inputChan is open
func countWorker(key int, inputChan chan int, outChan chan int) {
	count := 0
	for num := range inputChan {
		if num == key {
			count++
		}
	}
	outChan <- count
}

// List: [1, 2, 3, 2, 1, 2, 3, 2, 2]
// Key: 2
// Number of Threads: 3

// inputChan = 9
// outChan = 3

// worker waiting data from inputChan

// not equally distributed

// func count(list []int, key int, numThreads int) int {
// 	if numThreads > len(list) {
// 		numThreads = len(list)
// 	}

// 	outChan := make(chan int, numThreads)
// 	doneChan := make(chan struct{})

// 	workAssignments := divideWorkFairly(list, numThreads)

// 	for _, assignment := range workAssignments {
// 		go countWorker(key, assignment, outChan, doneChan)
// 	}

// 	totalCount := 0
// 	for i := 0; i < numThreads; i++ {
// 		totalCount += <-outChan
// 	}

// 	// wait for all routines to be done
// 	for i := 0; i < numThreads; i++ {
// 		<-doneChan
// 	}

// 	return totalCount
// }
 
// // if 3 -> [[], [], []]
// func divideWorkFairly(list []int, numThreads int) [][]int {
// 	// Round-robin distribution
// 	workAssignments := make([][]int, numThreads)
// 	for i := 0; i < len(list); i++ {
// 		workerIndex := i % numThreads
// 		workAssignments[workerIndex] = append(workAssignments[workerIndex], list[i])
// 	}
// 	return workAssignments
// }

// func countWorker(key int, sublist []int, outChan chan int, doneChan chan struct{}) {
// 	count := 0
// 	for _, num := range sublist {
// 		if num == key {
// 			count++
// 		}
// 	}
// 	outChan <- count
// 	// Signal to indicate done
// 	doneChan <- struct{}{}
// }




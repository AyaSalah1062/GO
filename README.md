# Overview

This package provides solutions for four tasks involving mathematical computations, string manipulations, data structuring, and concurrent programming in Go. Below is a detailed explanation of each task along with usage examples.

## Sum of Squares

This function calculates the sum of squares of all integers from 0 to the absolute value of the input `n`.

### Function Signature

```go
func getSumSquares(n int) int
Example
go
Copy code
fmt.Println(getSumSquares(3)) // Output: 14 (0^2 + 1^2 + 2^2 + 3^2) 

Extract Words Ending with a Specific Letter
This function extracts all words from the input string text that end with the specified endLetter.

Function Signature
go
Copy code
func getWords(text string, endLetter rune) []string
Example
go
Copy code
text := "hello world, welcome to the new world"
endLetter := 'd'
fmt.Println(getWords(text, endLetter)) // Output: ["world", "world"]
Course Registration Information
This function receives a list of student registration records and returns a map that shows the number of students registered per course, ignoring duplicate records.

Struct Definition
go
Copy code
type RegRecord struct {
    studentId  int
    courseName string
}
Function Signature
go
Copy code
func getCourseInfo(records []RegRecord) map[string]int
Example
go
Copy code
records := []RegRecord{
    {studentId: 1, courseName: "Math"},
    {studentId: 2, courseName: "Math"},
    {studentId: 1, courseName: "Science"},
    {studentId: 1, courseName: "Math"}, // Duplicate
}
fmt.Println(getCourseInfo(records)) // Output: map[Math:2 Science:1]
Count Occurrences in Parallel
This function counts the occurrences of a specific key in a list of integers using parallel goroutines. Each goroutine processes a part of the list, and the communication between the main thread and the workers is done via channels.

Function Signature
go
Copy code
func count(list []int, key int, numThreads int) int
Example
go
Copy code
list := []int{1, 2, 3, 2, 1, 2, 3, 2, 2}
key := 2
numThreads := 3
fmt.Println(count(list, key, numThreads)) // Output: 5
Helper Functions and Worker
Divide Work Fairly
Divides the list into sublists for each worker in a round-robin fashion.

go
Copy code
func divideWorkFairly(list []int, numThreads int) [][]int
Count Worker
Worker function that counts occurrences of the key in a sublist.

go
Copy code
func countWorker(key int, sublist []int, outChan chan int, doneChan chan struct{})
Usage
To use the package, import it into your Go project and call the desired functions with the appropriate arguments.

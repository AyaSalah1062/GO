# Overview

This package provides solutions for four tasks involving mathematical computations, string manipulations, data structuring, and concurrent programming in Go. Below is a detailed explanation of each task along with usage examples.

## Sum of Squares
This function calculates the sum of squares of all integers from 0 to the absolute value of the input `n`.

## Extract Words Ending with a Specific Letter
This function extracts all words from the input string text that end with the specified endLetter.

## Course Registration Information
This function receives a list of student registration records and returns a map that shows the number of students registered per course, ignoring duplicate records.

## Count Occurrences in Parallel
This function counts the occurrences of a specific key in a list of integers using parallel goroutines. Each goroutine processes a part of the list, and the communication between the main thread and the workers is done via channels.

## Helper Functions and Worker
Divide Work Fairly
Divides the list into sublists for each worker in a round-robin fashion.

## Usage
To use the package, import it into your Go project and call the desired functions with the appropriate arguments.

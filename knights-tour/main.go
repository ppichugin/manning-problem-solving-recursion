package main

import (
	"fmt"
	"sort"
	"time"
)

// The board dimensions.
const numRows = 9
const numCols = numRows

// Whether we want an open or closed tour.
const requireClosedTour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Offset defines offsets for the knight's movement.
type Offset struct {
	dr, dc int
}

// moveOffsets slice holds all the possible moves that a knight can make.
var moveOffsets []Offset

// numCalls variable will hold the number of times we call the recursive
// function, so we can be impressed with the number of calls the program needs
var numCalls int64

func main() {
	numCalls = 0

	// Initialize the move offsets.
	initializeOffsets()

	// Create the blank board.
	board := makeBoard(numRows, numCols)

	// Try to find a tour.
	start := time.Now()
	board[0][0] = 0
	if findTour(board, numRows, numCols, 0, 0, 1) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Could not find a tour.")
	}
	elapsed := time.Since(start)
	dumpBoard(board)
	fmt.Printf("%f seconds\n", elapsed.Seconds())
	fmt.Printf("%d calls\n", numCalls)
}

// Fill the Offset slice.
func initializeOffsets() {
	moveOffsets = []Offset{
		{-2, -1},
		{-1, -2},
		{+2, -1},
		{+1, -2},
		{-2, +1},
		{-1, +2},
		{+2, +1},
		{+1, +2},
	}
}

// Make a board filled with -1s.
func makeBoard(numRows, numCols int) [][]int {
	board := make([][]int, numRows)
	for r := range board {
		board[r] = make([]int, numCols)
		for c := 0; c < numCols; c++ {
			board[r][c] = unvisited
		}
	}
	return board
}

func dumpBoard(board [][]int) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			fmt.Printf("%02d ", board[r][c])
		}
		fmt.Println()
	}
}

// Try to extend a knight's tour starting at (curRow, curCol).
// Return true or false to indicate whether we have found a solution.
func findTour(board [][]int, numRows, numCols, curRow, curCol, numVisited int) bool {
	numCalls++
	board[curRow][curCol] = numVisited
	if numVisited == numRows*numCols {
		if requireClosedTour {
			for _, offset := range moveOffsets {
				if curRow+offset.dr == 0 && curCol+offset.dc == 0 {
					return true
				}
			}
			board[curRow][curCol] = unvisited
			return false
		}
		return true
	}

	accessibility := countAccessibility(board, numRows, numCols, curRow, curCol)
	moves := make([]Offset, len(moveOffsets))
	copy(moves, moveOffsets)

	sort.Slice(moves, func(i, j int) bool {
		return accessibility[i] < accessibility[j]
	})

	for _, offset := range moves {
		r := curRow + offset.dr
		c := curCol + offset.dc
		if r >= 0 && r < numRows && c >= 0 && c < numCols && board[r][c] == unvisited {
			board[r][c] = numVisited
			if findTour(board, numRows, numCols, r, c, numVisited+1) {
				return true
			}
			board[r][c] = unvisited
		}
	}
	board[curRow][curCol] = unvisited
	return false
}

// countAccessibility is a function that calculates the accessibility of each possible move for the knight.
// The accessibility of a move is defined as the number of unvisited squares that can be reached from the target square of the move.
func countAccessibility(board [][]int, numRows, numCols, curRow, curCol int) []int {
	// Initialize an array to hold the accessibility of each move.
	accessibility := make([]int, len(moveOffsets))

	// Loop through each possible move.
	for i, offset := range moveOffsets {
		// Calculate the target square of the move.
		r := curRow + offset.dr
		c := curCol + offset.dc

		// Check if the move is within the board and leads to an unvisited square.
		if r >= 0 && r < numRows && c >= 0 && c < numCols && board[r][c] == unvisited {
			// If the move is valid, loop through each possible move from the target square.
			for _, offset := range moveOffsets {
				// Calculate the square that can be reached from the target square.
				rr := r + offset.dr
				cc := c + offset.dc

				// Check if the move is within the board and leads to an unvisited square.
				if rr >= 0 && rr < numRows && cc >= 0 && cc < numCols && board[rr][cc] == unvisited {
					// If the move is valid, increment the accessibility of the original move.
					accessibility[i]++
				}
			}
		} else {
			// If the original move is not valid, set its accessibility to -1.
			accessibility[i] = -1
		}
	}

	// Return the array of accessibility values.
	return accessibility
}

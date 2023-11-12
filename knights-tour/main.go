package main

import (
	"fmt"
	"time"
)

// The board dimensions.
const numRows = 8
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
		// We've visited every square.
		if requireClosedTour {
			// We need to end at a square that is one knight's move from the start.
			for _, offset := range moveOffsets {
				if curRow+offset.dr == 0 && curCol+offset.dc == 0 {
					// We found a closed tour.
					return true
				}
			}
			// We didn't find a closed tour.
			board[curRow][curCol] = unvisited
			return false
		}
		// We found a tour.
		return true
	}
	// Try extending the tour to each valid move.
	// Loop through the possible moves.
	for _, offset := range moveOffsets {
		// Get the move.
		r := curRow + offset.dr
		c := curCol + offset.dc

		// See if this move is on the board.
		if r < 0 || r >= numRows {
			continue
		}
		if c < 0 || c >= numCols {
			continue
		}

		// See if we have already visited this position.
		if board[r][c] != unvisited {
			continue
		}

		// The move to [r][c] is viable.
		// Make this move.
		board[r][c] = numVisited

		// Try to find the rest of a tour.
		// If we succeed, return true.
		if findTour(board, numRows, numCols, r, c, numVisited+1) {
			return true
		}

		// We did not find a tour with this move. Unmake this move.
		board[r][c] = unvisited
	}
	// Backtrack.
	board[curRow][curCol] = unvisited
	return false
}

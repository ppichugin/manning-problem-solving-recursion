package main

import (
	"fmt"
	"time"
)

func main() {
	const numRows = 20
	board := makeBoard(numRows)

	start := time.Now()
	//success := placeQueens1(board, numRows, 0, 0)
	success := placeQueens4(board, numRows, 0)

	elapsed := time.Since(start)
	if success {
		fmt.Println("Success!")
		dumpBoard(board)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}

/*
Success!
Q . . . . . . . . . . . . . . . . . . .
. . . Q . . . . . . . . . . . . . . . .
. Q . . . . . . . . . . . . . . . . . .
. . . . Q . . . . . . . . . . . . . . .
. . Q . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . Q .
. . . . . . . . . . . . . . . . Q . . .
. . . . . . . . . . . . . . Q . . . . .
. . . . . . . . . . . Q . . . . . . . .
. . . . . . . . . . . . . . . Q . . . .
. . . . . . . . . . . . . . . . . . . Q
. . . . . . . Q . . . . . . . . . . . .
. . . . . Q . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . Q . .
. . . . . . Q . . . . . . . . . . . . .
. . . . . . . . . . . . Q . . . . . . .
. . . . . . . . . . Q . . . . . . . . .
. . . . . . . . Q . . . . . . . . . . .
. . . . . . . . . . . . . Q . . . . . .
. . . . . . . . . Q . . . . . . . . . .
Elapsed: 2.700308 seconds
*/

// Make a board filled with periods.
func makeBoard(numRows int) [][]string {
	numCols := numRows
	board := make([][]string, numRows)
	for r := range board {
		board[r] = make([]string, numCols)
		for c := 0; c < numCols; c++ {
			board[r][c] = "."
		}
	}
	return board
}

// Display the board.
func dumpBoard(board [][]string) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			fmt.Printf("%s ", board[r][c])
		}
		fmt.Println()
	}
}

// Return true if this series of squares contains at most one queen.
func seriesIsLegal(board [][]string, r0, c0, dr, dc int) bool {
	numRows := len(board)
	numCols := numRows
	hasQueen := false

	r := r0
	c := c0
	for {
		if board[r][c] == "Q" {
			// If we already have a queen on this row,
			// then this board is not legal.
			if hasQueen {
				return false
			}

			// Remember that we have a queen on this row.
			hasQueen = true
		}

		// Move to the next square in the series.
		r += dr
		c += dc

		// If we fall off the board, then the series is legal.
		if r >= numRows ||
			c >= numCols ||
			r < 0 ||
			c < 0 {
			return true
		}
	}
}

// Return true if the board is legal.
func boardIsLegal(board [][]string) bool {
	numRows := len(board)

	// See if each row is legal.
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, 0, 0, 1) {
			return false
		}
	}

	// See if each column is legal.
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, 0) {
			return false
		}
	}

	// See if diagonals down to the right are legal.
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, 0, 1, 1) {
			return false
		}
	}
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, 1) {
			return false
		}
	}

	// See if diagonals down to the left are legal.
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, numRows-1, 1, -1) {
			return false
		}
	}
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, -1) {
			return false
		}
	}

	// If we survived this long, then the board is legal.
	return true
}

// Return true if the board is legal and a solution.
func boardIsASolution(board [][]string) bool {
	// See if it is legal.
	if !boardIsLegal(board) {
		return false
	}

	// See if the board contains exactly numRows queens.
	numRows := len(board)
	numQueens := 0
	for r := 0; r < numRows; r++ {
		for c := 0; c < numRows; c++ {
			if board[r][c] == "Q" {
				numQueens += 1
			}
		}
	}
	return numQueens == numRows
}

// Try placing a queen at position [r][c].
// Return true if we find a legal board.
func placeQueens1(board [][]string, numRows, r, c int) bool {
	// If we are past the end of the board, then see if this is a solution.
	if r >= numRows {
		return boardIsASolution(board)
	}

	// Try placing a queen in each column in this row.
	for c := 0; c < numRows; c++ {
		board[r][c] = "Q"
		if placeQueens1(board, numRows, r+1, c) {
			return true
		}
		board[r][c] = "."
	}

	// If we get here, then we could not find a solution.
	return false
}

// Try to place a queen in this column.
// Return true if we find a legal board.
func placeQueens4(board [][]string, numRows, c int) bool {
	// a. If c == numRows, then we have assigned a queen to every column.
	// Call the boardIsLegal function and return true or false accordingly.
	if c == numRows {
		return boardIsLegal(board)
	}

	// b. If c < numRows, then call boardIsLegal to see whether the solution so far is still legal.
	// If it is not, return false.
	if !boardIsLegal(board) {
		return false
	}

	// c. If we reach this point, then we need to assign a queen to column c.
	// Loop through the possible rows with a for r := 0; r < numRows; r++ loop.
	for r := 0; r < numRows; r++ {
		// i. Set the board entry to Q.
		board[r][c] = "Q"

		// ii. Call placeQueens4 recursively to assign a queen to the next column.
		if placeQueens4(board, numRows, c+1) {
			// iii. If the recursive call returns true, then return true.
			return true
		}

		// iv. If the recursive call returns false, then reset board[r][c] to a period and continue the for loop.
		board[r][c] = "."
	}

	// d. If the for loop ends and we did not find a valid row where we could place this column’s queen,
	// then there is no solution starting from the current board arrangement.
	// Return false to backtrack and try again higher up in the call stack.
	return false
}

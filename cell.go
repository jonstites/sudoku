package sudoku

import (
	"fmt"
)


// Represents one square of a Sudoku puzzle.
type cell struct {
	// The determined value
	value uint

	// Stores which values the cell is allowed to be
	options bitarray

	// Keep track of dependent guess when value is set
	guess int

	// any failed guesses in this cell
	tried bitarray
}

// Set the cell value 
func (myCell *cell) setValue(value uint, guessNum int) {
	myCell.value = value
	myCell.options = newBitArray(value)
	myCell.guess = guessNum
}

// add to tried values
func (myCell *cell) setTriedValue(value uint) {
	myCell.tried = setBitTrue(myCell.tried, value)
}

// Set cell options, excluding any previously-tried values
func (myCell *cell) setOptions(options bitarray) {
	myCell.options = options & bitNot(myCell.tried)
}

// If the cell depended on an old guess, reset the cell
func (myCell *cell) reset(guess int) {
	if myCell.guess > guess {
		myCell.value = 0
	}
}

// Check if the cell value has been set
func (myCell *cell) isKnown() bool {
	return myCell.value != 0
}

// Print the value of a cell f
func (myCell *cell) String() string {
	return fmt.Sprint(myCell.value)
}
// Make a new cell with all possible values
func newCell() *cell {
	cell := new(cell)
	cell.options = allTrue()
	return cell
}

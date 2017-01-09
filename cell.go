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
	guess uint

	// any failed guesses in this cell
	tried bitarray
}

// Set the cell value 
func (myCell *cell) setValue(value uint, guessNum uint) {
	myCell.value = value
	myCell.options = newBitArray(value)
	myCell.guess = guessNum
}

// add to tried values
func (myCell *cell) setTriedValue(value uint) {
	myCell.tried = setBitTrue(myCell.tried, value)
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

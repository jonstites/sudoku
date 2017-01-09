package sudoku


// The value for a 9 bit array all set to true
const allDigits = 511

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

func (myCell *cell) setValue(value uint, guessNum int) {
	myCell.value = value
	myCell.options = newBitArray(value)
	myCell.guess = guessNum
}


func (myCell *cell) addGuessValue(value uint) {
	myCell.tried = myCell.tried | newBitArray(value)
}

/*

// Just print the value of a cell f
func (c *cell) String() string {
	valueString := strconv.Itoa(c.value)
	if !(c.valueKnown) {
		valueString = "0"
	}
	return valueString
}
// Make a new cell with all possible values
func newCell() *cell {
	cell := new(cell)
	cell.options = allDigits
	return cell
}
*/

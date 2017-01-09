package sudoku

// The value for a 9 bit array all set to true
const allDigits = 511

type square interface {
}

// Represents one square of a Sudoku puzzle.
type cell struct {
	// The determined value
	value uint

	// Stores which values the cell is allowed to be
	options uint

	// Keep track of dependent guess when value is set
	guess int

	// any failed guesses in this cell
	tried uint
}

func (myCell *cell) setValue(value uint, guessNum int) {
	myCell.value = value
	myCell.options = createBits(value)
	myCell.guess = guessNum
}

/*
func (c *cell) addGuess(value uint) {
	c.guesses = c.guesses | addBits(value)
}


func (c *cell) hasOptions(options ...uint) (bool, error) {
	cellOptions := c.options
	for _, option := range options {
		// check if bit is unset
	
			return false
		}
	}		
	return true
}
		
// Returns the number of values the cell is allowed to be in
func (c *cell) numValueOptions() int {
	num := 0
	for options := c.options; options != 0; options = options >> 1 {
		if (options & 1) == 1 {
			num += 1
		}
	}
	return num
}


func (c *cell) chooseValue() int {
	
	i := 1
	for options := c.options; options != 0; options = options >> 1 {
		if (options & 1) == 1 {
			return i
		}
		i++
	}
	return i
}

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
// Create a bit array with given values
// Offset by 1 because there is no 0 in sudoku
func createBits(values ...uint) uint {
	var bits uint
	bits = 0
	for _, value := range values {
		bits += 1 << (value - 1)
	}
	return bits
}

func valueSet(bitarray uint, value uint) bool {
	return (bitarray >> (value - 1) & 1) == 1
}


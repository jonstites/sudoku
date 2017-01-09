package sudoku


type bitarray uint

// Create a bit array with given values
// Offset by 1 because there is no 0 in sudoku
func newBitArray(values ...uint) bitarray {
	var bits bitarray
	bits = 0
	for _, value := range values {
		bits += 1 << (value - 1)
	}
	return bits
}


// Return true if the value is set to true
func (myBits *bitarray) valueSet(value uint) bool {
	return (*myBits >> (value - 1) & 1) == 1
}


// Return the number of values set to true
func (myBits *bitarray) numValuesSet() int {
	num := 0
	for i := uint(1); i <= 9; i++ {
		if myBits.valueSet(i) {
			num += 1
		}
	}
	return num
}

// Return the lowest value set to true
func (myBits *bitarray) lowestValue() uint {
	for i := uint(1); i <= 9; i++ {
		if myBits.valueSet(i) {
			return i
		}
	}
	return 0
}

package sudoku


type bitarray uint


// The value for a 9 bit array all set to true
const allDigits = 511

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
func (myBits *bitarray) valueTrue(value uint) bool {
	return (*myBits >> (value - 1) & 1) == 1
}

// Return the number of values set to true
func (myBits *bitarray) numValuesTrue() int {
	num := 0
	for i := uint(1); i <= 9; i++ {
		if myBits.valueTrue(i) {
			num += 1
		}
	}
	return num
}

// Return the lowest value set to true
func (myBits *bitarray) lowestValue() uint {
	for i := uint(1); i <= 9; i++ {
		if myBits.valueTrue(i) {
			return i
		}
	}
	return 0
}


// Set value to true
func setBitTrue(myBits bitarray, value uint) bitarray {
	return myBits | newBitArray(value)
}

// Set value to false
func setBitFalse(myBits bitarray, value uint) bitarray {
	return myBits &^ newBitArray(value)
}

// A bitarray with everything set to true
func allTrue() bitarray {
	return allDigits
}

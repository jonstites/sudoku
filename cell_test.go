package sudoku

import (
	"testing"
)


//defaultCell := Cell{bitArray: 511}

func intInSlice(values []uint, target uint) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func TestBitArray(t *testing.T) {
	var bitsTest = []struct {
		bits uint
		set []uint
	}{
		{0, []uint {}},
		{1, []uint {1}},
		{5, []uint {1, 5}},
		{30, []uint {2, 3, 4}},
		{128, []uint {7}},
		{createBits(2, 5, 9), []uint {2, 5, 9}},
	}

	for _, bits := range bitsTest {
		for i := uint(1); i <= 9; i++ {
			if valueSet(bits.bits, i) != intInSlice(bits.set, i) {
				t.Error("Bits: ", bits.bits, "value:", i, "does not match expected.")
			}
		}
	}
}
/*
func TestValueSet(t *testing.T) {
	value := newValueSet(1, 2, 7)
	var expected valueSet
	expected = 67
	if value != expected {
		t.Error("ValueSet of 1,2,7 should be", expected, "not: ", value)
	}
}


func TestConstructor(t *testing.T) {
	for i := 1; 
}


func TestCellPrint(t *testing.T) {
	myCell := Cell{3, 0, true, 0, 0}
	if !(myCell.String() == "3") {
		t.Error("Cell{3} should print as 3.")
	}
}

func TestCellPrintEmpty(t *testing.T) {
	myCell := newCell()
	if !(myCell.String() == "0") {
		t.Error("Unknown Cell should print as 0.")
	}
}

func TestCellOptions(t *testing.T) {
	myCell := Cell{3, newValueSet(3, 5), true, 0, 0}
	hasOptions, err := myCell.hasOptions(3, 5)
	if err != nil {
		t.Error("3 and 5 should be acceptable sudoku values")
	}
	if !hasOptions {
		t.Error("Cell should be allowed to have values 3 and 5")
	}
	var i uint
	for i = 1; i <= 9; i++ {
		hasOptions, err = myCell.hasOptions(i)
		if err != nil {
			t.Error(i, " should be an acceptable sudoku value.")
		}
		if i != 3 && i != 5 && hasOptions {
			t.Error("Cell of 3, 5 should not have option ", i)
		}
	}
}

func TestCellOptionErr(t *testing.T) {
	myCell := Cell{3, newValueSet(3, 5), true, 0, 0}
	_, err := myCell.hasOptions(3, 5, 7, 11)
	if err != nil {
		t.Error("Should raise error when testing for value 11")
	}
}

func TestCellNumOptions(t *testing.T) {
	myCell := Cell{3, newValueSet(3, 5), true, 0, 0}
	if (myCell.numValueOptions() != 2) {
		t.Error("Cell should have 2 options")
	}

}

func TestCellValueOptions(t *testing.T) {
	myCell := newCell()
	var i uint
	for i = 1; i <= 9; i++ {
		hasOption, _ := myCell.hasOptions(i)
		if !(hasOption) {
			t.Error("Cell should have options 1-9")
		}
	}
}

func TestCellChooseValue(t *testing.T) {
	myCell := newCell()
	myCell.valueOptions = 28
	value := myCell.chooseValue()
	expected := 3
	if value != expected {
		t.Error("Expected: ", expected, " got: ", value)
	}
}
*/

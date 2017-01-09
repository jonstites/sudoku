package sudoku


import (
	"testing"
)


//defaultCell := Cell{bitArray: 511}

func TestConstructor(t *testing.T) {
	myCell := newCell()
	for i := uint(1); i <= 9; i++ {
		expected := true
		got := myCell.options.valueTrue(i)
		if got != expected {
			t.Error("Expected %q set to %q but got %q.",
				i, expected, got)
		}
	}
}

/*
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

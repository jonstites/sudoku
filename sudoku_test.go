package sudoku

import (
	"testing"
	"reflect"
)


func TestValueSet(t *testing.T) {
	value := newValueSet(1, 2, 7)
	var expected valueSet
	expected = 67
	if value != expected {
		t.Error("ValueSet of 1,2,7 should be", expected, "not: ", value)
	}
}

func TestCellPrint(t *testing.T) {
	myCell := Cell{3, 0, true}
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
	myCell := Cell{3, newValueSet(3, 5), true}
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
	myCell := Cell{3, newValueSet(3, 5), true}
	_, err := myCell.hasOptions(3, 5, 7, 11)
	if err != nil {
		t.Error("Should raise error when testing for value 11")
	}
}

func TestCellNumOptions(t *testing.T) {
	myCell := Cell{3, newValueSet(3, 5), true}
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

func TestPuzzleNumCellCols(t *testing.T) {
	myPuzzle := newPuzzle()
	if (len(myPuzzle.puzzle) != 9) {
		t.Error("Puzzle should have 9 columns")
	}
}

func TestPuzzleNumCellRows(t *testing.T) {
	myPuzzle := newPuzzle()
	if (len(myPuzzle.puzzle[0]) != 9) {
		t.Error("Puzzle should have 9 rows")
	}
}

func TestPuzzleCellNumOptions(t *testing.T) {
	myPuzzle := newPuzzle()
	numOptions := myPuzzle.puzzle[3][5].numValueOptions()
	if (numOptions != 9) {
		t.Error("Puzzle cells should have 9 options, not ", numOptions)
	}
}

func TestPuzzleSetValue(t *testing.T) {
	myPuzzle := newPuzzle()
	myPuzzle.setValue(3, 2, 8)
	if (myPuzzle.puzzle[3][2].value != 8) {
		t.Error("Puzzle should have 8 set at 3,2.")
	}
	if (myPuzzle.puzzle[3][2].valueKnown != true) {
		t.Error("Puzzle should be set at 3,2.")
	}

}

func TestIsNumeric(t *testing.T) {
	t1 := "00003897509120398"
	t2 := "09287308a98098"
	if !(isNumeric(t1)) {
		t.Error(t1, " should be marked numeric.")
	}

	if isNumeric(t2) {
		t.Error(t2, " should not be marked numeric.")
	}
}

func TestValidateRowFormat(t *testing.T) {
	row1 := "123456788"
	row2 := "12345678"
	row3 := "12345678a"
	row4 := "1234567899"

	if validateRowFormat(row1) != nil {
 		t.Error("Row should be valid: ", row1)
	}

	if validateRowFormat(row2) == nil {
		t.Error("Row should not be valid: ", row2)
	}

	if validateRowFormat(row3) == nil {
		t.Error("Row should not be valid: ", row3)
	}

	if validateRowFormat(row4) == nil {
		t.Error("Row should not be valid: ", row4)
	}
}

func TestGetValue(t *testing.T) {
	myPuzzle := newPuzzle()
	myPuzzle.setValue(3, 2, 8)
	value, ok := myPuzzle.getValue(3, 2)
	if ok != nil {
		t.Error("")
	}

	if value != 8 {
		t.Error()
	}

	value, ok = myPuzzle.getValue(3, 3)
	if ok == nil {
		t.Error()
	}
}

func TestInsertRow(t *testing.T) {
	myPuzzle := newPuzzle()
	row := "123456789"
	myPuzzle.insertRow(1, row)
}

func TestPuzzleFromFile(t *testing.T) {
	puzzleOne := newPuzzle()
	puzzleOne.setValue(2, 1, 2)
	puzzleTwo := puzzleFromFile("almost_empty_test.txt")
	if !(reflect.DeepEqual(puzzleOne, puzzleTwo)) {
		t.Error("Puzzles should be equal.", puzzleOne, puzzleTwo)
	}
	
	
}


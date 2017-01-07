package sudoku

import (
	"testing"
	"reflect"
)


func TestUnion(t *testing.T) {
	value1 := valueSet{1: true, 2:true}
	value2 := valueSet{1: true, 3:true}
	value3 := union(value1, value2)
	if !(reflect.DeepEqual(value3, valueSet{1:true, 2:true, 3:true})) {
		t.Error("Union of {1,2} and {1,3} should be {1,2,3}")
	}
}

func TestIntersection(t *testing.T) {
	value1 := valueSet{1: true, 2:true}
	value2 := valueSet{1: true, 3:true}
	value3 := intersection(value1, value2)
	if !(reflect.DeepEqual(value3, valueSet{1:true})) {
		t.Error("Intersection of {1,2} and {1,3} should be {1}")
	}
}

func TestCellPrint(t *testing.T) {
	myCell := Cell{3, valueSet{}, true}
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
	myCell := Cell{3, map[int]bool{3: true, 5: true}, true}
	if !(myCell.valueOptions[3]) {
		t.Error("Cell should be allowed to have value 3")
	}
}

func TestCellNumOptions(t *testing.T) {
	myCell := Cell{3, map[int]bool{3: true, 5: true}, true}
	if (myCell.numValueOptions() != 2) {
		t.Error("Cell should have 2 options")
	}

}

func TestCellValueOptions(t *testing.T) {
	myCell := newCell()
	for i := 1; i <= 9; i++ {
		if !(myCell.valueOptions[i]) {
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

/*
func TestPuzzleFromFile(t *testing.T) {
	puzzleOne := newPuzzle()
	puzzleOne.puzzle[3][2].value = 2
	puzzleTwo := puzzleFromFile("almost_empty.txt")
	if !(reflect.DeepEqual(puzzleOne, puzzleTwo)) {
		t.Error("Puzzles should be equal.")
	}
	
	
}

*/

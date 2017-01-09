package sudoku
/*
import (
	"testing"
	"reflect"
)



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

func TestPuzzleGetCell(t *testing.T) {
	myPuzzle := newPuzzle()
	myPuzzle.puzzle[3][5].value = 5
	myCell := myPuzzle.getCell(3, 5)
	if myCell.value != 5 {
		t.Error("Function getCell should be getting a cell with value of 5.")
	}
}


func TestPuzzleCellNumOptions(t *testing.T) {
	myPuzzle := newPuzzle()
	myCell := myPuzzle.getCell(3, 5)
	numOptions := myCell.numValueOptions()
	if (numOptions != 9) {
		t.Error("Puzzle cells should have 9 options, not ", numOptions)
	}
}

func TestPuzzleSetValue(t *testing.T) {
	myPuzzle := newPuzzle()
	myPuzzle.setValue(3, 2, 8)
	myCell := myPuzzle.getCell(3, 2)
	if (myCell.value != 8) {
		t.Error("Puzzle should have 8 set at 3,2.")
	}
	if (myCell.valueKnown != true) {
		t.Error("Puzzle should be set at 3,2.")
	}
}


func TestPuzzleSetValueOptions(t *testing.T) {
	myPuzzle := newPuzzle()
	myPuzzle.setValueOptions(3, 2, 8)
	myCell := myPuzzle.getCell(3, 2)
	if (myCell.valueOptions != 8) {
		t.Error("Puzzle should have valueOption '8' at 3,2.")
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

func TestCalcValueOptions(t *testing.T) {
	myCell := new(Cell)
	myCell.setValue(3, 0)
	cells := []Cell{ *myCell }
	valueOption, _ := calcValueOptions(cells)
	if valueOption != 507 {
		t.Error("Value option should be ", 507, " not ", valueOption)
	}

}

func TestUpdateRow(t *testing.T) {
	myPuzzle := newPuzzle()
	myPuzzle.setValue(1, 2, 8)
	myPuzzle.updateRow(1)
	checkCell := myPuzzle.getCell(1, 5)
	hasOption, _ := checkCell.hasOptions(8)
	if hasOption {
		t.Error("Cell should not have option 8 if 8 already in row: ",
		checkCell.valueOptions)
	}
}

func TestUpdateCol(t *testing.T) {
	myPuzzle := newPuzzle()
	myPuzzle.setValue(1, 2, 8)
	myPuzzle.updateCol(2)
	checkCell := myPuzzle.getCell(5, 2)
	hasOption, _ := checkCell.hasOptions(8)
	if hasOption {
		t.Error("Cell should not have option 8 if 8 already in col: ",
		checkCell.valueOptions)
	}
}

func TestUpdateBlock(t *testing.T) {
	myPuzzle := newPuzzle()
	myPuzzle.setValue(1, 7, 8)
	myPuzzle.updateBlock(1, 7)
	checkCell := myPuzzle.getCell(2, 6)
	hasOption, _ := checkCell.hasOptions(8)
	if hasOption {
		t.Error("Cell should not have option 8 if 8 already in block: ",
		checkCell.valueOptions)
	}
}

func TestUpdateAll(t *testing.T) {
	myPuzzle := puzzleFromFile("almost_empty_test.txt")
	myPuzzle.updateAll()
	var cells [3]Cell
	cells[0] = *myPuzzle.getCell(0, 1)
	cells[1] = *myPuzzle.getCell(2, 6)
	cells[2] = *myPuzzle.getCell(8, 1)

	for i, cell := range cells {
		hasOption, _ := cell.hasOptions(2)
		if hasOption {
			t.Error("Cell ", i, " should not have option 2: ", cell)
		}
	}
}

func TestPuzzleComplete(t *testing.T) {
	myPuzzle := puzzleFromFile("almost_complete_test.txt")
	if isComplete, _ := myPuzzle.isComplete(); isComplete {
		t.Error("Puzzle should not be considered complete: ", myPuzzle)
	}

	myPuzzle.fillOneCell()
	if isComplete, _ := myPuzzle.isComplete(); !isComplete {
		t.Error("Puzzle should be considered complete: ", myPuzzle)
	}
	
}

func TestFillOneCell(t *testing.T) {
	myPuzzle := puzzleFromFile("almost_complete_test.txt")
	myPuzzle.fillOneCell()
	myCell := myPuzzle.getCell(4, 7)
	if myCell.value != 5 {
		t.Error("Cell should have 3 at 4, 7, not: ", myCell.value)
	}
}

func TestFillAllCells(t *testing.T) {
	myPuzzle := puzzleFromFile("row_missing_test.txt")
	err := myPuzzle.fillAllCells()
	if err != nil {
		t.Error(err)
	}
	if isComplete, _ := myPuzzle.isComplete(); !isComplete {
		t.Error("Puzzle should have been completed: ", myPuzzle)
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
	puzzleOne.updateAll()
	puzzleTwo := puzzleFromFile("almost_empty_test.txt")
	if !(reflect.DeepEqual(puzzleOne, puzzleTwo)) {
		t.Error("Puzzles should be equal.", puzzleOne, puzzleTwo)
	}
	
	
}
*/

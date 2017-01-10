package sudoku

import (
	"testing"
)



func TestGridNumCells(t *testing.T) {
	myGrid := newGrid()
	got := len(myGrid.grid)
	expected := 9
		
	if got != expected {
		t.Errorf("Expected %d rows, not %d.",
			expected, got)
	}

	for _, row := range myGrid.grid {
		got = len(row)
		if (got != expected) {
			t.Error("Expected %d columns, not %d.",
				expected, got)
		}
	}
}


func TestGridSetValue(t *testing.T) {
	var testCells = []struct {
		row int
		col int
		value uint
	}{
		{0, 0, 1},
		{1, 0, 4},
		{6, 8, 9},
		{8, 8, 2},
	}
	
	myGrid := newGrid()
	for _, testCell := range testCells {
		row := testCell.row
		col := testCell.col
		value := testCell.value
		myGrid.setCellValue(row, col, value)
		expected := value
		got, _ := myGrid.getCellValue(row, col)
		if got != expected {
			t.Errorf("Expected %d at %d,%d but got %d.",
				expected, row, col, got)
		}
	}
}

func TestSetOptionValues(t *testing.T) {
	var testOptions = []struct {
		rowNum int
		colNum int
		value uint
		expected bitarray
	}{
		{8, 1, 9, allTrue()},
		{0, 5, 3, bitNot(newBitArray(3))},
		{1, 1, 5, bitNot(newBitArray(3, 5))},
		{5, 0, 8, bitNot(newBitArray(3, 5, 8))},

	}

	myGrid := newGrid()
	for _, testOption := range testOptions {
		myGrid.setCellValue(testOption.rowNum, testOption.colNum, testOption.value)
		myGrid.updateOptions(0, 0)
		myCell, _ := myGrid.getCell(0, 0)
		got := myCell.options
		expected := testOption.expected
		if got != expected {
			t.Errorf("Set %d, expected %d but got %d",
				testOption.value, expected, got)
		}
	}
}


/*
func TestPuzzleCellNumOptions(t *testing.T) {
	myGrid := newPuzzle()
	myCell := myGrid.getCell(3, 5)
	numOptions := myCell.numValueOptions()
	if (numOptions != 9) {
		t.Error("Puzzle cells should have 9 options, not ", numOptions)
	}
}

func TestPuzzleSetValue(t *testing.T) {
	myGrid := newPuzzle()
	myGrid.setValue(3, 2, 8)
	myCell := myGrid.getCell(3, 2)
	if (myCell.value != 8) {
		t.Error("Puzzle should have 8 set at 3,2.")
	}
	if (myCell.valueKnown != true) {
		t.Error("Puzzle should be set at 3,2.")
	}
}


func TestPuzzleSetValueOptions(t *testing.T) {
	myGrid := newPuzzle()
	myGrid.setValueOptions(3, 2, 8)
	myCell := myGrid.getCell(3, 2)
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
	myGrid := newPuzzle()
	myGrid.setValue(3, 2, 8)
	value, ok := myGrid.getValue(3, 2)
	if ok != nil {
		t.Error("")
	}

	if value != 8 {
		t.Error()
	}

	value, ok = myGrid.getValue(3, 3)
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
	myGrid := newPuzzle()
	myGrid.setValue(1, 2, 8)
	myGrid.updateRow(1)
	checkCell := myGrid.getCell(1, 5)
	hasOption, _ := checkCell.hasOptions(8)
	if hasOption {
		t.Error("Cell should not have option 8 if 8 already in row: ",
		checkCell.valueOptions)
	}
}

func TestUpdateCol(t *testing.T) {
	myGrid := newPuzzle()
	myGrid.setValue(1, 2, 8)
	myGrid.updateCol(2)
	checkCell := myGrid.getCell(5, 2)
	hasOption, _ := checkCell.hasOptions(8)
	if hasOption {
		t.Error("Cell should not have option 8 if 8 already in col: ",
		checkCell.valueOptions)
	}
}

func TestUpdateBlock(t *testing.T) {
	myGrid := newPuzzle()
	myGrid.setValue(1, 7, 8)
	myGrid.updateBlock(1, 7)
	checkCell := myGrid.getCell(2, 6)
	hasOption, _ := checkCell.hasOptions(8)
	if hasOption {
		t.Error("Cell should not have option 8 if 8 already in block: ",
		checkCell.valueOptions)
	}
}

func TestUpdateAll(t *testing.T) {
	myGrid := puzzleFromFile("almost_empty_test.txt")
	myGrid.updateAll()
	var cells [3]Cell
	cells[0] = *myGrid.getCell(0, 1)
	cells[1] = *myGrid.getCell(2, 6)
	cells[2] = *myGrid.getCell(8, 1)

	for i, cell := range cells {
		hasOption, _ := cell.hasOptions(2)
		if hasOption {
			t.Error("Cell ", i, " should not have option 2: ", cell)
		}
	}
}

func TestPuzzleComplete(t *testing.T) {
	myGrid := puzzleFromFile("almost_complete_test.txt")
	if isComplete, _ := myGrid.isComplete(); isComplete {
		t.Error("Puzzle should not be considered complete: ", myGrid)
	}

	myGrid.fillOneCell()
	if isComplete, _ := myGrid.isComplete(); !isComplete {
		t.Error("Puzzle should be considered complete: ", myGrid)
	}
	
}

func TestFillOneCell(t *testing.T) {
	myGrid := puzzleFromFile("almost_complete_test.txt")
	myGrid.fillOneCell()
	myCell := myGrid.getCell(4, 7)
	if myCell.value != 5 {
		t.Error("Cell should have 3 at 4, 7, not: ", myCell.value)
	}
}

func TestFillAllCells(t *testing.T) {
	myGrid := puzzleFromFile("row_missing_test.txt")
	err := myGrid.fillAllCells()
	if err != nil {
		t.Error(err)
	}
	if isComplete, _ := myGrid.isComplete(); !isComplete {
		t.Error("Puzzle should have been completed: ", myGrid)
	}
}

func TestInsertRow(t *testing.T) {
	myGrid := newPuzzle()
	row := "123456789"
	myGrid.insertRow(1, row)
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

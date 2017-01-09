package sudoku

import (

)

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

*/

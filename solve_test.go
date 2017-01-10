package sudoku

import (
	"testing"
)

func TestEasiestCell(t *testing.T) {
	var testEasiest = []struct {
		rowNum int
		colNum int
		value uint
		easyRow int
		easyCol int
		easyValue uint
	}{
		{0, 1, 1, 0, 0, 2},
		{2, 8, 3, 0, 6, 4},
		{5, 2, 5, 0, 2, 3},
	}
	
	myGrid := newGrid()
	tried := 0
	for _, testEasy := range testEasiest {
		myGrid.setCellValue(testEasy.rowNum, testEasy.colNum, testEasy.value, 0)
		myGrid.reset(tried)
		expectedRow := testEasy.easyRow
		expectedCol := testEasy.easyCol

		t.Run("Get easiest cell", func(t *testing.T) {
			gotRow, gotCol := getEasiestCell(myGrid)
			if (gotRow != expectedRow) || (gotCol != expectedCol) {
				t.Errorf("Expected coordinates %d, %d but got %d, %d in:\n%s",
					expectedRow, expectedCol, gotRow, gotCol, myGrid)
			}
		})
		t.Run("Fill easiest cell", func(t *testing.T) {
			fillCell(myGrid, tried)
			expectedValue := testEasy.easyValue
			gotValue, _ := myGrid.getCellValue(expectedRow, expectedCol)
			if gotValue != expectedValue {
				t.Errorf("Expected %d, %d to be %d but got %d.",
				expectedRow, expectedCol, expectedValue, gotValue)
			}
		})

		tried += 1
	}
	
}

func TestMistakes(t *testing.T) {
	myGrid := newGrid()
	myGrid.setCellValue(0, 0, 9, 1)
	myGrid.setCellValue(0, 1, 8, 2)
	myCell, _ := myGrid.getCell(0, 2)
	for i:= uint(1) ; i < 9; i++ {
		myCell.setTriedValue(i)
	}

	tried, _ := fillCell(myGrid, 2)
	fillCell(myGrid, tried)
	expectedValue := uint(1)
	got, _ := myGrid.getCellValue(0, 0)
	
	if got != expectedValue {
		t.Errorf("Cell %d, %d was expected to be %d but got %d in:\n%s",
			0, 0, expectedValue, got, myGrid)
	}
}

func TestSolve(t *testing.T) {
	myGrid := newGrid()
	err := Solve(myGrid, 0)

	if err != nil {
		t.Error(err)
	}

	if !(myGrid.isComplete() && myGrid.isValid()) {
		t.Errorf("Expected grid to be complete: ", myGrid)
	}

}

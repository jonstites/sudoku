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
		myGrid.updateAllOptions()
		myCell, _ := myGrid.getCell(0, 0)
		got := myCell.options
		expected := testOption.expected
		if got != expected {
			t.Errorf("Set %d, expected %d but got %d",
				testOption.value, expected, got)
		}
	}
}

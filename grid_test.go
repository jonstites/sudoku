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
		myGrid.setCellValue(row, col, value, 0)
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
		myGrid.setCellValue(testOption.rowNum, testOption.colNum, testOption.value, 0)
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

func TestReset(t *testing.T) {
	var testResetValues = []struct {
		row int
		col int
		value uint
		reset int
		options bitarray
	}{
		{0, 8, 0, 5, newBitArray(6, 7, 8, 9)},
		{0, 4, 5, 5, newBitArray(5)},
		{0, 4, 0, 4, newBitArray(5, 6, 7, 8, 9)},
		{0, 3, 4, 4, newBitArray(4)},
		{0, 0, 1, 1, newBitArray(1)},
		{0, 0, 0, 0, allTrue()},
	}

	myGrid := newGrid()
	for i := 0; i < 5; i++ {
		myGrid.setCellValue(0, i, uint(i + 1), i+1)
	}

	for _, testReset := range testResetValues {
		row := testReset.row
		col := testReset.col
		myGrid.reset(testReset.reset)
		t.Run("Test reset values: ", func(t *testing.T) {
			expected := testReset.value
			got, _ := myGrid.getCellValue(row, col)
			if expected != got {
				t.Errorf("Cell %d, %d was expected to have %d but got %d in:\n%s",
					row, col, expected, got, myGrid)
			}
		})

		t.Run("Test reset options: ", func(t *testing.T) {
			expected := testReset.options
			myCell, _ := myGrid.getCell(row, col)
			got := myCell.options
			if got != expected {
				t.Errorf("Cell %d, %d was expected to have %d but got %d in:\n%s",
					row, col, expected, got, myGrid)
			}
		})
	}
}

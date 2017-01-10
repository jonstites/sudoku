package sudoku


import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Read a grid from a file.
func Read(filename string) *grid {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	myGrid, err := fileToGrid(file)
	
	return myGrid
}

// Scan the file
func fileToGrid(file *os.File) (*grid, error) {
	myGrid := newGrid()
	
	scanner := bufio.NewScanner(file)

	cellsFilled := 0
	for scanner.Scan() {
		row := scanner.Text()

		num := len(row)
		if num != 9 && num != 81 {
			
		}
		
		if err := validateRow(row); err != nil {
			log.Fatal(err)
		}

		insertRow(myGrid, cellsFilled, row)
		cellsFilled += len(row)
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return myGrid, nil
}

// Insert the characters into the puzzle
func insertRow(myGrid *grid, cellsFilled int, row string) {
	for colNum, char := range row {

		// Throw error if too large
		cellsFilled += 1
		if cellsFilled > 81 {
			log.Fatal("Puzzle has more than 81 squares.")
		}

		// Rune to int
		value := uint(char - '0')

		// Identify row number using total number of cells filled
		rowNum := cellsFilled / 9

		// Put into the grid
		if value != 0 {
			myGrid.setCellValue(rowNum, colNum, value, 0)
		} else {
			cell, _ := myGrid.getCell(rowNum, colNum)
			cell.reset(0)
		}
	}
}

// Check to make sure string contains numeric characters only
func isNumeric(row string) bool {
	for _, char := range row {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// Raise error if row does not conform
func validateRow(row string) error {
	if !(isNumeric(row)) {
		return fmt.Errorf("File contains non-numeric characters in row.")
	}

	if len(row) != 9 && len(row) != 81 {
		return fmt.Errorf("Row must contain 9 or 81 digits.")
	}

	return nil
}


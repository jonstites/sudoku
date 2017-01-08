package sudoku

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type valueSet int
const fullSet = 511

// Initialize a valueSet with given values
func newValueSet(values ...uint) valueSet {
	var myValueSet valueSet
	for _, value := range values {
		myValueSet += 1 << (value - 1)
	}
	return myValueSet
}

// Represents one square of a Sudoku puzzle.
type Cell struct {
	// The determined value (meaningless if valueKnown is false)
	value int
	// Stores which values the cell is allowed to be (meaningless if valueKnown is true)
	valueOptions valueSet
	// Whether the value is known
	valueKnown bool
}

func newCell() *Cell {
	cell := new(Cell)
	cell.valueOptions = fullSet
	return cell
}

func (c *Cell) hasOptions(options ...uint) (bool, error) {
	cellOptions := c.valueOptions
	for _, option := range options {
		if option < 1 || option > 9 {
			return false, fmt.Errorf("Not a valid sudoku value: ", option)
		}

		// check if bit is unset
		if (cellOptions >> (option - 1) & 1) != 1 {
			return false, nil
		}
	}		
	return true, nil
}
		
// Returns the number of values the cell is allowed to be in
func (c *Cell) numValueOptions() int {
	num := 0
	for valueOptions := c.valueOptions; valueOptions != 0; valueOptions = valueOptions >> 1 {
		if (valueOptions & 1) == 1 {
			num += 1
		}
	}
	return num
}

func (c *Cell) setValue(value int) {
	c.value = value
	c.valueKnown = true
}

func (c *Cell) chooseValue() int {
	
	i := 1
	for valueOptions := c.valueOptions; valueOptions != 0; valueOptions = valueOptions >> 1 {
		if (valueOptions & 1) == 1 {
			return i
		}
		i++
	}
	return i
}

// Just print the value of a Cell f
func (c *Cell) String() string {
	valueString := strconv.Itoa(c.value)
	if !(c.valueKnown) {
		valueString = "0"
	}
	return valueString
}

type Puzzle struct {
	puzzle [][]Cell
}

func newPuzzle() *Puzzle {
	myPuzzle := new(Puzzle)
	myPuzzle.puzzle = make([][]Cell, 9)
	for i := 0; i < 9; i++ {
		myPuzzle.puzzle[i] = make([]Cell, 9)
		for j := 0; j < 9; j++ {
			myPuzzle.puzzle[i][j] = *newCell()
		}
	}
	return myPuzzle
}

func (myPuzzle *Puzzle) getCell(rowNum int, colNum int) *Cell {
	return &myPuzzle.puzzle[rowNum][colNum]
}

func (myPuzzle *Puzzle) setValue(rowNum int, colNum int, value int) {
	myCell := myPuzzle.getCell(rowNum, colNum)
	myCell.setValue(value)
}

func (myPuzzle *Puzzle) setValueOptions(rowNum int, colNum int, valueOptions valueSet) {
	myCell := myPuzzle.getCell(rowNum, colNum)
	myCell.valueOptions = valueOptions & myCell.valueOptions
}

func (myPuzzle *Puzzle) getValue(rowNum int, colNum int) (int, error) {
	myCell := myPuzzle.getCell(rowNum, colNum)
	if !(myCell.valueKnown) {
		return -1, fmt.Errorf("cell not set.")
	}
	return myCell.value, nil
}


func calcValueOptions(cells []Cell) (valueSet, error) {
	var valueOptions valueSet
	valueOptions = fullSet
	for _, cell := range cells {
		if cell.valueKnown {
			valueOptions -= 1 << (uint(cell.value) - 1)
		}
	}

	if (valueOptions < 0) || (valueOptions > fullSet) {
		return valueOptions, fmt.Errorf("Row options out of range: ", valueOptions)
	}

	return valueOptions, nil
}


func (myPuzzle *Puzzle) updateRow(rowNum int) error {
	cells := make([]Cell, len(myPuzzle.puzzle))
	for colNum, cell := range myPuzzle.puzzle[rowNum] {
		cells[colNum] = cell
	}

	valueOptions, err := calcValueOptions(cells)

	for colNum, _ := range myPuzzle.puzzle[rowNum] {
		myPuzzle.setValueOptions(rowNum, colNum, valueOptions)
	}
	
	return err
}

func (myPuzzle *Puzzle) updateCol(colNum int) error {
	cells := make([]Cell, len(myPuzzle.puzzle))
	for rowNum, row := range myPuzzle.puzzle {
		cells[rowNum] = row[colNum]
	}

	valueOptions, err := calcValueOptions(cells)
	for rowNum, _ := range myPuzzle.puzzle {
		myPuzzle.setValueOptions(rowNum, colNum, valueOptions)
	}

	return err
}


func (myPuzzle *Puzzle) updateBlock(rowNum int, colNum int) error {
	cells := make([]Cell, len(myPuzzle.puzzle))
	blockRow := rowNum - (rowNum % 3)
	blockCol := colNum - (colNum % 3)
	i := 0
	for rowIndex := blockRow; rowIndex < blockRow + 3; rowIndex += 1 {
		for colIndex := blockCol; colIndex < blockCol + 3; colIndex += 1 {
			cells[i] = *myPuzzle.getCell(rowIndex, colIndex)
			i++
		}
	}

	valueOptions, err := calcValueOptions(cells)

	for rowIndex := blockRow; rowIndex < blockRow + 3; rowIndex += 1 {
		for colIndex := blockCol; colIndex < blockCol + 3; colIndex += 1 {
			myPuzzle.setValueOptions(rowIndex, colIndex, valueOptions)
		}
	}
	return err
		
}

func (myPuzzle *Puzzle) updateAll() error {
	for i := 0; i < 9; i++ {
		err := myPuzzle.updateRow(i)
		if err != nil {
			return err
		}
		err = myPuzzle.updateCol(i)
		if err != nil {
			return err
		}
	}

	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			err := myPuzzle.updateBlock(i, j)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (myPuzzle *Puzzle) fillOneCell() error {
	bestRowNum := 0
	bestColNum := 0
	bestNumOptions := 10
	bestCellValue := 0
	isComplete := true

	for rowNum, row := range myPuzzle.puzzle {
		for colNum, cell := range row {
			if cell.valueKnown {
				continue
			}

			isComplete = false
			numOptions := cell.numValueOptions()

			if (numOptions >= bestNumOptions) {
				continue
			}

			bestNumOptions = numOptions
			bestRowNum = rowNum
			bestColNum = colNum
			bestCellValue = cell.chooseValue()
		}
	}

	if isComplete {
		return fmt.Errorf("Can't fill a cell: sudoku complete.")
	}

	// if stack is empty... (TODO)
	if bestNumOptions == 0 {
		return fmt.Errorf("This sudoku puzzle cannot be solved.")
	}

	if bestNumOptions > 1 {
		return fmt.Errorf("Guessing not implemented.")
	}

	myPuzzle.setValue(bestRowNum, bestColNum, bestCellValue)
	myPuzzle.updateRow(bestRowNum)
	myPuzzle.updateCol(bestColNum)
	myPuzzle.updateBlock(bestRowNum, bestColNum)
	return nil
}

func (myPuzzle *Puzzle) fillAllCells() error {
	for !(myPuzzle.isComplete()) {
		err := myPuzzle.fillOneCell() 
		if err != nil {
			return err
		}
	}
	return nil
}

func (myPuzzle *Puzzle) isComplete() bool {
	for _, row := range myPuzzle.puzzle {
		for _, cell := range row {
			if !cell.valueKnown {
				return false
			}
		}
	}
	return true
}

func (myPuzzle *Puzzle) insertRow(rowNum int, row string) {
	for colNum, char := range row {
		value := int(char-'0')
		if value != 0 {
			myPuzzle.setValue(rowNum, colNum, value)
		}
	}
}

func box(cells [][]Cell) string {
	var box []string
	for _, row := range cells {
		var rowValues []string
		for _, col := range row {
			rowValues = append(rowValues, col.String())
		}
		box = append(box, strings.Join(rowValues, ""))
	}
	return strings.Join(box, "\n")
}

func (myPuzzle *Puzzle) String() string {
	return box(myPuzzle.puzzle)
}

func isNumeric(row string) bool {
	for _, char := range row {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func validateRowFormat(row string) error {
	if !(isNumeric(row)) {
		return fmt.Errorf("File contains non-numeric characters in row: ", row)
	}

	if len(row) != 9 {
		return fmt.Errorf("Row does not contain 9 digits: ", row) 
	}

	return nil
}

func puzzleFromFile(filename string) *Puzzle {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	myPuzzle := newPuzzle()
	
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		if i >= 9 {
			log.Fatal("Files contains more than 9 rows.")
		}
		err := validateRowFormat(row)
		if err != nil {
			log.Fatal(err)
		}

		myPuzzle.insertRow(i, row)
		i++
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	myPuzzle.updateAll()
	return myPuzzle
}

func main() {
	fmt.Printf("hello, world\n")
}

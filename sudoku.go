package sudoku

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

type valueSet int
const fullSet = 511

// Initialize a valueSet with 1-9
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

		// check if bit is set
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

func (myPuzzle *Puzzle) setValue(rowNum int, colNum int, value int) {
	myPuzzle.puzzle[rowNum][colNum].value = value
	myPuzzle.puzzle[rowNum][colNum].valueKnown = true
}

func (myPuzzle *Puzzle) getValue(rowNum int, colNum int) (int, error) {
	if !(myPuzzle.puzzle[rowNum][colNum].valueKnown) {
		return -1, fmt.Errorf("cell not set.")
	}

	return myPuzzle.puzzle[rowNum][colNum].value, nil
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
	return myPuzzle
}

func main() {
	fmt.Printf("hello, world\n")
}

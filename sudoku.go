package sudoku
/*
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


type Puzzle struct {
	puzzle [][]Cell
	guesses [][]int
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
	myCell.setValue(value, len(myPuzzle.guesses))
}

func (myPuzzle *Puzzle) setValueOptions(rowNum int, colNum int, valueOptions valueSet) {
	myCell := myPuzzle.getCell(rowNum, colNum)
	if rowNum == 1 && colNum == 0 {
		fmt.Println("1, 0: ", valueOptions, myCell.guesses, myCell.valueOptions,
			(valueOptions &^ myCell.guesses) & myCell.valueOptions)
	}

	myCell.valueOptions = (valueOptions &^ myCell.guesses) & myCell.valueOptions
	
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
	fmt.Println("Updating all")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			myCell := myPuzzle.getCell(i, j)
			if myCell.guessNum > len(myPuzzle.guesses) {
				myPuzzle.puzzle[i][j].valueKnown = false
				myPuzzle.puzzle[i][j].valueOptions = fullSet
			}
		}
	}
	
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


	if bestNumOptions == 0 {
		guessSize := len(myPuzzle.guesses)
		fmt.Println("Problem cell is: ", bestRowNum, bestColNum)
		if guessSize == 0 {
			return fmt.Errorf("This sudoku puzzle cannot be solved.",
			myPuzzle.guesses, "\n", myPuzzle)
		}

		fmt.Println("Removed a guess at: ", myPuzzle.guesses[guessSize - 1])
		fmt.Println("Was:\n", myPuzzle)
		myPuzzle.guesses = myPuzzle.guesses[:guessSize - 1]
		myPuzzle.updateAll()
		fmt.Println("Is:\n", myPuzzle)
	}
	

	if bestNumOptions > 1 {
		coordinates := []int {bestRowNum, bestColNum}


		myPuzzle.guesses = append(myPuzzle.guesses, coordinates)
		myPuzzle.addGuess(bestRowNum, bestColNum, bestCellValue)
		fmt.Println("Made a guess at: ", bestRowNum, bestColNum, bestCellValue)
		fmt.Println("Was:\n", myPuzzle)
	}

	myPuzzle.setValue(bestRowNum, bestColNum, bestCellValue)
	myPuzzle.updateRow(bestRowNum)
	myPuzzle.updateCol(bestColNum)
	myPuzzle.updateBlock(bestRowNum, bestColNum)
	if bestNumOptions > 1 {
		fmt.Println("Is:\n", myPuzzle)
	}
	return nil
}

func (myPuzzle *Puzzle) addGuess(rowNum int, colNum int, value int) {
	myCell := myPuzzle.getCell(rowNum, colNum)
	myCell.addGuess(value)
}

func (myPuzzle *Puzzle) fillAllCells() error {
	isComplete, err := myPuzzle.isComplete()
	if err != nil {
		return err
	}
	for !(isComplete) {
		err := myPuzzle.fillOneCell() 
		if err != nil {
			return err
		}
		isComplete, err = myPuzzle.isComplete()
		if err != nil {
			return err
		}
	}
	isComplete, _ = myPuzzle.isComplete()
	fmt.Println(isComplete)
	return nil
}

func (myPuzzle *Puzzle) isComplete() (bool, error) {
	for _, row := range myPuzzle.puzzle {
		for _, cell := range row {
			if !cell.valueKnown {
				return false, nil
			}
			if cell.value == 0 {
				return false, fmt.Errorf("Error: Sudoku solution contains a zero.")
			}
		}
	}
	return true, nil
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
	sudokuFile := flag.String("filename", "", "File of sudoku puzzle to solve.")
	flag.Parse()
	myPuzzle := puzzleFromFile(*sudokuFile)
	err := myPuzzle.fillAllCells()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myPuzzle)
}
*/

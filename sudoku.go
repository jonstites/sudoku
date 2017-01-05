package sudoku

import (
	"fmt"
	"strconv"
)

// Represents one square of a Sudoku puzzle.
type Cell struct {
	// The determined value (meaningless if valueSet is false)
	value int
	// Stores which values the cell is allowed to be
	valueOptions map[int]bool
	// Whether the value has been set
	valueSet bool
}

// Returns the number of values the cell is allowed to be in
func (c *Cell) numValueOptions() int {
	num := 0
	for _, value := range c.valueOptions {
		if value {
			num += 1
		}
	}
	return num
}

// Initialize a cell with possible values 1-9
func newCell() *Cell {
	c := Cell{}
	valueMap := map[int]bool{}
	for i := 1; i <= 9; i++ {
		valueMap[i] = true
	}
	c.valueOptions = valueMap
	return &c
}


// Just print the value of a Cell 
func (c *Cell) String() string {
	return strconv.Itoa(c.value)
}

func main() {
	fmt.Printf("hello, world\n")
}

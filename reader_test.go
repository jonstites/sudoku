package sudoku


import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidateRowFormat(t *testing.T) {
	var testRows = []struct {
		row string
		expectedErr error
	}{
		{"123456788", nil},
		{"12345", fmt.Errorf("Row must contain 9 or 81 digits.")},
	}

	for _, testRow := range testRows {
		got := validateRow(testRow.row)
		expected := testRow.expectedErr
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %q but got %q",
				expected, got)
		}
	}
}


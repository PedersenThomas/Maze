package test

import (
	"testing"

	. "github.com/pedersenthomas/Maze/model"
)

func TestInit(t *testing.T) {
	b := &Board{Height: 10, Width: 10, Cells: nil}
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell.Flag != 15 {
				t.Errorf("Init() should init all Flag to %d, got %d for %+v", ALLDIRECTIONS, cell.Flag, cell)
			}
		}
	}
}

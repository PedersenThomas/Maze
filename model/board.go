package model

import (
	"io"
	"math"
)

type Board struct {
	Height uint16 // Row
	Width  uint16 // Column
	Cells  [][]Cell
}

func (b *Board) Init() {
	b.Cells = make([][]Cell, b.Height)
	for i := uint16(0); i < b.Height; i++ {
		b.Cells[i] = make([]Cell, b.Width)
	}

	for h := uint16(0); h < b.Height; h++ {
		for w := uint16(0); w < b.Width; w++ {
			// set the flag field with 15, which in binary will be 0b00001111,
			// the 4 bits indicates that all 4 walls are up, so the cells are
			// isolated/sealed from each other initially. ex:
			//  _ _
			// |_|_|
			// |_|_|
			//
			b.Cells[h][w].Flag = ALLDIRECTIONS

			// set the relative [x,y] position of the cell on the board
			b.Cells[h][w].X = h
			b.Cells[h][w].Y = w
		}
	}
}

func (b *Board) Cell(x, y uint16) *Cell {
	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width {
		return &b.Cells[x][y]
	}

	return nil
}

func (b *Board) Neighbours(cell *Cell) []*Cell {
	result := make([]*Cell, 0, 4)

	if cell.X != 0 {
		if neighbour := b.Cell(cell.X-1, cell.Y); neighbour != nil {
			result = append(result, neighbour)
		}
	}

	if cell.Y != 0 {
		if neighbour := b.Cell(cell.X, cell.Y-1); neighbour != nil {
			result = append(result, neighbour)
		}
	}

	if cell.X != math.MaxUint16 {
		if neighbour := b.Cell(cell.X+1, cell.Y); neighbour != nil {
			result = append(result, neighbour)
		}
	}

	if cell.Y != math.MaxUint16 {
		if neighbour := b.Cell(cell.X, cell.Y+1); neighbour != nil {
			result = append(result, neighbour)
		}
	}

	return result
}

func (b *Board) BreakWall(from, to *Cell) {
	direction := b.FindDirection(from, to)

	switch direction {
	case EAST:
		from.ClearBit(EAST)
		to.ClearBit(WEST)
	case SOUTH:
		from.ClearBit(SOUTH)
		to.ClearBit(NORTH)
	case WEST:
		from.ClearBit(WEST)
		to.ClearBit(EAST)
	case NORTH:
		from.ClearBit(NORTH)
		to.ClearBit(SOUTH)
	}
}

func (b *Board) FindDirection(from, to *Cell) FlagPosition {
	// X denotes row, Y denotes col
	//
	//        col 0  | col 1 | col 2
	// --------------------------------
	// row 0  [x0,y0] [x0,y1] [x0,y2]
	// row 1  [x1,y0] [x1,y1] [x1,y2]
	if from.X < to.X {
		return SOUTH
	}
	if from.X > to.X {
		return NORTH
	}
	if from.Y < to.Y {
		return EAST
	}
	if from.Y > to.Y {
		return WEST
	}
	//TODO: This is really an error case here
	return EAST
}

func (b *Board) Write(writer io.Writer) {
	writer.Write([]byte("  "))
	for i := uint16(1); i < b.Width; i++ {
		writer.Write([]byte(" _"))
	}
	writer.Write([]byte("\n"))

	for h := uint16(0); h < b.Height; h++ {
		writer.Write([]byte("|"))
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]
			if w == b.Width-1 && h == b.Height-1 {
				writer.Write([]byte(" |"))
				break
			}
			if c.IsSet(SOUTH) {
				writer.Write([]byte("_"))
			} else {
				writer.Write([]byte(" "))
			}

			if c.IsSet(EAST) {
				writer.Write([]byte("|"))
			} else {
				writer.Write([]byte(" "))
			}
		}
		writer.Write([]byte("\n"))
	}
}

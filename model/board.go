package model

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

package algo

import (
	"math/rand"

	"github.com/pedersenthomas/Maze/model"
	"github.com/pedersenthomas/Maze/utils"
)

type DepthFirstMazeGenerator struct {
	Board model.Board
	Rng   *rand.Rand
}

func NewDepthFirstMazeGenerator(height, width uint16) *DepthFirstMazeGenerator {
	p := &DepthFirstMazeGenerator{Board: model.Board{Height: height, Width: width, Cells: nil}}
	p.Board.Init()
	return p
}

func (maze *DepthFirstMazeGenerator) Generate() error {
	opStack := utils.Stack{}

	// Pick Starting location
	start := maze.Board.Cell(0, 0)
	start.SetBit(model.VISITED)
	start.SetBit(model.START)
	end := maze.Board.Cell(maze.Board.Height-1, maze.Board.Width-1)
	end.SetBit(model.END)
	opStack.Push(start)

	for !opStack.IsEmpty() {
		currentCell := opStack.Peek().(*model.Cell)

		neighbours := getUnvisitedNeighbours(&maze.Board, currentCell)
		if len(neighbours) == 0 {
			opStack.Pop()
			continue
		}

		idx := utils.RandomIntN(maze.Rng, len(neighbours))
		nextCell := neighbours[idx]
		nextCell.SetBit(model.VISITED)
		maze.Board.BreakWall(currentCell, nextCell)
		if !nextCell.IsSet(model.END) {
			opStack.Push(nextCell)
		}
	}

	return nil
}

func getUnvisitedNeighbours(b *model.Board, Cell *model.Cell) []*model.Cell {
	result := make([]*model.Cell, 0, 4)
	for _, neightbour := range b.Neighbours(Cell) {
		if !neightbour.IsSet(model.VISITED) {
			result = append(result, neightbour)
		}
	}

	return result
}

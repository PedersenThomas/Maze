package test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/pedersenthomas/Maze/algo"
	"github.com/pedersenthomas/Maze/model"
)

func TestPrim(t *testing.T) {
	t.Skip("skipping this test for now")
	rand.Seed(time.Now().UTC().UnixNano())
	count := rand.Intn(50) + 1
	for c := 0; c < count; c++ {
		k := algo.NewDepthFirstMazeGenerator(uint16(rand.Intn(100)+1), uint16(rand.Intn(100)+1))
		k.Generate()
		for h := uint16(0); h < k.Board.Height; h++ {
			for w := uint16(0); w < k.Board.Width; w++ {
				if !k.Board.Cells[h][w].IsSet(model.VISITED) {
					t.Errorf("Prim, every cell should be visited, but not [%d,%d]", h, w)
				}
			}
		}
	}
}

func TestNewPrim(t *testing.T) {
	k := algo.NewDepthFirstMazeGenerator(10, 10)
	for _, row := range k.Board.Cells {
		for _, cell := range row {
			if cell.Flag != 15 {
				t.Errorf("NewPrim(), every celll should have flag set to 15, got %d", cell.Flag)
			}
		}
	}
}

func BenchmarkPrimAlgo_1000x500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := algo.NewDepthFirstMazeGenerator(1000, 500)
		k.Generate()
	}
}

func BenchmarkPrimAlgo_100x50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := algo.NewDepthFirstMazeGenerator(100, 50)
		k.Generate()
	}
}

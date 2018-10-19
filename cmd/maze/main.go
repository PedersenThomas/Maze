package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/pedersenthomas/Maze/algo"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	k := algo.NewDepthFirstMazeGenerator(10, 15)
	k.Generate()
	k.Board.Write(os.Stdout)
}

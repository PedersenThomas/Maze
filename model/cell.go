package model

import "fmt"

type FlagPosition uint16

const (
	NORTH FlagPosition = 1 << iota
	SOUTH
	EAST
	WEST
	VISITED
	START
	END
	DEAD  //a dead cell is not part of the solution path
	CROSS //for weave maze
)

const (
	ALLDIRECTIONS FlagPosition = NORTH | SOUTH | EAST | WEST
	ALL                        = NORTH | SOUTH | EAST | WEST | VISITED | START | END | DEAD | CROSS
)

type Cell struct {
	Flag FlagPosition
	X    uint16
	Y    uint16
}

func (c *Cell) SetBit(pos FlagPosition) {
	c.Flag |= pos
}

func (c *Cell) ClearBit(pos FlagPosition) {
	c.Flag &= ^pos
}

func (c *Cell) IsSet(pos FlagPosition) bool {
	return c.Flag&(pos) != 0
}

func (c *Cell) String() string {
	return fmt.Sprintf("[%d,%d]", c.X, c.Y)
}

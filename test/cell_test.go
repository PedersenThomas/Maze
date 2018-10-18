package test

import (
	"testing"

	. "github.com/pedersenthomas/Maze/model"
)

func TestSetBit(t *testing.T) {
	var tests = []struct {
		inputs []FlagPosition
		want   FlagPosition
	}{
		{[]FlagPosition{NORTH}, 1},
		{[]FlagPosition{SOUTH}, 2},
		{[]FlagPosition{EAST}, 4},
		{[]FlagPosition{WEST}, 8},
		{[]FlagPosition{VISITED}, 16},
		{[]FlagPosition{START}, 32},
		{[]FlagPosition{END}, 64},
		{[]FlagPosition{DEAD}, 128},
		{[]FlagPosition{NORTH, SOUTH}, 3},
		{[]FlagPosition{NORTH, SOUTH, EAST, WEST, VISITED, START, END, DEAD}, 255},
	}

	for _, test := range tests {
		c := Cell{}
		for _, input := range test.inputs {
			c.SetBit(input)
		}
		if c.Flag != test.want {
			t.Errorf("SetBit(%q), Flag = %v, want %v", test.inputs, c.Flag, test.want)
		}
	}
}

func TestClearBit(t *testing.T) {
	var tests = []struct {
		inputs []FlagPosition
		want   FlagPosition
	}{
		{[]FlagPosition{NORTH}, 510},
		{[]FlagPosition{SOUTH}, 509},
		{[]FlagPosition{EAST}, 507},
		{[]FlagPosition{WEST}, 503},
		{[]FlagPosition{VISITED}, 495},
		{[]FlagPosition{START}, 479},
		{[]FlagPosition{END}, 447},
		{[]FlagPosition{DEAD}, 383},
		{[]FlagPosition{NORTH, SOUTH}, 508},
		{[]FlagPosition{NORTH, SOUTH, EAST, WEST, VISITED, START, END, DEAD}, 256},
	}

	for _, test := range tests {
		c := Cell{Flag: ALL}
		for _, input := range test.inputs {
			c.ClearBit(input)
		}
		if c.Flag != test.want {
			t.Errorf("ClearBit(%b), Flag = %v, want %v", test.inputs, c.Flag, test.want)
		}
	}
}

func TestIsSet(t *testing.T) {
	c := Cell{Flag: ALL}
	if !c.IsSet(EAST) {
		t.Errorf("IsSet(%q), Flag = %v, want %v", "EAST", c.Flag, true)
	}
}

func BenchmarkSetBit1(b *testing.B) {
	c := Cell{}
	for i := 0; i < b.N; i++ {
		c.SetBit(EAST)
	}
}

func BenchmarkSetBit8(b *testing.B) {
	c := Cell{}
	inputs := []FlagPosition{NORTH, SOUTH, EAST, WEST, VISITED, START, END, DEAD}
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			c.SetBit(input)
		}
	}
}

func BenchmarkClearBit1(b *testing.B) {
	c := Cell{Flag: ALL}
	for i := 0; i < b.N; i++ {
		c.ClearBit(EAST)
	}
}

func BenchmarkClearBit8(b *testing.B) {
	c := Cell{Flag: ALL}
	inputs := []FlagPosition{NORTH, SOUTH, EAST, WEST, VISITED, START, END, DEAD}
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			c.ClearBit(input)
		}
	}
}

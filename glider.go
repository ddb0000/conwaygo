package main

import (
	"fmt"
	"time"
)

const (
	width  = 10
	height = 10
)

// Grid type
type Grid [height][width]bool

// Display the grid
func (g *Grid) Display() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if g[y][x] {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Seed some initial life forms with a more complex pattern
func (g *Grid) Seed() {
	// Glider
	g[1][2] = true
	g[2][3] = true
	g[3][1] = true
	g[3][2] = true
	g[3][3] = true
}

// Next generation
func (g *Grid) Next() Grid {
	var next Grid
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			alive := g[y][x]
			neighbors := g.countNeighbors(x, y)
			if alive && (neighbors == 2 || neighbors == 3) {
				next[y][x] = true
			} else if !alive && neighbors == 3 {
				next[y][x] = true
			}
		}
	}
	return next
}

// Count neighbors
func (g *Grid) countNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			if nx >= 0 && nx < width && ny >= 0 && ny < height && g[ny][nx] {
				count++
			}
		}
	}
	return count
}

func main() {
	var g Grid
	g.Seed()
	for {
		g.Display()
		g = g.Next()
		time.Sleep(time.Second)
	}
}

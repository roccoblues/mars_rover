package main

import "fmt"

type plateau struct {
	x     int
	y     int
	cells [][]*rover
}

func NewPlateau(x, y int) *plateau {
	cells := make([][]*rover, x+1)
	for i := range cells {
		cells[i] = make([]*rover, y+1)
	}

	return &plateau{x, y, cells}
}

func (p *plateau) Put(x int, y int, r *rover) error {
	if !p.insidePlateau(x, y) {
		return fmt.Errorf("Position %d %d is outside of plateau.", x, y)
	}
	if !p.cellEmpty(x, y) {
		return fmt.Errorf("Cell %d %d is occupied.", x, y)
	}

	p.cells[x][y] = r

	return nil
}

func (p *plateau) Get(x int, y int) *rover {
	if !p.insidePlateau(x, y) {
		return nil
	}

	return p.cells[x][y]
}

func (p *plateau) Update(oldX, oldY, newX, newY int) error {
	r := p.Get(oldX, oldY)
	if r == nil {
		return fmt.Errorf("No rover on position %d %d.", oldX, oldY)
	}

	p.cells[oldX][oldY] = nil
	err := p.Put(newX, newY, r)
	if err != nil {
		// If the rover can't move, it should stay in place and not disappear.
		p.cells[oldX][oldY] = r
		return err
	}

	return nil
}

func (p *plateau) cellEmpty(x int, y int) bool {
	if p.cells[x][y] == nil {
		return true
	}
	return false
}

func (p *plateau) insidePlateau(x int, y int) bool {
	if x > p.x || y > p.y {
		return false
	}
	return true
}

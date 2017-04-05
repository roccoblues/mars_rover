package main

import (
	"fmt"
	"strings"
)

type plateau struct {
	x     uint
	y     uint
	cells [][]*rover
}

func NewPlateau(s string) (*plateau, error) {
	x, y, err := convertCoordinates(strings.Fields(s))
	if err != nil {
		return nil, err
	}

	cells := make([][]*rover, x+1)
	for i := range cells {
		cells[i] = make([]*rover, y+1)
	}

	return &plateau{x, y, cells}, nil
}

func (p *plateau) Put(x uint, y uint, r *rover) error {
	if !p.insidePlateau(x, y) {
		return fmt.Errorf("Position out of plateau: %d %d", x, y)
	}
	if !p.cellEmpty(x, y) {
		return fmt.Errorf("Cell occupied: %d %d", x, y)
	}

	p.cells[x][y] = r

	return nil
}

func (p *plateau) Update(oldX, oldY, newX, newY uint, r *rover) error {
	p.cells[oldX][oldY] = nil
	return p.Put(newX, newY, r)
}

func (p *plateau) cellEmpty(x uint, y uint) bool {
	if p.cells[x][y] == nil {
		return true
	}
	return false
}

func (p *plateau) insidePlateau(x uint, y uint) bool {
	if x > p.x || y > p.y {
		return false
	}
	return true
}

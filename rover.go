package main

import (
	"strings"
)

type rover struct {
	X         uint
	Y         uint
	Direction rune
	Commands  []rune
	Plateau   *plateau
}

func NewRover(pos, commands string) (*rover, error) {
	posFields := strings.Fields(pos)
	x, y, err := convertCoordinates(posFields[0:2])
	if err != nil {
		return nil, err
	}

	d, err := convertDirection(posFields[2])
	if err != nil {
		return nil, err
	}

	c, err := convertCommands(commands)
	if err != nil {
		return nil, err
	}

	return &rover{x, y, d, c, nil}, nil
}

func (r *rover) Deploy(p *plateau) error {
	if err := p.Put(r.X, r.Y, r); err != nil {
		return err
	}

	r.Plateau = p

	return nil
}

func (r *rover) Run() error {
	for _, c := range r.Commands {
		err := r.applyCommand(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rover) applyCommand(c rune) error {
	switch c {
	case 'R':
		switch r.Direction {
		case 'N':
			r.Direction = 'E'
		case 'E':
			r.Direction = 'S'
		case 'S':
			r.Direction = 'W'
		case 'W':
			r.Direction = 'N'
		}
	case 'L':
		switch r.Direction {
		case 'N':
			r.Direction = 'W'
		case 'E':
			r.Direction = 'N'
		case 'S':
			r.Direction = 'E'
		case 'W':
			r.Direction = 'S'
		}
	case 'M':
		newX := r.X
		newY := r.Y
		switch r.Direction {
		case 'N':
			newY = r.Y + 1
		case 'E':
			newX = r.X + 1
		case 'S':
			newY = r.Y - 1
		case 'W':
			newX = r.X - 1
		}

		err := r.Plateau.Update(r.X, r.Y, newX, newY, r)
		if err != nil {
			return err
		}

		r.X = newX
		r.Y = newY
	}

	return nil
}

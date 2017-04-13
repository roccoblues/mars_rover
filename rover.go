package main

type rover struct {
	x int
	y int
	d rune
	p *plateau
}

func NewRover(x, y int, d rune) *rover {
	return &rover{x, y, d, nil}
}

func (r *rover) Deploy(p *plateau) error {
	err := p.Put(r.x, r.y, r)
	if err != nil {
		return err
	}

	r.p = p
	return nil
}

func (r *rover) ApplyCommand(c rune) error {
	switch c {
	case 'R':
		switch r.d {
		case 'N':
			r.d = 'E'
		case 'E':
			r.d = 'S'
		case 'S':
			r.d = 'W'
		case 'W':
			r.d = 'N'
		}
	case 'L':
		switch r.d {
		case 'N':
			r.d = 'W'
		case 'E':
			r.d = 'N'
		case 'S':
			r.d = 'E'
		case 'W':
			r.d = 'S'
		}
	case 'M':
		newX := r.x
		newY := r.y
		switch r.d {
		case 'N':
			newY = r.y + 1
		case 'E':
			newX = r.x + 1
		case 'S':
			newY = r.y - 1
		case 'W':
			newX = r.x - 1
		}

		err := r.p.Update(r.x, r.y, newX, newY)
		if err != nil {
			return err
		}

		r.x = newX
		r.y = newY
	}

	return nil
}

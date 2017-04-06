package main

type rover struct {
	x         uint
	y         uint
	direction rune
	commands  []rune
	plateau   *plateau
}

func newRover(p, c string) (*rover, error) {
	x, y, direction, err := convertPosition(p)
	if err != nil {
		return nil, err
	}

	commands, err := convertCommands(c)
	if err != nil {
		return nil, err
	}

	return &rover{x, y, direction, commands, nil}, nil
}

func (r *rover) deploy(p *plateau) error {
	err := p.put(r.x, r.y, r)
	if err != nil {
		return err
	}

	r.plateau = p

	return nil
}

func (r *rover) run() error {
	for _, c := range r.commands {
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
		switch r.direction {
		case 'N':
			r.direction = 'E'
		case 'E':
			r.direction = 'S'
		case 'S':
			r.direction = 'W'
		case 'W':
			r.direction = 'N'
		}
	case 'L':
		switch r.direction {
		case 'N':
			r.direction = 'W'
		case 'E':
			r.direction = 'N'
		case 'S':
			r.direction = 'E'
		case 'W':
			r.direction = 'S'
		}
	case 'M':
		newX := r.x
		newY := r.y
		switch r.direction {
		case 'N':
			newY = r.y + 1
		case 'E':
			newX = r.x + 1
		case 'S':
			newY = r.y - 1
		case 'W':
			newX = r.x - 1
		}

		err := r.plateau.update(r.x, r.y, newX, newY, r)
		if err != nil {
			return err
		}

		r.x = newX
		r.y = newY
	}

	return nil
}

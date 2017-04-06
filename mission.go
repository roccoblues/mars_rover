package main

import (
	"errors"
)

type mission struct {
	rovers  []*rover
	plateau *plateau
}

func newMission(input []string) (*mission, error) {
	m := new(mission)
	if len(input) < 3 {
		return nil, errors.New("Invalid mission definition. Need at least plateau coordinates and one rover with commands.")
	}
	for i, _ := range input {
		if i == 0 {
			// first line is plateau size specification
			p, err := newPlateau(input[0])
			if err != nil {
				return nil, err
			}
			m.plateau = p
		} else if i%2 != 0 {
			// rover specifications are in odd rows
			// and the next line are the commands
			r, err := newRover(input[i], input[i+1])
			if err != nil {
				return nil, err
			}
			m.rovers = append(m.rovers, r)
		}
	}
	return m, nil
}

func (m *mission) run() error {
	for _, r := range m.rovers {
		err := r.deploy(m.plateau)
		if err != nil {
			return err
		}
		err = r.run()
		if err != nil {
			return err
		}
	}
	return nil
}

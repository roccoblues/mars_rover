package main

import (
	"errors"
)

type mission struct {
	Rovers  []*rover
	Plateau *plateau
}

func NewMission(input []string) (*mission, error) {
	m := new(mission)
	if len(input) < 3 {
		return nil, errors.New("Invalid mission definition. Need at least plateau coordinates and one rover with commands.")
	}
	for i, _ := range input {
		if i == 0 {
			// first line is plateau size specification
			p, err := NewPlateau(input[0])
			if err != nil {
				return nil, err
			}
			m.Plateau = p
		} else if i%2 != 0 {
			// rover specifications are in odd rows
			// and the next line are the commands
			r, err := NewRover(input[i], input[i+1])
			if err != nil {
				return nil, err
			}
			m.Rovers = append(m.Rovers, r)
		}
	}
	return m, nil
}

func (m *mission) Run() error {
	for _, r := range m.Rovers {
		if err := r.Deploy(m.Plateau); err != nil {
			return err
		}
		if err := r.Run(); err != nil {
			return err
		}
	}
	return nil
}

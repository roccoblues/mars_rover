package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type mission struct {
	plateau  *plateau
	rovers   []*rover
	commands [][]rune
}

func NewMission(r io.Reader) (*mission, error) {
	m := &mission{nil, nil, nil}

	i := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			// The first line is the plateau size specification.
			x, y, err := parsePlateauSize(line)
			if err != nil {
				return nil, err
			}
			m.plateau = NewPlateau(x, y)
		} else if i%2 != 0 {
			// Rover positions are in odd rows.
			x, y, d, err := parsePosition(line)
			if err != nil {
				return nil, err
			}
			m.rovers = append(m.rovers, NewRover(x, y, d))
		} else {
			// Even lines are the rover commands.
			c, err := parseCommands(line)
			if err != nil {
				return nil, err
			}
			m.commands = append(m.commands, c)
		}
		i++
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *mission) Run() error {
	// First deploy all rovers to plateau.
	for _, r := range m.rovers {
		err := r.Deploy(m.plateau)
		if err != nil {
			return err
		}
	}

	// Then run then sequentially.
	for i, r := range m.rovers {
		for _, c := range m.commands[i] {
			err := r.ApplyCommand(c)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *mission) Result() string {
	var out = new(bytes.Buffer)
	for _, r := range m.rovers {
		out.WriteString(fmt.Sprintf("%d %d %c\n", r.x, r.y, r.d))
	}

	return out.String()
}

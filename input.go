package main

import (
	"fmt"
	"strconv"
	"strings"
)

var validDirections = [4]rune{'N', 'E', 'S', 'W'}
var validCommands = [3]rune{'L', 'R', 'M'}

func convertCoordinates(c []string) (uint, uint, error) {
	if len(c) != 2 {
		return 0, 0, fmt.Errorf("Invalid coordinates: %s", c)
	}

	x, err := convertCoordinate(c[0])
	if err != nil {
		return 0, 0, err
	}

	y, err := convertCoordinate(c[1])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func convertCoordinate(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Invalid coordinate: %s", s)
	}
	if i <= 0 {
		return 0, fmt.Errorf("Invalid coordinate: %d", i)
	}
	return uint(i), nil
}

func convertDirection(s string) (rune, error) {
	d := rune(strings.ToUpper(s)[0])
	for _, v := range validDirections {
		if d == v {
			return d, nil
		}
	}
	return 'X', fmt.Errorf("Invalid direction: %s", s)
}

func convertCommands(s string) (commands []rune, err error) {
	for _, c := range strings.ToUpper(s) {
		if !validCommand(c) {
			return nil, fmt.Errorf("Invalid command: %c", c)
		}
		commands = append(commands, c)
	}
	return commands, nil
}

func validCommand(r rune) bool {
	for _, v := range validCommands {
		if r == v {
			return true
		}
	}
	return false
}

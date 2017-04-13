package main

import (
	"fmt"
	"strconv"
	"strings"
)

var validDirections = []rune{'N', 'E', 'S', 'W'}
var validCommands = []rune{'L', 'R', 'M'}

func parsePlateauSize(s string) (x int, y int, err error) {
	f := strings.Fields(s)
	if f == nil || len(f) != 2 {
		return 0, 0, fmt.Errorf("could't parse plateau size: %s", s)
	}

	x, y, err = convertCoordinates(f[0], f[1])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func parsePosition(s string) (x int, y int, direction rune, err error) {
	f := strings.Fields(s)
	if f == nil || len(f) != 3 {
		return 0, 0, 'X', fmt.Errorf("could't parse position: %s", s)
	}

	x, y, err = convertCoordinates(f[0], f[1])
	if err != nil {
		return 0, 0, 'X', err
	}

	direction = rune(f[2][0])
	if !validDirection(direction) {
		return 0, 0, 'X', fmt.Errorf("invalid rover position: %c", direction)
	}

	return x, y, direction, nil
}

func parseCommands(s string) (commands []rune, err error) {
	for _, c := range strings.ToUpper(s) {
		if !validCommand(c) {
			return nil, fmt.Errorf("invalid command: %c", c)
		}
		commands = append(commands, c)
	}
	return commands, nil
}

func convertCoordinates(sx string, sy string) (x int, y int, err error) {
	x, err = strconv.Atoi(sx)
	if err != nil {
		return 0, 0, err
	}
	y, err = strconv.Atoi(sy)
	if err != nil {
		return 0, 0, err
	}

	if x < 0 || y < 0 {
		return 0, 0, fmt.Errorf("invalid coordinates: %s %s", sx, sy)
	}

	return x, y, nil
}

func validCommand(r rune) bool {
	return runeInList(r, validCommands)
}

func validDirection(r rune) bool {
	return runeInList(r, validDirections)
}

func runeInList(r rune, l []rune) bool {
	for _, v := range l {
		if r == v {
			return true
		}
	}
	return false
}

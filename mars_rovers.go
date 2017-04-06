package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	input := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	err := scanner.Err()
	checkError(err)

	mission, err := newMission(input)
	checkError(err)

	err = mission.run()
	checkError(err)

	for _, r := range mission.rovers {
		fmt.Printf("%d %d %c\n", r.x, r.y, r.direction)
	}
}

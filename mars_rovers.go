package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	mission, err := NewMission(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = mission.Run(); err != nil {
		log.Fatal(err)
	}
	for _, r := range mission.Rovers {
		fmt.Printf("%d %d %c\n", r.X, r.Y, r.Direction)
	}
}

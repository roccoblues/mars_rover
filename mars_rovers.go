package main

import (
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
	m, err := NewMission(os.Stdin)
	checkError(err)

	err = m.Run()
	checkError(err)

	fmt.Print(m.Result())
}

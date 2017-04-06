package main

import "testing"

func TestNewMission(t *testing.T) {
    validMission := []string{"5 5", "1 2 N", "LMLMLMLMM", "3 3 E", "MMRMMRMRRM"}

    mission, err := newMission(validMission)
    if err != nil {
       t.Error("unexpected error received")
    }
    if mission == nil {
        t.Fatal("no mission created")
    }

    actual := len(mission.rovers)
    expected := 2
    if actual != expected {
        t.Errorf("expected %d rovers in mission, got %d", expected, actual)
    }
}

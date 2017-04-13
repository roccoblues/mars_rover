package main

import (
	"bytes"
	"testing"
)

var validMissionTests = []struct {
	in  string
	exp string
}{
	{"5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM\n", "1 3 N\n5 1 E\n"},      // example mission
	{"5 5\n1 2 N\nRM\n3 3 E\nMMR\n4 5 N\nLMM\n", "2 2 E\n5 3 S\n2 5 W\n"}, // 3 rovers
	{"5 5\n1 2 N\n\n3 3 E\n\n4 5 N\n\n", "1 2 N\n3 3 E\n4 5 N\n"},         // 3 rovers without commands
}

var parseErrorTests = []struct {
	in string
}{
	{"-5 5\n1 1 N\nM\n"}, // invalid plateau size
	{"x 5\n1 1 N\nM\n"},  // invalid plateau size
	{"5 5\n-1 1 N\nM\n"}, // invalid rover position
	{"5 5\nX 1 N\nM\n"},  // invalid rover position
	{"5 5\n1 1 F\nM\n"},  // invalid rover direction
	{"5 5\n1 1 F\nMX\n"}, // invalid rover command
}

var runErrorTests = []struct {
	in string
}{
	{"5 5\n5 6 N\nLM\n"},        // put rover outside of plateau
	{"5 5\n4 4 N\nMM\n"},        // drive rover of the map
	{"5 5\n1 1 N\nM\n1 1 S\nM"}, // put rover on top of another
	{"5 5\n1 1 N\nM\n1 2 N\nM"}, // crash rover into another
}

func TestValidMission(t *testing.T) {
	for _, tt := range validMissionTests {
		input := bytes.NewBufferString(string(tt.in))
		m, err := NewMission(input)
		if err != nil {
			t.Errorf("NewMission(%s) returned unexpected error: %s.", tt.in, err)
		}
		if m == nil {
			t.Errorf("NewMission(%s) didn't return a mission.", tt.in)
			continue
		}

		err = m.Run()
		if err != nil {
			t.Errorf("Mission:%s: Run() unexpected error received: %s", tt.in, err)
		}

		r := m.Result()
		if r != tt.exp {
			t.Errorf("Mission:%s Result() got:\n%sexpected:\n%s", tt.in, r, tt.exp)
		}
	}
}

func TestParseError(t *testing.T) {
	for _, tt := range parseErrorTests {
		input := bytes.NewBufferString(string(tt.in))
		m, err := NewMission(input)
		if err == nil {
			t.Errorf("NewMission(%s) expected error but got nil.", tt.in)
		}
		if m != nil {
			t.Errorf("NewMission(%s) returned a mission, expected nil.", tt.in)
		}
	}
}

func TestRunError(t *testing.T) {
	for _, tt := range runErrorTests {
		input := bytes.NewBufferString(string(tt.in))
		m, err := NewMission(input)
		if err != nil {
			t.Errorf("NewMission(%s) returned unexpected error: %s.", tt.in, err)
		}
		if m == nil {
			t.Errorf("NewMission(%s) didn't return a mission.", tt.in)
			continue
		}

		err = m.Run()
		if err == nil {
			t.Errorf("Mission:%s: Run() succeeded but we expected an error.", tt.in)
		}
	}
}

package midi

import (
	"encoding/json"
	"testing"
)

func TestParseNoteOn(t *testing.T) {
	// GIVEN
	input := make(chan byte, 6)
	defer close(input)
	input <- 0x90
	input <- 0x3C
	input <- 0x40
	input <- 0x90
	input <- 0x3C
	input <- 0x40

	// WHEN
	output := Parse(input)

	// THEN
	expected := Message{NoteOn, 60, 64}
	found := <-output
	if found != expected {
		expected_out, _ := json.Marshal(expected)
		found_out, _ := json.Marshal(found)
		t.Errorf("Expected %s, found %s", string(expected_out), string(found_out))
	}

	expected = Message{NoteOn, 60, 64}
	found = <-output
	if found != expected {
		expected_out, _ := json.Marshal(expected)
		found_out, _ := json.Marshal(found)
		t.Errorf("Expected %s, found %s", string(expected_out), string(found_out))
	}

}

func TestParseNoteOff(t *testing.T) {
	// GIVEN
	input := make(chan byte, 3)
	defer close(input)
	input <- 0x80
	input <- 0x3C
	input <- 0x40

	// WHEN
	output := Parse(input)

	// THEN
	expected := Message{NoteOff, 60, 64}
	found := <-output

	if found != expected {
		expected_out, _ := json.Marshal(expected)
		found_out, _ := json.Marshal(found)
		t.Errorf("Expected %s, found %s", string(expected_out), string(found_out))
	}
}

func TestParseNoteOnWithVelocity0(t *testing.T) {
	// GIVEN
	input := make(chan byte, 3)
	defer close(input)
	input <- 0x90
	input <- 0x3C
	input <- 0x00

	// WHEN
	output := Parse(input)

	// THEN
	expected := Message{NoteOff, 60, 0}
	found := <-output

	if found != expected {
		expected_out, _ := json.Marshal(expected)
		found_out, _ := json.Marshal(found)
		t.Errorf("Expected %s, found %s", string(expected_out), string(found_out))
	}
}

func TestParseRunningStatus(t *testing.T) {
	input := make(chan byte, 7)
	defer close(input)
	input <- 0x90
	input <- 0x3C
	input <- 0x7F
	input <- 0x7F
	input <- 0x41
	input <- 0x3C
	input <- 0x00

	// WHEN
	output := Parse(input)
	// THEN
	expected := Message{NoteOn, 60, 127}
	found := <-output
	if found != expected {
		expected_out, _ := json.Marshal(expected)
		found_out, _ := json.Marshal(found)
		t.Errorf("Expected %s, found %s", string(expected_out), string(found_out))
	}

	expected = Message{NoteOn, 127, 65}
	found = <-output
	if found != expected {
		expected_out, _ := json.Marshal(expected)
		found_out, _ := json.Marshal(found)
		t.Errorf("Expected %s, found %s", string(expected_out), string(found_out))
	}

	expected = Message{NoteOff, 60, 0}
	found = <-output
	if found != expected {
		expected_out, _ := json.Marshal(expected)
		found_out, _ := json.Marshal(found)
		t.Errorf("Expected %s, found %s", string(expected_out), string(found_out))
	}
}

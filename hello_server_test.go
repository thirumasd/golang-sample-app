package main

import "testing"

func TestGreetingSpecific(t *testing.T) {
	greeting := CreateGreeting("John")
	if greeting != "Hello, John\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingDefault(t *testing.T) {
	greeting := CreateGreeting("")
	ifgreeting != "Hello, Guest\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
	}
}

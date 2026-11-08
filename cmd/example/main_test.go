package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("nix")
	want := "Hello, nix"
	if got != want {
		t.Fatalf("Greeting(%q) = %q, want %q", "nix", got, want)
	}
}

func TestGreetingEmpty(t *testing.T) {
	got := Greeting("")
	want := "Hello, world"
	if got != want {
		t.Fatalf("Greeting(%q) = %q, want %q", "", got, want)
	}
}

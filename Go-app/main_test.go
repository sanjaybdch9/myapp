// main_test.go
package main

import (
    "testing"
)

func TestHelloWorld(t *testing.T) {
    expected := "Hello, World!"
    result := HelloWorld()
    if result != expected {
        t.Errorf("Expected %s but got %s", expected, result)
    }
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 2)
    expected := 3
    if result != expected {
        t.Errorf("Subtract(5, 2) returned %d, expected %d", result, expected)
    }
}

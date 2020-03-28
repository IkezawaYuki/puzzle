package main

import (
	"fmt"
	"testing"
)

func TestPleaseConform(t *testing.T) {
	caps1 := []string{"F", "F", "B", "B", "F", "F", "B", "B"}
	caps2 := []string{"F", "B", "B", "B", "F", "B", "B", "F"}
	fmt.Println("--- caps1 ---")
	pleaseConform(caps1)
	fmt.Println("--- caps2 ---")
	pleaseConform(caps2)
}
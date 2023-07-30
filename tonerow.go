package main

import (
	"fmt"
	"math/rand"
	"strings"
)


func main() {
	// Create a random 12-tone row
	row := make([]int, 12)
	for i := 0; i < 12; i++ {
		row[i] = i
	}
	for i := 0; i < 12; i++ {
		for j := i+1; j < 11; j++ {
			if rand.Float64() < 0.5 {
				row[i], row[j] = row[j], row[i]
			}
		}
	}
	
	// Generate its inversion
	inv := invert(row)

	// Generate its retrograde
	ret := retrograde(row)

	// Generate its retrograde inversion
	ri := invert(retrograde(row))

	// Print all four
	printNotes("Row:", row)
	printNotes("Inv:", inv)
	printNotes("Ret:", ret)
	printNotes("RI: ", ri)
}

func invert(row []int) []int {
	inv := make([]int, 12)
	inv[0] = row[0]
	for i := 1; i < 12; i++ {
		interval := row[i-1] - row[i]
		inv[i] = inv[i-1] + interval
		if inv[i] < 0 {
			inv[i] += 12
		}
		if inv[i] > 11 {
			inv[i] -= 12
		}
	}
	return inv
}

func retrograde(row []int) []int {
	ret := make([]int, 12)
	j := 12
	for i := 0; i < 12; i++ {
		j--
		ret[i] = row[j]
	}
	return ret
}

func printNotes(label string, row []int) {
	notes := []string{
		"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B",
	}
	parts := make([]string, 12)
	for i := 0; i < 12; i++ {
		parts[i] = notes[row[i]]
	}
	line := strings.Join(parts, " ")
	fmt.Println(label, line)
}
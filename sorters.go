package main

import (
	"time"
)

type Sorter interface {
	start(output Visualizer)
}

type Bubblesort struct {
}

func (*Bubblesort) start(output Visualizer) {
	values := output.GetSlice()
	n := len(values)
	swapped := true
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 1; i < n; i++ {
			// if the current element is greater than the next
			// element, swap them
			if values[i-1] > values[i] {
				time.Sleep(100 * time.Millisecond)
				output.SwitchPositions(i, i-1)
				swapped = true
			}
		}
	}
}

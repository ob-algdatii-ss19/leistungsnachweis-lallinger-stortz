package sortingalgorithms

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/visualization"
	"time"
)

type Sorter interface {
	Start(output visualization.Visualizer)
}

type Bubblesort struct {
}

func (*Bubblesort) Start(output visualization.Visualizer) {
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
				time.Sleep(10 * time.Millisecond)
				output.SwitchPositions(i, i-1)
				swapped = true
			}
		}
	}
}

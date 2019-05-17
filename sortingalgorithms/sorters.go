package sortingalgorithms

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/visualization"
	"time"
)

var sleeptime = 100 * time.Millisecond

type Sorter interface {
	Start(output visualization.Visualizer)
}

type Bubblesort struct {
}

type Insertionsort struct {
}

type Selectionsort struct {
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

				output.SwitchPositions(i, i-1, -1)
				time.Sleep(sleeptime)
				swapped = true
			}
		}
	}
}

func (*Insertionsort) Start(output visualization.Visualizer) {
	values := output.GetSlice()
	var n = len(values)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if values[j-1] > values[j] {
				output.SwitchPositions(j, j-1, i)
				time.Sleep(sleeptime)
			}
			j = j - 1
		}
	}
}

func (*Selectionsort) Start(output visualization.Visualizer) {
	values := output.GetSlice()
	var n = len(values)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if values[j] < values[minIdx] {
				minIdx = j
			}
		}
		output.SwitchPositions(i, minIdx, i)
		time.Sleep(sleeptime)
	}
}

package sortingalgorithms

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/visualization"
	"math/rand"
)

type Sorter interface {
	Start(output visualization.Visualizer)
}

type Bubblesort struct {
	values []int
}

type Insertionsort struct {
	values []int
}

type Selectionsort struct {
	values []int
}

type Bogosort struct {
	values []int
}

type Quicksort struct {
	values  []int
	indeces map[int]int
	out     visualization.Visualizer
}

func (me *Bubblesort) Start(output visualization.Visualizer) {
	values := output.GetSlice()
	me.values = values
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

				if output.SwitchPositions(i, i-1, -1) {
					return
				}
				swapped = true
			}
		}
	}
}

func (me *Insertionsort) Start(output visualization.Visualizer) {
	values := output.GetSlice()
	me.values = values
	var n = len(values)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if values[j-1] > values[j] {
				if output.SwitchPositions(j, j-1, i) {
					return
				}
			}
			j = j - 1
		}
	}
}

func (me *Selectionsort) Start(output visualization.Visualizer) {
	values := output.GetSlice()
	me.values = values
	var n = len(values)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if values[j] < values[minIdx] {
				minIdx = j
			}
		}
		if output.SwitchPositions(i, minIdx, i) {
			return
		}
	}
}

func (me *Bogosort) Start(output visualization.Visualizer) {
	values := output.GetSlice()
	me.values = values

	for {
		if output.SwitchPositions(rand.Intn(len(values)), rand.Intn(len(values)), -1) {
			return
		}

		for i := 0; i < len(values); i++ {
			if i+1 == len(values) {
				return
			}
			if values[i] > values[i+1] {
				break
			}
		}
	}
}

func (me *Quicksort) Start(output visualization.Visualizer) {
	values := output.GetSlice()
	me.values = values

	me.out = output
	me.indeces = make(map[int]int)

	for k, v := range me.values {
		me.indeces[v] = k
	}
	me.quicksort(values)

}

func (me *Quicksort) quicksort(a []int) bool {

	stop := false
	if len(a) < 2 {
		return stop
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	if me.indeces[a[pivot]] != me.indeces[a[right]] {
		stop = me.out.SwitchPositions(me.indeces[a[pivot]], me.indeces[a[right]], -1)
		if stop {
			return stop
		}
		me.indeces[a[pivot]], me.indeces[a[right]] = me.indeces[a[right]], me.indeces[a[pivot]]
	}

	for i, _ := range a {
		if a[i] < a[right] {
			if me.indeces[a[left]] != me.indeces[a[i]] {
				stop = me.out.SwitchPositions(me.indeces[a[left]], me.indeces[a[i]], -1)
			}
			if stop {
				return stop
			}
			me.indeces[a[left]], me.indeces[a[i]] = me.indeces[a[i]], me.indeces[a[left]]
			left++
		}
	}

	if me.indeces[a[left]] != me.indeces[a[right]] {
		stop = me.out.SwitchPositions(me.indeces[a[left]], me.indeces[a[right]], -1)
		me.indeces[a[left]], me.indeces[a[right]] = me.indeces[a[right]], me.indeces[a[left]]
	}

	stop = me.quicksort(a[:left])
	if stop {
		return stop
	}
	stop = me.quicksort(a[left+1:])

	return stop
}

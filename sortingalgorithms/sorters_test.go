package sortingalgorithms

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/visualization"
	"testing"
	"time"
)

const n = 100

func TestBubble(t *testing.T) {
	visual := &visualization.ShellVisualizer{Test: n}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Bubblesort{}
	sorter.Start(visual)
	for k, v := range sorter.values {
		if k > 0 && v < sorter.values[k-1] {
			t.Error("Bubblesort array is not sorted")
		}
	}
	visual.Clear()
}

func TestInsertion(t *testing.T) {
	visual := &visualization.ShellVisualizer{Test: n}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Insertionsort{}
	sorter.Start(visual)
	for k, v := range sorter.values {
		if k > 0 && v < sorter.values[k-1] {
			t.Error("Insertionsort array is not sorted")
		}
	}
	visual.Clear()
}

func TestSelection(t *testing.T) {
	visual := &visualization.ShellVisualizer{Test: n}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Selectionsort{}
	sorter.Start(visual)
	for k, v := range sorter.values {
		if k > 0 && v < sorter.values[k-1] {
			t.Error("Selectionsort array is not sorted")
		}
	}
	visual.Clear()
}

func TestBogo(t *testing.T) {
	visual := &visualization.ShellVisualizer{Test: 4}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Bogosort{}
	sorter.Start(visual)
	for k, v := range sorter.values {
		if k > 0 && v < sorter.values[k-1] {
			t.Error("Bogosort array is not sorted")
		}
	}
	visual.Clear()
}

func TestQuick(t *testing.T) {
	visual := &visualization.ShellVisualizer{Test: n}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Quicksort{}
	sorter.Start(visual)
	for k, v := range sorter.values {
		if k > 0 && v < sorter.values[k-1] {
			t.Error("Quicksort array is not sorted")
		}
	}
	visual.Clear()
}

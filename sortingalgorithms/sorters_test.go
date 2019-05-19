package sortingalgorithms

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/visualization"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const n = 100

func TestBubble(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	values := rand.Perm(n)
	cp := make([]int, n)
	copy(cp, values)

	visual := &visualization.ShellVisualizer{Values: values, Test: true}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Bubblesort{values: values}
	sorter.Start(visual)
	sort.Ints(cp)
	if !isIdentical(cp, sorter.values) {
		t.Error("Bubblesort arrays are not identical")
	}
	visual.Clear()
}

func TestInsertion(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	values := rand.Perm(n)
	cp := make([]int, n)
	copy(cp, values)

	visual := &visualization.ShellVisualizer{Values: values, Test: true}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Insertionsort{values: values}
	sorter.Start(visual)
	sort.Ints(cp)
	if !isIdentical(cp, sorter.values) {
		t.Error("Insertionsort arrays are not identical")
	}
	visual.Clear()
}

func TestSelection(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	values := rand.Perm(n)
	cp := make([]int, n)
	copy(cp, values)

	visual := &visualization.ShellVisualizer{Values: values, Test: true}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Selectionsort{values: values}
	sorter.Start(visual)
	sort.Ints(cp)
	if !isIdentical(cp, sorter.values) {
		t.Error("Selectionsort arrays are not identical")
	}
	visual.Clear()
}

func TestBogo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	values := rand.Perm(4)
	cp := make([]int, 4)
	copy(cp, values)

	visual := &visualization.ShellVisualizer{Values: values, Test: true}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Bogosort{values: values}
	sorter.Start(visual)
	sort.Ints(cp)
	if !isIdentical(cp, sorter.values) {
		t.Error("Bubblesort arrays are not identical")
	}
	visual.Clear()
}

func TestQuick(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	values := rand.Perm(4)
	cp := make([]int, 4)
	copy(cp, values)

	visual := &visualization.ShellVisualizer{Values: values, Test: true}
	visual.Init(false, 0*time.Millisecond)
	sorter := &Quicksort{values: values}
	sorter.Start(visual)
	sort.Ints(cp)
	if !isIdentical(cp, sorter.values) {
		t.Error("Bubblesort arrays are not identical")
	}
	visual.Clear()
}

func isIdentical(first []int, second []int) bool {
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

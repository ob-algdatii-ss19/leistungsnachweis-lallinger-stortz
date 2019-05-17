package main

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/sortingalgorithms"
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/visualization"
)

func main() {

	var sorter sortingalgorithms.Sorter
	var visual visualization.Visualizer
	visual = &visualization.ShellVisualizer{}

	a := "Selectionsort"
	switch a {
	case "Bubblesort":
		sorter = &sortingalgorithms.Bubblesort{}
	case "Insertionsort":
		sorter = &sortingalgorithms.Insertionsort{}
	case "Selectionsort":
		sorter = &sortingalgorithms.Selectionsort{}
	}

	sorter.Start(visual)
}

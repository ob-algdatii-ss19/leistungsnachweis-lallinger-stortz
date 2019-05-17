package main

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/sortingalgorithms"
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/visualization"
)

func main() {

	var sorter sortingalgorithms.Sorter
	var visual visualization.Visualizer
	visual = &visualization.ShellVisualizer{}

	a := "Bubblesort"
	switch a {
	case "Bubblesort":
		sorter = &sortingalgorithms.Bubblesort{}
	}

	sorter.Start(visual)
}

package main

func main() {

	var sorter Sorter
	var visual Visualizer
	visual = &ShellVisualizer{}

	a := "Bubblesort"
	switch a {
	case "Bubblesort":
		sorter = &Bubblesort{}
	}

	sorter.start(visual)
}

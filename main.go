package main

import (
	"flag"
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/sortingalgorithms"
	"github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz/visualization"
	"time"
)

var (
	sleepTime = flag.Int("sleep", 100, "Time to wait between steps while automatic stepping in miliseconds")
	stepping  = flag.Bool("step", false, "Activate manual stepping")
)

func main() {

	origUsage := flag.Usage
	flag.Usage = func() {
		fmt.Println("This program visualizes sorting algorithms. You can choose a sorting algorithm and stepping mode (manual or automatic).")
		fmt.Println("While automatic stepping is on you can press space to switch into manual stepping mode.")
		fmt.Println("When in manual stepping mode you can step with space or go back to automatic stepping by pressing 'c'.")
		fmt.Println("To exit the program press ctrl+c at any moment.")
		fmt.Println("Resizing the terminal after starting is not possible. Please set you desired window size before starting")
		fmt.Println("leistungsnachweis-lalllinger-stortz [FLAGS] SORTER")
		fmt.Println("FLAGS")
		origUsage()
		fmt.Println()
		fmt.Println("SORTER")
		fmt.Println("  bubble")
		fmt.Println("	Sort values using bubblesort")
		fmt.Println("  insertion")
		fmt.Println("	Sort values using insertionsort")
		fmt.Println("  selection")
		fmt.Println("	Sort values using selectionsort")
		fmt.Println("  quick")
		fmt.Println("	Sort values using quicksort")
		fmt.Println("  bogo")
		fmt.Println("	Sort values using bogosort. Caution this probably won't terminate by itself!")
		fmt.Println("")
		fmt.Println("Example:")
		fmt.Println("  leistungsnachweis-lalllinger-stortz --sleep=50 --step=false selection")
	}

	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Invalid number of arguments!")
		flag.Usage()
		return
	}

	var sorter sortingalgorithms.Sorter
	var visual visualization.Visualizer
	visual = &visualization.ShellVisualizer{}

	switch flag.Args()[0] {
	case "bubble":
		sorter = &sortingalgorithms.Bubblesort{}
	case "insertion":
		sorter = &sortingalgorithms.Insertionsort{}
	case "selection":
		sorter = &sortingalgorithms.Selectionsort{}
	case "bogo":
		sorter = &sortingalgorithms.Bogosort{}
	case "quick":
		sorter = &sortingalgorithms.Quicksort{}
	}

	visual.Init(*stepping, time.Duration(*sleepTime)*time.Millisecond)
	sorter.Start(visual)
	visual.Clear()
}

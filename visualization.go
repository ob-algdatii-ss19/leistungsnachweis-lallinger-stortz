package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"math"
	"math/rand"
	"os"
)

type Visualizer interface {
	SwitchPositions(first int, second int)
	GetSlice() []int
}

type ShellVisualizer struct {
	values []int
	s      tcell.Screen
}

var style = tcell.StyleDefault

var screen tcell.Screen
var quit chan struct{}
var clear map[string]func() //create a map for storing clear funcs

const halfslab = '▄'

func (me *ShellVisualizer) drawSlice() {

	st := tcell.StyleDefault
	const gl = ' '
	for x := 0; x < len(me.values); x++ {
		val := me.values[x]
		for y := 0; y < len(me.values); y++ {
			difVal := len(me.values) - (val + 1)
			if difVal == y {
				st = st.Background(tcell.ColorGrey)
				st = st.Foreground(tcell.ColorBlue)
				me.s.SetCell(x, y, st, rune((val%26)+65))
			} else {
				if y > difVal {
					st = st.Background(tcell.ColorWhite)
				} else {
					st = st.Background(tcell.ColorBlack)
				}
				me.s.SetCell(x, y, st, gl)
			}
		}
	}

	me.s.Show()

}

func (me *ShellVisualizer) GetSlice() []int {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	var e error
	me.s, e = tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	encoding.Register()

	if e = me.s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	quit = make(chan struct{})
	//go pollEvents(s)

	me.s.Show()

	w, h := me.s.Size()
	values := rand.Perm(int(math.Min(float64(w), float64(h))))
	for i := range values {
		values[i]++
	}

	me.values = values

	me.drawSlice()
	return values
}

func (me *ShellVisualizer) SwitchPositions(first int, second int) {
	me.values[first], me.values[second] = me.values[second], me.values[first]
	me.drawSlice()
}

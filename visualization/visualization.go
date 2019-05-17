package visualization

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
	values      []int
	s           tcell.Screen
	doubleWidth bool
}

var style = tcell.StyleDefault

var screen tcell.Screen
var quit chan struct{}
var clear map[string]func() //create a map for storing clear funcs

func (me *ShellVisualizer) drawSlice() {

	st := tcell.StyleDefault
	sym := ' '
	for x := 0; x < len(me.values); x++ {
		val := me.values[x]
		useHalf := false
		factor := 1
		if me.doubleWidth {
			if val%2 == 0 {
				val = val / 2
				factor = 2
			} else {
				val = (val - 1) / 2
				useHalf = true
				factor = 2
			}
		}
		difVal := len(me.values)/factor - val - 1
		for y := 0; y < len(me.values)/factor; y++ {
			if y > difVal {
				st = st.Background(tcell.ColorWhite)
				sym = ' '
			} else if y == difVal && useHalf {
				st = st.Background(tcell.ColorBlack)
				st = st.Foreground(tcell.ColorWhite)
				sym = '▄'
			} else {
				st = st.Background(tcell.ColorBlack)
				sym = ' '
			}

			me.s.SetCell(x, y, st, sym)
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
	var values []int
	w, h := me.s.Size()
	if w >= 2*h {
		me.doubleWidth = true
		values = rand.Perm(2 * h)
	} else {
		me.doubleWidth = false
		values = rand.Perm(int(math.Min(float64(w), float64(h))))
	}

	me.values = values
	me.drawSlice()
	return values
}

func (me *ShellVisualizer) SwitchPositions(first int, second int) {
	me.values[first], me.values[second] = me.values[second], me.values[first]
	me.drawSlice()
}

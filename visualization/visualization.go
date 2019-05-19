package visualization

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"math/rand"
	"os"
	"time"
)

type Visualizer interface {
	SwitchPositions(first int, second int, green int) bool
	GetSlice() []int
	Init(step bool, sleeptime time.Duration)
	Clear()
}

type ShellVisualizer struct {
	Values    []int
	s         tcell.Screen
	green     int
	stepping  bool
	sleeptime time.Duration
	stop      bool
	doStep    bool
	keyPress  bool
	Test      bool
	ratio     int
	heigth    int
}

const quarter = '▂'
const threequarters = '▆'

var style = tcell.StyleDefault

var screen tcell.Screen
var clear map[string]func() //create a map for storing clear funcs

func (me *ShellVisualizer) Clear() {
	me.s.Clear()
	me.s.Fini()
}

func (me *ShellVisualizer) Init(step bool, sleeptime time.Duration) {

	me.stepping = step
	me.sleeptime = sleeptime
	me.stop = false
	me.doStep = false
	me.keyPress = false
	me.green = -1

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
	st := tcell.StyleDefault
	st.Background(tcell.ColorGray)
	me.s.Fill(' ', st)
	me.s.HideCursor()
	me.s.DisableMouse()
	me.s.Resize(200, 200, 200, 200)
	me.s.Clear()
	me.s.Show()
	go func() {
		for {
			ev := me.s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Rune() == ' ' {
					me.stepping = true
					me.doStep = !me.doStep
				}
				if ev.Rune() == 'c' {
					me.keyPress = true
					me.stepping = false
				}
				if ev.Key() == tcell.KeyCtrlC {
					me.keyPress = true
					me.stop = true
				}
			case *tcell.EventResize:
				me.s.Sync()
			}
		}
	}()

}

func (me *ShellVisualizer) drawSlice(indeces ...int) {

	first := -1
	second := -1
	if len(indeces) == 2 {
		first = indeces[0]
		second = indeces[1]
	}

	st := tcell.StyleDefault
	sym := ' '
	for x := 0; x < len(me.Values); x++ {
		val := me.Values[x]
		remainder := val % me.ratio

		if remainder == 0 {
			val = val / me.ratio
		} else {
			val = (val - remainder) / me.ratio
		}

		difVal := me.heigth - val - 1
		//for y := 0; y < len(me.Values)/me.ratio; y++ {
		for y := 0; y < me.heigth; y++ {
			if y > difVal {
				if me.green == x {
					st = st.Background(tcell.ColorGreen)
				} else if first == x || second == x {
					st = st.Background(tcell.ColorRed)
				} else {
					st = st.Background(tcell.ColorWhite)
				}

				sym = ' '
			} else if y == difVal && remainder != 0 {
				st = st.Background(tcell.ColorBlack)
				if me.green == x {
					st = st.Foreground(tcell.ColorGreen)
				} else if first == x || second == x {
					st = st.Foreground(tcell.ColorRed)
				} else {
					st = st.Foreground(tcell.ColorWhite)
				}

				sym = '▄'
				if me.ratio == 4 {
					if remainder == 1 {
						sym = quarter
					} else if remainder == 3 {
						sym = threequarters
					}
				}
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
	var values []int
	w, h := me.s.Size()
	me.heigth = h
	rand.Seed(time.Now().UnixNano())
	me.ratio = 1
	if w >= 2*h {
		me.ratio = 4
	} else if w >= h {
		me.ratio = 2
	}
	values = rand.Perm(w)

	me.Values = values
	me.drawSlice()
	return values
}

func (me *ShellVisualizer) SwitchPositions(first int, second int, green int) bool {
	if green >= 0 {
		me.green = green
	}

	if me.stepping {
		for !me.doStep && !me.keyPress {
			time.Sleep(10 * time.Millisecond)
		}
		me.keyPress = false
		me.doStep = false
	} else {
		time.Sleep(me.sleeptime)
	}
	me.Values[first], me.Values[second] = me.Values[second], me.Values[first]
	if !me.Test {
		me.drawSlice(first, second)
	}
	return me.stop
}

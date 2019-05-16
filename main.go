package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	//"github.com/fatih/color"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var style = tcell.StyleDefault

var screen tcell.Screen
var quit chan struct{}
var clear map[string]func() //create a map for storing clear funcs

const halfslab = '▄'

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	encoding.Register()

	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	quit = make(chan struct{})
	//go pollEvents(s)

	s.Show()

	go func() {
		for {
			drawScreen(s)
			time.Sleep(time.Millisecond * 10000)
		}

	}()

	<-quit
	s.Fini()
}

func drawScreen(s tcell.Screen) {

	w, h := s.Size()

	if w == 0 || h == 0 {
		return
	}

	list := rand.Perm(h)
	for i, _ := range list {
		list[i]++
	}
	sel1 := rand.Intn(h) + 1
	sel2 := rand.Intn(h) + 1
	st := tcell.StyleDefault
	const gl = ' '
	for x := 0; x < h; x++ {
		val := list[x]
		for y := 0; y < h; y++ {
			/*if y < val{
				st = st.Background(tcell.ColorGrey)
				st = st.Foreground(tcell.ColorBlue)
				s.SetCell(x, y, st, rune((y % 26)+65))
			}else{
				st = st.Background(tcell.ColorBlack)
				s.SetCell(x, y, st, gl)
			}*/
			difVal := h - (val + 1)
			if difVal == y {
				st = st.Background(tcell.ColorGrey)
				st = st.Foreground(tcell.ColorBlue)
				s.SetCell(x, y, st, rune((val%26)+65))

			} else {
				if y > difVal {
					if difVal == sel1 {
						st = st.Background(tcell.ColorRed)
					} else if difVal == sel2 {
						st = st.Background(tcell.ColorGreen)
					} else {
						st = st.Background(tcell.ColorWhite)
					}
				} else {
					st = st.Background(tcell.ColorBlack)
				}

				s.SetCell(x, y, st, gl)
			}

		}
	}

	s.Show()
}

package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	if 1 >= len(os.Args) {
		fmt.Printf("usage: countdown <30m22s>\n")
		os.Exit(1)
	}

	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		log.Fatalf("do not understand input: %s", err)
	}

	go func() {
		w := app.NewWindow()
		err := run(w, d)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window, d time.Duration) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	end := time.Now().Truncate(time.Second).Add(d)
	go func() {
		for {
			time.Sleep(time.Second)
			w.Invalidate()
		}
	}()

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			paint.Fill(&ops, color.NRGBA{R: 0, G: 0, B: 0, A: 255})

			duration := end.Sub(time.Now().Truncate(time.Second))

			title := material.H1(th, duration.String())
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			e.Frame(gtx.Ops)

		case key.Event:
			if e.Name == "⎋" {
				return nil
			}
			if e.Name == "Q" {
				return nil
			}
			if e.Name == "R" {
				end = time.Now().Truncate(time.Second).Add(d)
			}
		}
	}
}

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
)

type (
	SymbolCell struct {
		BoxContainer     *fyne.Container
		SymbolTextCanvas *canvas.Text
		XAxis            float32
	}

	CursorTape struct {
		Widget    *fyne.Container
		XAxis     float32
		StateBind binding.String
		State     string
	}
)

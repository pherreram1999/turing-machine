package main

import (
	"bytes"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

//go:embed TM.png
var diagram []byte

func OpenTuringDiagram() {
	win := App.NewWindow("Turing Diagram")
	win.Resize(fyne.NewSize(800, 300))
	imgReader := bytes.NewReader(diagram)
	img := canvas.NewImageFromReader(imgReader, "TM.png")
	img.FillMode = canvas.ImageFillContain
	win.SetContent(img)
	win.Show()
}

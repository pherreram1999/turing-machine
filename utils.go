package main

import (
	"bytes"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

//go:embed cursor.png
var cursorImage []byte

func NewCursor(stateBind binding.String) *fyne.Container {
	imgCanvas := canvas.NewImageFromReader(bytes.NewReader(cursorImage), "cursor.png")
	imgCanvas.FillMode = canvas.ImageFillOriginal
	cursorLbl := widget.NewLabel("")
	cursorLbl.Bind(stateBind)
	cursorLbl.Move(fyne.NewPos(22, 3))
	cursorCont := container.NewWithoutLayout(imgCanvas, cursorLbl)
	imgCanvas.Resize(fyne.NewSize(80, 80))
	return cursorCont
}

func NewBoxSymbol(symbol string) *fyne.Container {
	boxBackground := canvas.NewRectangle(color.RGBA{238, 242, 255, 255})
	boxSymbol := canvas.NewText(symbol, color.Black)
	boxCont := container.NewWithoutLayout()
	boxCont.Add(boxBackground)
	boxCont.Add(boxSymbol)
	boxBackground.Resize(fyne.NewSize(boxSize, boxSize))
	boxCont.Resize(fyne.NewSize(boxSize, boxSize))
	middle := float32(boxSize / 2)
	boxSymbol.Move(fyne.NewPos(middle, middle-10))
	return boxCont
}

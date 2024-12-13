package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"regexp"
	"strings"
)

var Win fyne.Window

func main() {

	a := app.NewWithID("turing-machine")
	inputWordLbl := widget.NewLabel("Word")
	inputWord := widget.NewEntry()
	inputWord.SetPlaceHolder("Enter word")
	inputString := binding.NewString()
	inputWord.Bind(inputString)

	regex, err := regexp.Compile("01")

	if err != nil {
		panic(err)
	}

	cursorStateBind := binding.NewString()
	_ = cursorStateBind.Set("q0")

	cursorWidget := NewCursor(cursorStateBind)

	cursorTape := &CursorTape{
		Widget:    cursorWidget,
		StateBind: cursorStateBind,
		State:     "q0",
	}

	tmContainer := container.NewWithoutLayout(cursorWidget)

	cursorWidget.Move(fyne.NewPos(-cursorOffset, cursorTop))

	var tape []*TapeCell

	buildBtn := widget.NewButton("Build", func() {
		word, _ := inputString.Get()
		word = strings.TrimSpace(word)
		if !regex.MatchString(word) {
			dialog.ShowError(errors.New("No es una cadena valida de 0^n1^n"), Win)
			return
		}

		var xAxis float32 = 0
		for _, s := range word {
			symbolBind := binding.NewString()
			_ = symbolBind.Set(string(s))
			boxSymbol := NewBoxSymbol(symbolBind)
			boxSymbol.Move(fyne.NewPos(xAxis, tapeTop))
			tmContainer.Add(boxSymbol)

			tape = append(tape, &TapeCell{
				XAxis:        xAxis,
				SymbolBind:   symbolBind,
				BoxContainer: boxSymbol,
				Symbol:       string(s),
			})

			xAxis += boxSize

		}
	})

	animateBtn := widget.NewButton("Animate", func() {
		if len(tape) == 0 {
			dialog.ShowError(errors.New("no elements in tape"), Win)
			return
		}
		turingAnimate(cursorTape, &tape)
	})

	Win = a.NewWindow("Turing Machine")
	Win.Resize(fyne.NewSize(1200, 800))

	menuCont := container.NewVBox(
		inputWordLbl,
		inputWord,
		buildBtn,
		animateBtn,
	)

	title := canvas.NewText("Turing Machine", color.Black)
	title.TextStyle.Bold = true

	tmContainer.Add(title)

	cont := container.NewWithoutLayout(menuCont, tmContainer)

	tmContainer.Move(fyne.NewPos(320, 0))

	menuCont.Resize(fyne.NewSize(300, 500))

	Win.SetContent(cont)
	Win.ShowAndRun()
}

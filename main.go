package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var Win fyne.Window
var App fyne.App

var slideDurationBind binding.Float

const LimitAnimation = 10

const MaxRand = 20

func main() {

	App = app.NewWithID("turing-machine")
	inputWordLbl := widget.NewLabel("Word")
	inputWord := widget.NewEntry()
	inputWord.SetPlaceHolder("Enter word")
	inputString := binding.NewString()
	inputWord.Bind(inputString)

	slideDuration := widget.NewSlider(1, 20)
	slideDurationBind = binding.NewFloat()
	_ = slideDurationBind.Set(10)
	slideDuration.Bind(slideDurationBind)

	// bind counter
	strLenBind := binding.NewString()
	strLenLbl := widget.NewLabel("")
	strLenLbl.Bind(strLenBind)
	// counter info widget
	strLenCont := container.New(
		layout.NewHBoxLayout(),
		canvas.NewText("Str Length", color.Black),
		strLenLbl,
	)

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

	inputWord.OnChanged = func(s string) {
		cursorTape.StrLen = len(s)
		strLenBind.Set(fmt.Sprintf("%d", cursorTape.StrLen))
	}

	// turing machine container
	tmContainer := container.NewWithoutLayout(cursorWidget)

	cursorWidget.Move(fyne.NewPos(-cursorOffset, cursorTop))

	var tape []*TapeCell

	animateBtn := widget.NewButton("Animate", func() {
		word, _ := inputString.Get()
		word = strings.TrimSpace(word)
		if !regex.MatchString(word) {
			dialog.ShowError(errors.New("No es una cadena valida de 0^n1^n"), Win)
			return
		}

		var xAxis float32 = 0
		// antes de vaciar la cinta quitamos los nodos
		for _, cell := range tape {
			if cell.BoxContainer != nil {
				cell.BoxContainer.Hide()
				cell.BoxContainer = nil
			}
		}
		tape = []*TapeCell{} // vaciamos el arreglo
		cursorTape.Reset()
		for _, s := range word {

			cell := &TapeCell{
				XAxis:  xAxis,
				Symbol: string(s),
			}

			if cursorTape.StrLen <= LimitAnimation {
				symbolBind := binding.NewString()
				_ = symbolBind.Set(string(s))
				boxSymbol := NewBoxSymbol(symbolBind)
				boxSymbol.Move(fyne.NewPos(xAxis, tapeTop))
				tmContainer.Add(boxSymbol)
				xAxis += boxSize
				cell.SymbolBind = symbolBind
				cell.BoxContainer = boxSymbol
			}
			tape = append(tape, cell)
		}

		if cursorTape.StrLen <= LimitAnimation {
			time.Sleep(time.Second)
		}

		_ = os.Remove("transitions.txt")
		// creamos un archivo donde se guardar las transiciones
		transitionFile, err := os.Create("transitions.txt")

		if err != nil {
			return
		}

		defer transitionFile.Close()

		turingAnimate(cursorTape, &tape, transitionFile)
		dialog.ShowInformation("Se ha terminado", "la maquina se detuvo", Win)
	})

	Win = App.NewWindow("Turing Machine")
	Win.Resize(fyne.NewSize(1200, 180))

	// semilla randmon
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)

	btnRandom := widget.NewButton("Random", func() {
		maxLength := r.Intn(MaxRand/2) * 2 // recordar que rand da n-1
		randEntry := ""
		for i := 0; i < maxLength; i++ {
			if r.Intn(2) == 1 {
				randEntry += "1"
			} else {
				randEntry += "0"
			}
		}
		_ = inputString.Set(randEntry)
	})

	showDiagramBtn := widget.NewButton("Show Diagram", OpenTuringDiagram)

	menuCont := container.NewVBox(
		inputWordLbl,
		inputWord,
		animateBtn,
		slideDuration,
		strLenCont,
		btnRandom,
		showDiagramBtn,
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

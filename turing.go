package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"time"
)

func turingAnimate(cursor *CursorTape, tapeRef *[]*TapeCell) {
	tape := *tapeRef
	cell := tape[cursor.Index]
	// si estoy en q0 y recibo un cero, lo mando a q1 y muevo a la derecha, e imprimo X
	if cursor.State == "q0" && cell.Symbol == "0" {
		cursor.SetState("q1")
		cell.SetSymbol("X")
		// movemos el cursor a la derecha
		cursor.Index++
	} else if cursor.State == "q1" && cell.Symbol == "0" {
		cursor.SetState("q1")
		cell.SetSymbol("0")
		cursor.Index++
	} else if cursor.State == "q1" && cell.Symbol == "1" {
		cursor.SetState("q2")
		cell.SetSymbol("Y")
		cursor.Index--
	} else if cursor.State == "q2" && cell.Symbol == "0" {
		cursor.SetState("q2")
		cell.SetSymbol("0")
		cursor.Index--
	} else if cursor.State == "q2" && cell.Symbol == "X" {
		cursor.SetState("q0")
		cell.SetSymbol("X")
		cursor.Index++
	} else if cursor.State == "q0" && cell.Symbol == "Y" {
		cursor.SetState("q3")
		cell.SetSymbol("Y")
		cursor.Index++
	} else if cursor.State == "q1" && cell.Symbol == "Y" {
		cursor.SetState("q1")
		cell.SetSymbol("Y")
		cursor.Index++
	} else if cursor.State == "q2" && cell.Symbol == "Y" {
		cursor.SetState("q2")
		cell.SetSymbol("Y")
		cursor.Index--
	} else if cursor.State == "q3" && cell.Symbol == "Y" {
		cursor.SetState("q3")
		cell.SetSymbol("Y")
		cursor.Index++
	}

	if cursor.Index > len(tape)-1 {
		if cursor.State == "q3" {
			dialog.ShowInformation("Terminado", "La cadena esta balanciada", Win)
		}
		return // caso base
	}

	duration, _ := slideDurationBind.Get()

	timeDuration := time.Duration(duration) * (time.Millisecond * 100)

	cell = tape[cursor.Index]
	// movemos nuestro cursor a la celda
	newXAxis := cell.XAxis - cursorOffset
	moveCell := canvas.NewPositionAnimation(
		fyne.NewPos(cursor.XAxis, cursorTop),
		fyne.NewPos(newXAxis, cursorTop),
		timeDuration,
		cursor.Widget.Move,
	)
	cursor.XAxis = newXAxis
	moveCell.Start()
	time.Sleep(timeDuration) // esperamos la animacion
	tape = nil               // para que recoja el recolector
	turingAnimate(cursor, tapeRef)
}

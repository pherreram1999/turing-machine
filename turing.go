package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"os"
	"time"
)

func turingAnimate(cursor *CursorTape, tapeRef *[]*TapeCell, transitionFile *os.File) {
	tape := *tapeRef
	cell := tape[cursor.Index]
	direction := ""
	originState := cursor.State
	originSymbol := cell.Symbol
	// si estoy en q0 y recibo un cero, lo mando a q1 y muevo a la derecha, e imprimo X
	if cursor.State == "q0" && cell.Symbol == "0" {
		cursor.SetState("q1")
		cell.SetSymbol("X")
		// movemos el cursor a la derecha
		direction = "R"
		cursor.Index++
	} else if cursor.State == "q1" && cell.Symbol == "0" {
		cursor.SetState("q1")
		cell.SetSymbol("0")
		cursor.Index++
		direction = "R"
	} else if cursor.State == "q1" && cell.Symbol == "1" {
		cursor.SetState("q2")
		cell.SetSymbol("Y")
		direction = "L"
		cursor.Index--
	} else if cursor.State == "q2" && cell.Symbol == "0" {
		cursor.SetState("q2")
		cell.SetSymbol("0")
		direction = "L"
		cursor.Index--
	} else if cursor.State == "q2" && cell.Symbol == "X" {
		cursor.SetState("q0")
		cell.SetSymbol("X")
		cursor.Index++
		direction = "R"
	} else if cursor.State == "q0" && cell.Symbol == "Y" {
		cursor.SetState("q3")
		cell.SetSymbol("Y")
		cursor.Index++
		direction = "R"
	} else if cursor.State == "q1" && cell.Symbol == "Y" {
		cursor.SetState("q1")
		cell.SetSymbol("Y")
		cursor.Index++
		direction = "R"
	} else if cursor.State == "q2" && cell.Symbol == "Y" {
		cursor.SetState("q2")
		cell.SetSymbol("Y")
		cursor.Index--
		direction = "L"
	} else if cursor.State == "q3" && cell.Symbol == "Y" {
		cursor.SetState("q3")
		cell.SetSymbol("Y")
		cursor.Index++
		direction = "R"
	} else if cursor.State == "q3" && cell.Symbol == "1" {
		// DETENEMOS la maquina dado este estado no esta contemplado
		return
	}

	if cursor.Index+1 > len(tape) {
		// si no hay elementos en la cinta y no encontramos en el estaod q3
		// indica que hemos llegado al finla de la maquina
		if cursor.State == "q3" {
			dialog.ShowInformation("Terminado", "La cadena esta balanciada", Win)
		}
		return // caso base
	}

	duration, _ := slideDurationBind.Get()
	cell = tape[cursor.Index]

	if cursor.StrLen <= 10 {
		timeDuration := time.Duration(duration) * (time.Millisecond * 100)
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

	}
	// es una copia del apuntador
	tape = nil // para que recoja el recolector
	// transition
	transition := fmt.Sprintf(
		"&(%s,%s) = (%s,%s,%s)\n",
		originState,
		originSymbol,
		cursor.State,
		cell.Symbol,
		direction,
	)

	_, _ = fmt.Fprint(transitionFile, transition)

	turingAnimate(cursor, tapeRef, transitionFile)
}

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type (
	TapeCell struct {
		BoxContainer *fyne.Container
		SymbolBind   binding.String
		Symbol       string
		XAxis        float32
	}

	CursorTape struct {
		Widget    *fyne.Container
		XAxis     float32
		StateBind binding.String
		State     string
		Index     int // lleva registro donde se ubica
	}
)

func (c *CursorTape) SetState(state string) {
	c.State = state
	_ = c.StateBind.Set(state)
}

func (tc *TapeCell) SetSymbol(symbol string) {
	tc.Symbol = symbol
	_ = tc.SymbolBind.Set(symbol)
}

func (ct *CursorTape) Reset() {
	ct.XAxis = 0
	ct.Index = 0
	ct.SetState("q0")
}

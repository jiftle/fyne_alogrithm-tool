package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var w fyne.Window

func main() {
	a := app.New()

	w = a.NewWindow("algorithm tool by jiftle 2022")

	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	gui := newGUI()

	w.SetContent(gui.makeUI())
	w.ShowAndRun()

}

package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var w fyne.Window

func init() {
	// linux 指定支持中文的字体
	// os.Setenv("FYNE_FONT", "/usr/share/fonts/truetype/droid/DroidSansFallbackFull.ttf")
	os.Setenv("FYNE_THEME", "light")
}

func main() {
	a := app.New()

	w = a.NewWindow("algorithm tool by jiftle 2022")

	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	gui := newGUI()

	w.SetContent(gui.makeUI())
	w.ShowAndRun()

}

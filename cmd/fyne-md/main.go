package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/zivlakmilos/fyne-md/private/gui"
)

func main() {
	app := app.New()

	win := gui.NewMainWindow(app)
	win.Show()

	app.Run()
}

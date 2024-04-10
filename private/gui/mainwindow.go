package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	app fyne.App
	win fyne.Window

	editWidget    *widget.Entry
	previewWidget *widget.RichText
	currentFyle   fyne.URI
	saveMenuItem  *fyne.MenuItem
}

func NewMainWindow(app fyne.App) *MainWindow {
	w := &MainWindow{
		app: app,
		win: app.NewWindow("Markdown"),
	}

	w.setupUi()
	w.setupHandlers()

	return w
}

func (w *MainWindow) Show() {
	w.win.Resize(fyne.NewSize(800, 500))
	w.win.CenterOnScreen()
	w.win.Show()
}

func (w *MainWindow) setupUi() {
	w.editWidget = widget.NewMultiLineEntry()
	w.previewWidget = widget.NewRichTextFromMarkdown("")

	w.win.SetContent(container.NewHSplit(w.editWidget, w.previewWidget))
}

func (w *MainWindow) setupHandlers() {
	w.editWidget.OnChanged = w.previewWidget.ParseMarkdown
}

package gui

import (
	"io"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func NewMainWindow(app fyne.App) *MainWindow {
	w := &MainWindow{
		app: app,
		win: app.NewWindow("Markdown"),
	}

	w.setupUi()
	w.setupHandlers()
	w.createMenuItems()

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

func (w *MainWindow) createMenuItems() {
	openMenu := fyne.NewMenuItem("Open...", w.openFunc())
	saveMenu := fyne.NewMenuItem("Save", w.saveFunc())
	saveAsMenu := fyne.NewMenuItem("Save as...", w.saveAsFunc())
	fileMenu := fyne.NewMenu("File", openMenu, saveMenu, saveAsMenu)

	w.saveMenuItem = saveMenu
	w.saveMenuItem.Disabled = true

	menu := fyne.NewMainMenu(fileMenu)
	w.win.SetMainMenu(menu)
}

func (w *MainWindow) setupHandlers() {
	w.editWidget.OnChanged = w.previewWidget.ParseMarkdown
}

func (w *MainWindow) openFunc() func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w.win)
				return
			}

			if read == nil {
				return
			}

			defer read.Close()

			data, err := io.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, w.win)
				return
			}

			w.editWidget.SetText(string(data))
			w.currentFyle = read.URI()
			w.win.SetTitle(w.win.Title() + " - " + read.URI().Name())
			w.saveMenuItem.Disabled = false
		}, w.win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (w *MainWindow) saveFunc() func() {
	return func() {
		if w.currentFyle != nil {
			write, err := storage.Writer(w.currentFyle)
			if err != nil {
				dialog.ShowError(err, w.win)
				return
			}

			write.Write([]byte(w.editWidget.Text))
			defer write.Close()
		}
	}
}

func (w *MainWindow) saveAsFunc() func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w.win)
				return
			}

			if write == nil {
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with .md extension!", w.win)
				return
			}

			write.Write([]byte(w.editWidget.Text))
			w.currentFyle = write.URI()

			defer write.Close()

			w.win.SetTitle(w.win.Title() + " - " + write.URI().Name())
			w.saveMenuItem.Disabled = false
		}, w.win)

		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}

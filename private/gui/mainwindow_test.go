package gui

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestMainWindow(t *testing.T) {
	w := &MainWindow{}
	w.setupUi()
	w.setupHandlers()

	test.Type(w.editWidget, "Hello")

	if w.previewWidget.String() != "Hello" {
		t.Error("failed - did not find expected value in preview")
	}
}

func TestRunApp(t *testing.T) {
	testApp := test.NewApp()
	w := NewMainWindow(testApp)

	testApp.Run()

	test.Type(w.editWidget, "Some text")
	if w.previewWidget.String() != "Some text" {
		t.Error("failed")
	}
}

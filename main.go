package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	mApp := app.New()

	mainWindow := mApp.NewWindow("Markdown Uploader")

	pathLabel := widget.NewLabel("Obsidian Path")
	pathEntry := widget.NewEntry()
	pathEntry.SetPlaceHolder("Enter your Obsidian path...")

	githubLabel := widget.NewLabel("GitHub Repository")
	githubEntry := widget.NewEntry()
	githubEntry.SetPlaceHolder("Enter your GitHub URL...")

	grid := container.New(
		layout.NewFormLayout(),
		pathLabel,
		pathEntry,
		githubLabel,
		githubEntry,
	)

	startButton := widget.NewButton(
		"Watch Changes",
		func() {
			fmt.Println("Introduced:", githubEntry.Text)
		},
	)

	mainWindow.SetContent(
		container.NewVBox(
			grid,
			startButton,
		),
	)

	mainWindow.Show()
	mainWindow.Resize(fyne.NewSize(200, 100))
	mApp.Run()
}

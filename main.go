package main

import (
	//"fmt"
	"fmt"
	"yeel/config"
	"yeel/globals"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var app = app.New()
	var window = app.NewWindow("yeel")
	window.Resize(fyne.NewSize(800, 600))

	var isFirstTimeOpeningApplication bool 
	var _, configError = config.GetConfigurationPath(&isFirstTimeOpeningApplication)
	if configError != nil {
		fyne.CurrentApp().SendNotification(fyne.NewNotification("Error", "Failed to create configuration path!"))
	}

	if (isFirstTimeOpeningApplication) {
		firstTimeWelcomeDialog(window)
	}
	
	// top bar section
	var addWidgetBtn = widget.NewButtonWithIcon("Add Widget", theme.ContentAddIcon(), func () {})
	var linkCommandBtn = widget.NewButtonWithIcon("Command", theme.MailAttachmentIcon(), func (){})

	// here's still part of the top bar but anchored to the right
	var saveProjectButton = widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func () {});
	var runAppBtn = widget.NewButtonWithIcon("Run", theme.MediaPlayIcon(), func () {})
	var settingsButton = widget.NewButtonWithIcon("", theme.SettingsIcon(), func () {})
	var topBar = container.NewVBox(container.NewHBox(addWidgetBtn, linkCommandBtn, layout.NewSpacer(),
		saveProjectButton, runAppBtn, settingsButton),
		widget.NewSeparator()) // end of the top bar section
	
	// start of the widget panel (left side)
	// list is initialized in yeel/globals 
	var widgetPanel = widget.NewList(func () int {return len(globals.WIDGET_LIST)}, func () fyne.CanvasObject {return widget.NewLabel("")},
		func (lii widget.ListItemID, wi fyne.CanvasObject) {
			wi.(*widget.Label).SetText(globals.WIDGET_LIST[lii].Title)
		}) // end of widget panel

	var mainContainer = container.NewBorder(nil, nil, nil, nil, widget.NewLabel("Hello World!"))
	var mainSplit = container.NewHSplit(widgetPanel, mainContainer)
        mainSplit.SetOffset(0.3)

	var bottomPanel = container.NewVBox(widget.NewSeparator(), widget.NewLabel("Logs"))
	
	var mainContent = container.NewBorder(topBar, bottomPanel, nil, nil, mainSplit)
	window.SetContent(mainContent)

	// the exit sequence
        // maybe i should move it somewhere else
	window.SetCloseIntercept(func () {
		dialog.NewConfirm("Exit?", "Are you certain you want to exit? you may have unsaved work", func (b bool) {
			if b {
				fmt.Println("Bye bye :)")
                                window.Close()
			}
		}, window).Show()
	})
        window.ShowAndRun()
}


func firstTimeWelcomeDialog(parent fyne.Window) {
	dialog.NewInformation("Welcome to Yeel", "Looks like it's your first time opening Yeel, the commands should be easy intuitive to use, Enjoy!", parent).Show()
}

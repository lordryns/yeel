package main

import (
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

	// initialising widgets
	var widgetPanel *widget.List
	
	// top bar section
	var addWidgetBtn = widget.NewButtonWithIcon("Add Widget", theme.ContentAddIcon(), func () { addWidgetDialog(window, widgetPanel) })
	var linkCommandBtn = widget.NewButtonWithIcon("Command", theme.MailAttachmentIcon(), func (){})

	// here's still part of the top bar but anchored to the right
	var saveProjectButton = widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func () {});
	saveProjectButton.Importance = widget.HighImportance
	
	var runAppBtn = widget.NewButtonWithIcon("Run", theme.MediaPlayIcon(), func () {})
	var settingsButton = widget.NewButtonWithIcon("", theme.SettingsIcon(), func () {})
	var topBar = container.NewVBox(container.NewHBox(addWidgetBtn, linkCommandBtn, layout.NewSpacer(),
		saveProjectButton, runAppBtn, settingsButton),
		widget.NewSeparator()) // end of the top bar section
	
	// start of the widget panel (left side)
	// list is initialized in yeel/globals 
	widgetPanel = widget.NewList(func () int {return len(globals.ProjectSchema.Widgets)}, func () fyne.CanvasObject {return widget.NewLabel("")},
		func (lii widget.ListItemID, wi fyne.CanvasObject) {
			wi.(*widget.Label).SetText(globals.ProjectSchema.Widgets[lii].Title)
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


func addWidgetDialog(parent fyne.Window, list *widget.List) {
	var selected string
	var cdialog dialog.Dialog
	var selectWidget = widget.NewSelect(globals.AvailableWidgets, func (s string) { selected = s })
	var errorDialog = widget.NewLabel("")
	errorDialog.Importance = widget.DangerImportance
	errorDialog.Alignment = fyne.TextAlignCenter
	
	var cancelButton = widget.NewButton("Cancel", func () {cdialog.Dismiss()})
	var addButton = widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func (){
		var widgets = globals.ProjectSchema.Widgets
		var selectedEnum globals.WIDGET_TYPE
		switch (selected) {
		case "Button":
		       selectedEnum = globals.WIDGET_BUTTON
		case "Entry":
			selectedEnum = globals.WIDGET_ENTRY
		case "Checkbox":
		        selectedEnum = globals.WIDGET_CHECKBOX
		default:
		    selectedEnum = globals.WIDGET_NONE
		}
		if len(selected) > 2 {
			var widget = globals.Widget{Widget: selectedEnum, Title: fmt.Sprintf("%v%v", selected, len(widgets))}
		        globals.ProjectSchema.Widgets = append(globals.ProjectSchema.Widgets, widget)
			cdialog.Dismiss()
		        list.Refresh()
		} else {
			errorDialog.SetText("oopsie, looks like you forgot to select a valid window, try again")
		}
	})
	addButton.Importance = widget.HighImportance
	
	var content = container.NewVBox(selectWidget, container.NewBorder(nil, errorDialog, nil, container.NewHBox(cancelButton, addButton)))
	cdialog = dialog.NewCustomWithoutButtons("Create Widget", content, parent)
	cdialog.Resize(fyne.NewSize(450, cdialog.MinSize().Height))
	
	cdialog.Show()
}


func firstTimeWelcomeDialog(parent fyne.Window) {
	dialog.NewInformation("Welcome to Yeel", "Looks like it's your first time opening Yeel, the commands should be easy intuitive to use, Enjoy!", parent).Show()
}

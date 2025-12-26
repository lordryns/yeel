package globals

type WIDGET_TYPE int

const (
	WIDGET_BUTTON WIDGET_TYPE = iota
	WIDGET_ENTRY
	WIDGET_CHECKBOX
	WIDGET_FRAME
	WIDGET_NONE
)

type Widget struct {
	Widget WIDGET_TYPE
	Title string
	RelX int
	RelY int
	RelWidth int
	RelHeight int
}

type Command struct {

}

type PROJECT_SCHEMA struct {
	Widgets []Widget
	Commands []Command
}

var ProjectSchema PROJECT_SCHEMA 


var AvailableWidgets = []string{"Button", "Entry", "Checkbox"}

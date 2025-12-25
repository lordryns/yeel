package globals

type WIDGET_TYPE int

const (
	WIDGET_BUTTON WIDGET_TYPE = iota
	WIDGET_ENTRY
	WIDGET_CHECKBOX
	WIDGET_FRAME
)

type Widget struct {
	Widget WIDGET_TYPE
	Title string
}

var WIDGET_LIST []Widget

package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateWarningContent(window fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle("Warning", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Создатель сурса не отвечает за ваши действия. Действуйте своей головой!"),
	)
}

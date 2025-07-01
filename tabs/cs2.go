package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateCs2Content(window fyne.Window) fyne.CanvasObject {
	cs2Card := widget.NewCard(
		"Counter-Strike 2 Loader", // Полное название
		"Инжектор для CS2",        // Краткое описание
		container.NewVBox(
			widget.NewLabel("Этот раздел предназначен для инжекта читов в CS2. Используйте на свой страх и риск."),
			canvas.NewImageFromResource(theme.ComputerIcon()),
			widget.NewLabel("Информация"),
			widget.NewLabel("Подписка: 19.01.2038"),
			widget.NewLabel("Последнее обновление: 21.01.2024"),
			widget.NewLabel("Статус: Stable"),
			layout.NewSpacer(),
			widget.NewButton("Inject", func() {
				dialog.ShowInformation("CS2", "Инжект выполнен!", window)
			}),
		),
	)
	return container.NewVBox(
		widget.NewLabelWithStyle("Counter-Strike 2 Loader", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		cs2Card,
	)
}

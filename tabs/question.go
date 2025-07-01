package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var questions = []struct {
	Q string
	A string
}{
	{"Как пользоваться Loader?", "Выберите нужную вкладку и следуйте инструкциям."},
	{"Где взять скрипты?", "В разделе Lua есть дефолтные скрипты, а также поиск."},
	{"Что делать при ошибке Inject?", "Перезапустите игру и попробуйте снова."},
}

func CreateQuestionContent(window fyne.Window) fyne.CanvasObject {
	list := widget.NewList(
		func() int { return len(questions) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(questions[i].Q)
		},
	)
	answer := widget.NewMultiLineEntry()
	answer.SetPlaceHolder("Ответ...")
	list.OnSelected = func(id int) {
		answer.SetText(questions[id].A)
	}
	return container.NewVBox(
		widget.NewLabelWithStyle("Вопросы и ответы", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		container.NewHSplit(list, answer),
	)
}

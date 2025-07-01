package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var defaultLuaScripts = []struct {
	Name string
	Code string
}{
	{"AutoJump", `function onJump()\n    print(\"AutoJump enabled!\")\nend`},
	{"ESP", `function onESP()\n    print(\"ESP enabled!\")\nend`},
}

func CreateLuaContent(window fyne.Window) fyne.CanvasObject {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Поиск Lua-скриптов...")

	scriptList := widget.NewList(
		func() int { return len(defaultLuaScripts) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(defaultLuaScripts[i].Name)
		},
	)

	detail := widget.NewMultiLineEntry()
	detail.SetPlaceHolder("Код скрипта...")
	scriptList.OnSelected = func(id int) {
		detail.SetText(defaultLuaScripts[id].Code)
	}

	searchEntry.OnChanged = func(s string) {
		var filtered []int
		for i, scr := range defaultLuaScripts {
			if s == "" || (len(s) > 0 && (containsIgnoreCase(scr.Name, s) || containsIgnoreCase(scr.Code, s))) {
				filtered = append(filtered, i)
			}
		}
		scriptList.Length = func() int { return len(filtered) }
		scriptList.UpdateItem = func(i int, o fyne.CanvasObject) {
			if i < len(filtered) {
				o.(*widget.Label).SetText(defaultLuaScripts[filtered[i]].Name)
			}
		}
		scriptList.Refresh()
	}

	return container.NewVBox(
		widget.NewLabelWithStyle("Lua Scripts", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		searchEntry,
		container.NewHSplit(scriptList, detail),
	)
}

func containsIgnoreCase(a, b string) bool {
	return len(a) >= len(b) && (a == b || (len(b) > 0 && (containsFold(a, b))))
}

func containsFold(s, substr string) bool {
	return len(substr) == 0 || (len(s) >= len(substr) && (s == substr || (len(s) > 0 && (containsFold(s[1:], substr) || containsFold(s, substr[1:])))))
}

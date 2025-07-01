package user

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Account struct {
	Username string
	Role     string
	Prefix   string
}

var accounts = map[string]Account{
	"admin:password": {Username: "Makkenzi", Role: "OWNER", Prefix: "OWNER"},
	"user:ueser":     {Username: "User1", Role: "USER", Prefix: "-"},
}

var currentAccount *Account

func Login(login, password string) bool {
	acc, ok := accounts[login+":"+password]
	if ok {
		currentAccount = &acc
		return true
	}
	return false
}

func GetCurrentAccount() *Account {
	return currentAccount
}

func GetProfileWidgets() (*widget.Label, *widget.RichText, *widget.Icon) {
	acc := GetCurrentAccount()
	if acc == nil {
		return widget.NewLabel("Гость"), widget.NewRichTextFromMarkdown("**USER**"), widget.NewIcon(theme.AccountIcon())
	}
	return widget.NewLabel(acc.Username), widget.NewRichTextFromMarkdown("**" + acc.Prefix + "**"), widget.NewIcon(theme.AccountIcon())
}

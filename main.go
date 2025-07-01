package main

import (
	"time"

	"awesomeProject/tabs"
	"awesomeProject/user"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CollapsibleNavButton struct {
	widget.BaseWidget
	icon      fyne.Resource
	text      string
	label     *widget.Label
	iconImage *widget.Icon
	content   *fyne.Container // Container to hold icon and label
	onTapped  func()          // Store the tapped function
}

func NewCollapsibleNavButton(text string, icon fyne.Resource, tapped func()) *CollapsibleNavButton {
	label := widget.NewLabel(text)
	iconImage := widget.NewIcon(icon)

	content := container.NewHBox(iconImage, label)
	content.Layout = layout.NewHBoxLayout()

	btn := &CollapsibleNavButton{
		icon:      icon,
		text:      text,
		label:     label,
		iconImage: iconImage,
		content:   content,
		onTapped:  tapped, // Store the callback
	}
	btn.ExtendBaseWidget(btn) // Initialize the BaseWidget
	return btn
}

func (c *CollapsibleNavButton) CreateRenderer() fyne.WidgetRenderer {
	// For a button-like appearance, use a background rectangle and then your content.
	// Using BackgroundColor instead of deprecated ButtonColor
	rect := canvas.NewRectangle(theme.BackgroundColor())

	objects := []fyne.CanvasObject{rect, c.content} // Order matters: background first, then content

	return &collapsibleNavButtonRenderer{
		collapsibleNavButton: c,
		background:           rect,
		objects:              objects,
	}
}

func (c *CollapsibleNavButton) Tapped(event *fyne.PointEvent) {
	if c.onTapped != nil {
		c.onTapped()
	}
}

type collapsibleNavButtonRenderer struct {
	collapsibleNavButton *CollapsibleNavButton
	background           *canvas.Rectangle
	objects              []fyne.CanvasObject
}

func (r *collapsibleNavButtonRenderer) MinSize() fyne.Size {
	return r.collapsibleNavButton.content.MinSize().Add(fyne.NewSize(theme.Padding()*2, theme.Padding()*2))
}

func (r *collapsibleNavButtonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.collapsibleNavButton.content.Move(fyne.NewPos(theme.Padding(), theme.Padding()))
	r.collapsibleNavButton.content.Resize(size.Subtract(fyne.NewSize(theme.Padding()*2, theme.Padding()*2)))
}

func (r *collapsibleNavButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *collapsibleNavButtonRenderer) Destroy() {
}

func (r *collapsibleNavButtonRenderer) Refresh() {
	r.collapsibleNavButton.label.Refresh()
	r.collapsibleNavButton.iconImage.Refresh()
	r.collapsibleNavButton.content.Refresh()
	r.background.FillColor = theme.BackgroundColor()
	canvas.Refresh(r.background)
}

func (c *CollapsibleNavButton) SetCollapsed(collapsed bool) {
	if collapsed {
		c.label.Hide()
	} else {
		c.label.Show()
	}
	c.content.Refresh()
	c.Refresh()
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Личный кабинет")

	// --- Login Screen ---
	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Логин")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Пароль")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()
	errorLabel.TextStyle = fyne.TextStyle{Bold: true}

	loginForm := widget.NewForm(
		widget.NewFormItem("Логин", loginEntry),
		widget.NewFormItem("Пароль", passwordEntry),
	)

	loginButton := widget.NewButton("Войти", func() {
		if user.Login(loginEntry.Text, passwordEntry.Text) {
			showPersonalCabinet(myWindow)
		} else {
			errorLabel.SetText("Неверный логин или пароль")
			errorLabel.Show()
			shakeAnimation(errorLabel)
		}
	})

	exitButton := widget.NewButton("Выход", func() {
		myApp.Quit()
	})

	loginContent := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(widget.NewLabelWithStyle("Авторизация", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true})),
		loginForm,
		container.NewCenter(loginButton),
		container.NewCenter(errorLabel),
		container.NewCenter(exitButton),
		layout.NewSpacer(),
	)

	myWindow.SetContent(container.NewCenter(loginContent))
	myWindow.Resize(fyne.NewSize(400, 350))
	myWindow.ShowAndRun()
}

// showPersonalCabinet - displays the personal cabinet interface
func showPersonalCabinet(window fyne.Window) {
	// Main content area that will change dynamically
	mainContentArea := container.NewStack()
	mainContentArea.Add(createWelcomeScreen())

	var navBar *fyne.Container // Declare for easier access
	var split *container.Split
	isNavCollapsed := true // Initial state: navigation is collapsed

	// Define a fixed width for the collapsed (icon-only) state
	collapsedNavWidth := theme.IconInlineSize() + theme.Padding()*4 // Call functions to get float32 values

	// Navigation buttons
	navButtonCs2 := NewCollapsibleNavButton("Cs2", theme.ComputerIcon(), func() {
		mainContentArea.Objects = []fyne.CanvasObject{tabs.CreateCs2Content(window)}
		mainContentArea.Refresh()
		if !isNavCollapsed {
			isNavCollapsed = !isNavCollapsed
			animateSidebar(split, navBar, collapsedNavWidth, isNavCollapsed)
		}
	})

	navButtonLua := NewCollapsibleNavButton("Lua", theme.DocumentIcon(), func() {
		mainContentArea.Objects = []fyne.CanvasObject{tabs.CreateLuaContent(window)}
		mainContentArea.Refresh()
		if !isNavCollapsed {
			isNavCollapsed = !isNavCollapsed
			animateSidebar(split, navBar, collapsedNavWidth, isNavCollapsed)
		}
	})

	navButtonWarning := NewCollapsibleNavButton("Warning", theme.WarningIcon(), func() {
		mainContentArea.Objects = []fyne.CanvasObject{tabs.CreateWarningContent(window)}
		mainContentArea.Refresh()
		if !isNavCollapsed {
			isNavCollapsed = !isNavCollapsed
			animateSidebar(split, navBar, collapsedNavWidth, isNavCollapsed)
		}
	})

	navButtonQuestion := NewCollapsibleNavButton("Question", theme.QuestionIcon(), func() {
		mainContentArea.Objects = []fyne.CanvasObject{tabs.CreateQuestionContent(window)}
		mainContentArea.Refresh()
		if !isNavCollapsed {
			isNavCollapsed = !isNavCollapsed
			animateSidebar(split, navBar, collapsedNavWidth, isNavCollapsed)
		}
	})

	// User profile section at the bottom of the navigation panel
	userProfileLabel, userProfileRole, userProfileIcon := user.GetProfileWidgets()
	userProfileContent := container.NewHBox(
		userProfileIcon,
		userProfileLabel,
		userProfileRole,
	)

	// Top part of the sidebar with app name/logo and a collapse button
	appLogo := widget.NewLabelWithStyle("Loader", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	collapseButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() { // Changed to NavigateBackIcon
		isNavCollapsed = !isNavCollapsed
		animateSidebar(split, navBar, collapsedNavWidth, isNavCollapsed)
	})

	navBarHeader := container.NewHBox(
		appLogo,
		layout.NewSpacer(),
		collapseButton,
	)

	// Main navigation panel content
	navBar = container.NewVBox(
		navBarHeader,
		widget.NewSeparator(),
		container.NewVBox(
			navButtonCs2,
			navButtonLua,
			navButtonWarning,
			navButtonQuestion,
		),
		layout.NewSpacer(),
		widget.NewSeparator(),
		userProfileContent,
	)

	// Set initial collapsed state for all collapsible elements
	// It's important to access the actual CollapsibleNavButton objects
	navButtonCs2.SetCollapsed(isNavCollapsed)
	navButtonLua.SetCollapsed(isNavCollapsed)
	navButtonWarning.SetCollapsed(isNavCollapsed)
	navButtonQuestion.SetCollapsed(isNavCollapsed)

	appLogo.Hide()                                   // Hide logo when collapsed
	userProfileContent.Hide()                        // Hide full profile when collapsed
	collapseButton.SetIcon(theme.NavigateNextIcon()) // Initial icon for expand (changed to NavigateNextIcon)

	// Create a wrapper for the navBar to control its width
	navBarWrapper := container.New(layout.NewMaxLayout(), navBar)

	split = container.NewHSplit(navBarWrapper, mainContentArea)
	split.SetOffset(0.0) // Start with the sidebar collapsed

	// Top bar with menu button (now functions as an expander for the sidebar)
	topBar := container.NewHBox(
		widget.NewButtonWithIcon("", theme.MenuIcon(), func() {
			isNavCollapsed = !isNavCollapsed
			animateSidebar(split, navBar, collapsedNavWidth, isNavCollapsed)
		}),
		layout.NewSpacer(),
		widget.NewLabel("Личный кабинет"), // Title of the current section
		layout.NewSpacer(),
		widget.NewButton("Выйти", func() {
			// Logout functionality, simplified to return to login screen
			window.SetContent(container.NewCenter(createLoginContent(window)))
			window.Resize(fyne.NewSize(400, 350)) // Resize back to login screen size
		}),
	)

	cabinetLayout := container.NewBorder(topBar, nil, nil, nil, split)

	window.SetContent(cabinetLayout)
	window.Resize(fyne.NewSize(800, 600)) // Increase window size for better layout
	window.Show()
}

// animateSidebar handles the animation and state change of the sidebar
func animateSidebar(split *container.Split, navBar *fyne.Container, collapsedWidth float32, isExpanding bool) {
	startOffset := split.Offset
	endOffset := 0.0       // Target offset for collapsed state
	expandedOffset := 0.25 // Target offset for expanded state

	if isExpanding { // If we are expanding
		endOffset = expandedOffset
	}

	animation := fyne.NewAnimation(time.Millisecond*250, func(done float32) {
		// All UI changes within the animation are already in the correct thread
		currentOffset := startOffset + (endOffset-startOffset)*float64(done)
		split.SetOffset(currentOffset)
	})

	animation.Curve = fyne.AnimationEaseInOut // Smooth animation
	animation.Start()

	// Toggle visibility of elements after animation completes or based on a threshold
	// All UI modification calls inside this goroutine must be wrapped in fyne.Do()
	go func() {
		time.Sleep(time.Millisecond * 250) // Wait for animation to finish

		// Wrap all UI changes in fyne.Do()
		fyne.Do(func() { // Using fyne.Do to run this code on the main UI thread
			// Update visibility of elements based on the final state (isExpanding)
			if isExpanding { // Panel is now expanded
				// Show app logo and full user profile
				if navBar.Objects[0].(*fyne.Container).Objects[0] != nil {
					navBar.Objects[0].(*fyne.Container).Objects[0].Show() // K°oax label
				}
				if navBar.Objects[5] != nil {
					navBar.Objects[5].Show() // User profile content
				}
				// Update collapse button icon to "collapse" (arrow left)
				if navBar.Objects[0].(*fyne.Container).Objects[2] != nil {
					navBar.Objects[0].(*fyne.Container).Objects[2].(*widget.Button).SetIcon(theme.NavigateBackIcon())
				}

				// Show text for all collapsible buttons
				// Iterate over the objects in the navButtonsContainer and cast them
				for _, obj := range navBar.Objects[2].(*fyne.Container).Objects {
					if btn, ok := obj.(*CollapsibleNavButton); ok { // Corrected type assertion
						btn.SetCollapsed(false)
					}
				}
				split.SetOffset(expandedOffset) // Ensure it snaps to the expanded position
			} else { // Panel is now collapsed
				// Hide app logo and full user profile
				if navBar.Objects[0].(*fyne.Container).Objects[0] != nil {
					navBar.Objects[0].(*fyne.Container).Objects[0].Hide() // K°oax label
				}
				if navBar.Objects[5] != nil {
					navBar.Objects[5].Hide() // User profile content
				}
				// Update collapse button icon to "expand" (arrow right)
				if navBar.Objects[0].(*fyne.Container).Objects[2] != nil {
					navBar.Objects[0].(*fyne.Container).Objects[2].(*widget.Button).SetIcon(theme.NavigateNextIcon())
				}

				// Hide text for all collapsible buttons
				// Iterate over the objects in the navButtonsContainer and cast them
				for _, obj := range navBar.Objects[2].(*fyne.Container).Objects {
					if btn, ok := obj.(*CollapsibleNavButton); ok { // Corrected type assertion
						btn.SetCollapsed(true)
					}
				}
				split.SetOffset(0.0) // Ensure it snaps to the collapsed position
			}
			navBar.Refresh() // This also needs to be in the main Fyne thread
		}) // End of fyne.Do() block
	}()
}

// createLoginContent creates the login screen content for re-use
func createLoginContent(window fyne.Window) fyne.CanvasObject {
	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Логин")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Пароль")
	errorLabel := widget.NewLabel("")
	errorLabel.Hide()
	errorLabel.TextStyle = fyne.TextStyle{Bold: true}

	loginForm := widget.NewForm(
		widget.NewFormItem("Логин", loginEntry),
		widget.NewFormItem("Пароль", passwordEntry),
	)
	loginBtn := widget.NewButton("Войти", func() {
		if user.Login(loginEntry.Text, passwordEntry.Text) {
			showPersonalCabinet(window)
		} else {
			errorLabel.SetText("Неверный логин или пароль")
			errorLabel.Show()
			shakeAnimation(errorLabel)
		}
	})

	exitAppBtn := widget.NewButton("Выход", func() {
		fyne.CurrentApp().Quit()
	})

	return container.NewCenter(
		container.NewVBox(
			layout.NewSpacer(),
			container.NewCenter(widget.NewLabelWithStyle("Авторизация", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true})),
			loginForm,
			container.NewCenter(loginBtn),
			container.NewCenter(errorLabel),
			container.NewCenter(exitAppBtn),
			layout.NewSpacer(),
		),
	)
}

// createWelcomeScreen creates the initial content for the main area
func createWelcomeScreen() fyne.CanvasObject {
	return container.NewCenter(
		widget.NewLabelWithStyle("Выберите пункт меню слева", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true}),
	)
}

// createSkriptContent creates the content for the "Skript" section
func createSkriptContent(window fyne.Window) fyne.CanvasObject {
	// Simulate the Skript - Client section
	skriptClientCard := widget.NewCard(
		"Skript - Client", // Title
		"",                // Subtitle
		container.NewVBox(
			widget.NewLabel("Clicking 'Start' will load the cheat, after which the button will change to 'Destruct'. Clicking 'Destruct' will then initiate the cleaning process."),
			canvas.NewImageFromResource(theme.ComputerIcon()), // Using canvas.NewImageFromResource for icons
			widget.NewLabel("Informations"),
			widget.NewLabel("Subscription: 19.01.2038"),
			widget.NewLabel("Last update: 21.01.2024"),
			widget.NewLabel("Status: Stable"),
			layout.NewSpacer(), // Pushes button to bottom
			widget.NewButton("Start", func() {
				dialog.ShowInformation("Skript", "Настройки Skript применены!", window)
				// Here you would change the button to "Destruct" and handle logic
			}),
		),
	)

	return container.NewVBox(
		widget.NewLabelWithStyle("Skript - Client", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		skriptClientCard,
	)
}

// shakeAnimation adds a simple shake effect to a widget
func shakeAnimation(obj fyne.CanvasObject) {
	originalPos := obj.Position()
	shakeIntensity := float32(5)
	shakeDuration := time.Millisecond * 50

	animation := fyne.NewAnimation(shakeDuration*5, func(done float32) {
		offset := fyne.NewPos(0, 0)
		if int(done*10)%2 == 0 { // Alternate left/right for shaking effect
			offset.X = shakeIntensity
		} else {
			offset.X = -shakeIntensity
		}
		obj.Move(originalPos.Add(offset))
	})

	animation.Curve = fyne.AnimationEaseOut
	animation.Start()

	// Return to original position after animation
	// This also needs to be on the UI thread, so wrap in fyne.Do()
	time.AfterFunc(shakeDuration*5, func() {
		fyne.Do(func() { // Wrap in fyne.Do to ensure UI update on main thread
			obj.Move(originalPos)
		})
	})
}

package gui

import (
	"homework/probability/types"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Gui struct {
	service types.ProbabilityService
}

func NewGui(service types.ProbabilityService) *Gui {
	return &Gui{
		service: service,
	}
}

func (g *Gui) Run() {
	a := app.New()
	w := a.NewWindow("Main")
	w.Resize(fyne.NewSize(800, 600))

	title := widget.NewLabel("Дискретная математика")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	combo := widget.NewSelect(types.Themes, func(value string) {
		themeWindow := a.NewWindow(value)
		themeWindow.Resize(fyne.NewSize(600, 400))

		themeData := types.ThemeSwitcher(value)
		theoryLabel := widget.NewLabel(themeData.Theory)
		theoryLabel.Wrapping = fyne.TextWrapWord

		themeContent := container.NewVBox(
			widget.NewLabel("Теория:"),
			theoryLabel,
		)
		themeContent = container.NewBorder(nil, nil, nil, nil, themeContent)

		themeWindow.SetContent(themeContent)
		themeWindow.Show()
	})

	combo.SetSelected("Выберите нужную тему")

	centered := container.NewCenter(combo)
	content := container.NewVBox(
		title,
		layout.NewSpacer(),
		centered,
		layout.NewSpacer(),
	)


	w.SetContent(content)
	w.ShowAndRun()
}
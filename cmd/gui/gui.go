package gui

import (
	"homework/probability/themes"
	"homework/probability/types"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Gui struct {
	service types.MainService
}

func NewGui(service types.MainService) *Gui {
	return &Gui{
		service: service,
	}
}

func (g *Gui) Run() {
	a := app.New()
	a.Settings().SetTheme(themes.NewCustomTheme())

	w := a.NewWindow("Дискретная математика")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()

	title := canvas.NewText("Дискретная математика", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter

	subtitle := widget.NewLabel("Добро пожаловать! Это справочник по ключевым темам дискретной математики.")
	subtitle.Wrapping = fyne.TextWrapWord
	subtitle.Alignment = fyne.TextAlignCenter

	description := widget.NewLabel("Выберите тему ниже, чтобы изучить теорию, формулы, примеры и решения.")
	description.Wrapping = fyne.TextWrapWord
	description.Alignment = fyne.TextAlignCenter

	themeSelector := widget.NewSelect(types.Themes, func(value string) {
		if value == "Решатор задач" {
			reshator(a)
			return
		}

		data := types.ThemeSwitcher(value)
		detail := a.NewWindow(value)
		detail.Resize(fyne.NewSize(700, 500))

		section := func(label string, content fyne.CanvasObject) fyne.CanvasObject {
			header := canvas.NewText(label, color.White)
			header.TextStyle = fyne.TextStyle{Bold: true}
			header.TextSize = 14

			box := container.NewVBox(header, content)
			bg := canvas.NewRectangle(color.RGBA{45, 45, 45, 255})
			bg.CornerRadius = 12
			return container.NewStack(bg, container.NewPadded(box))
		}

		theory := section("Теория", widget.NewLabelWithData(binding.BindString(&data.Theory)))

		formula := section("Формула", container.NewVBox(
			func() fyne.CanvasObject {
				if data.FormulaPath == "" {
					return widget.NewLabel("")
				}
				img := canvas.NewImageFromFile(data.FormulaPath)
				img.FillMode = canvas.ImageFillContain
				img.SetMinSize(fyne.NewSize(300, 150))
				return img
			}(),
			widget.NewLabel(data.FormulaDescription),
		))
		
		taskContent := container.NewVBox()
		if data.TaskImage != "" {
			taskImg := canvas.NewImageFromFile(data.TaskImage)
			taskImg.FillMode = canvas.ImageFillContain
			taskImg.SetMinSize(fyne.NewSize(400, 250))
			taskContent.Add(container.NewCenter(taskImg))
		}
		taskContent.Add(widget.NewLabel(data.ExampleText))
		
		example := section("Пример", taskContent)
		
		buttons := container.NewHBox(
			widget.NewButton("Решение", func() {
				ShowSolutionWindow(a, "Решение", data.SolutionText, data.Hint)
			}),
			widget.NewButton("Подсказка", func() {
				ShowWindow(a, "Подсказка", data.Hint)
			}),
		)

		content := container.NewVBox(theory, formula, example, buttons)
		scroll := container.NewVScroll(content)
		detail.SetContent(scroll)
		detail.Show()
	})

	themeSelector.PlaceHolder = "Выберите тему"
	form := container.NewVBox(
		layout.NewSpacer(),
		title,
		layout.NewSpacer(),
		subtitle,
		description,
		layout.NewSpacer(),
		container.NewCenter(themeSelector),
		layout.NewSpacer(),
	)

	bg := canvas.NewRectangle(color.RGBA{R: 36, G: 36, B: 36, A: 255})
	formBox := container.NewStack(bg, container.NewPadded(form))

	w.SetContent(formBox)
	w.ShowAndRun()
}

func ShowWindow(a fyne.App, title, text string) {
	win := a.NewWindow(title)
	win.Resize(fyne.NewSize(400, 300))

	label := widget.NewLabel(text)
	label.Wrapping = fyne.TextWrapWord
	label.TextStyle = fyne.TextStyle{Italic: true}

	box := container.NewVBox(label)
	bg := canvas.NewRectangle(color.RGBA{45, 45, 45, 255})
	bg.CornerRadius = 12

	win.SetContent(container.NewStack(bg, container.NewPadded(box)))
	win.Show()
}

func ShowSolutionWindow(a fyne.App, title, solution, hint string) {
	win := a.NewWindow(title)
	win.Resize(fyne.NewSize(500, 400))

	answerLabel := widget.NewLabel(solution)
	answerLabel.TextStyle = fyne.TextStyle{Bold: true}
	
	explanationLabel := widget.NewLabel(hint)
	explanationLabel.Wrapping = fyne.TextWrapWord

	content := container.NewVBox(
		widget.NewLabel("Ответ:"),
		answerLabel,
		widget.NewSeparator(),
		widget.NewLabel("Пояснение:"),
		explanationLabel,
	)
	
	bg := canvas.NewRectangle(color.RGBA{45, 45, 45, 255})
	bg.CornerRadius = 12
	
	win.SetContent(container.NewStack(bg, container.NewPadded(content)))
	win.Show()
}
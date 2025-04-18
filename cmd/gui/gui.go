package gui

import (
	"homework/probability/themes"
	"homework/probability/types"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

	w := a.NewWindow("Main")
	w.Resize(fyne.NewSize(800, 600))

	title := canvas.NewText("Дискретная математика", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 15
	title.Alignment = fyne.TextAlignCenter

	welcomeText := widget.NewLabel("Данное приложение содержит справочный материал по дискретной математике. Выберите тему из списка ниже, чтобы изучить теорию, формулы и примеры задач.")
	welcomeText.Wrapping = fyne.TextWrapWord
	welcomeText.Alignment = fyne.TextAlignCenter

	combo := widget.NewSelect(types.Themes, func(value string) {
		themeWindow := a.NewWindow(value)
		themeWindow.Resize(fyne.NewSize(600, 400))

		// Теория
		themeData := types.ThemeSwitcher(value)
		theoryLabel := widget.NewLabel(themeData.Theory)
		theoryLabel.Wrapping = fyne.TextWrapWord
		theoryLabel.TextStyle = fyne.TextStyle{Italic: true}

		theoryMainText := canvas.NewText("Теория:", color.White)
		theoryMainText.TextStyle = fyne.TextStyle{Bold: true}
		theoryMainText.TextSize = 14 

		theoryBox := container.NewVBox(
			theoryMainText,
			theoryLabel,
		)
		theoryBg := canvas.NewRectangle(color.RGBA{R: 0x2D, G: 0x2D, B: 0x2D, A: 0xFF})
		theoryBg.CornerRadius = 8
		theoryBox = container.NewStack(theoryBg, container.NewPadded(theoryBox))

		// Формула
		formulaMainText := canvas.NewText("Формула:", color.White)
		formulaMainText.TextStyle = fyne.TextStyle{Bold: true}
		formulaMainText.TextSize = 14

		formulaImage := canvas.NewImageFromFile(themeData.FormulaPath)
		formulaImage.FillMode = canvas.ImageFillContain
		formulaImage.SetMinSize(fyne.NewSize(200, 100))

		formulaLabel := widget.NewLabel(themeData.FormulaDescription)
		formulaLabel.Wrapping = fyne.TextWrapWord
		formulaLabel.TextStyle = fyne.TextStyle{Italic: true}

		formulaBox := container.NewVBox(
			formulaMainText,
			formulaImage,
			formulaLabel,
		)
		formulaBg := canvas.NewRectangle(color.RGBA{R: 0x2D, G: 0x2D, B: 0x2D, A: 0xFF})
		formulaBg.CornerRadius = 8
		formulaBox = container.NewStack(formulaBg, container.NewPadded(formulaBox))

		// Пример
		exampleText := canvas.NewText("Пример:", color.White)
		exampleText.TextStyle = fyne.TextStyle{Bold: true}
		exampleText.TextSize = 14

		exampleImage := canvas.NewImageFromFile(themeData.ExampleImage)
		exampleImage.FillMode = canvas.ImageFillContain
		exampleImage.SetMinSize(fyne.NewSize(400, 300))

		exampleLabel := widget.NewLabel(themeData.ExampleText)
		exampleLabel.Wrapping = fyne.TextWrapWord
		exampleLabel.TextStyle = fyne.TextStyle{Italic: true}

		exampleBox := container.NewVBox(
			exampleText,
			exampleImage,
			exampleLabel,
		)
		exampleBg := canvas.NewRectangle(color.RGBA{R: 0x2D, G: 0x2D, B: 0x2D, A: 0xFF})
		exampleBg.CornerRadius = 8
		exampleBox = container.NewStack(exampleBg, container.NewPadded(exampleBox))

		// Решения
		solutionButton := widget.NewButton("Показать решение", func() {
			solutionWindow := a.NewWindow("Решение")
			solutionWindow.Resize(fyne.NewSize(400, 300))

			solutionText := widget.NewLabel(themeData.SolutionText)
			solutionText.Wrapping = fyne.TextWrapWord
			solutionText.TextStyle = fyne.TextStyle{Italic: true}

			solutionBox := container.NewVBox(solutionText)
			solutionBg := canvas.NewRectangle(color.RGBA{R: 0x2D, G: 0x2D, B: 0x2D, A: 0xFF})
			solutionBg.CornerRadius = 8
			solutionBox = container.NewStack(solutionBg, container.NewPadded(solutionBox))

			solutionWindow.SetContent(solutionBox)
			solutionWindow.Show()
		})

		// Подзказки
		hintBut := widget.NewButton("Показать подсказку", func() {

			hintWindow := a.NewWindow("Подсказка")
			hintWindow.Resize(fyne.NewSize(400, 300))

			hintText := widget.NewLabel(themeData.Hint)
			hintText.Wrapping = fyne.TextWrapWord
			hintText.TextStyle = fyne.TextStyle{Italic: true}

			hintBox := container.NewVBox(hintText)
			hintBg := canvas.NewRectangle(color.RGBA{R: 0x2D, G: 0x2D, B: 0x2D, A: 0xFF})
			hintBg.CornerRadius = 8
			hintBox = container.NewStack(hintBg, container.NewPadded(hintBox))

			hintWindow.SetContent(hintBox)
			hintWindow.Show()
		})

		themeContent := container.NewVBox(
			theoryBox,
			layout.NewSpacer(),
			formulaBox,
			layout.NewSpacer(),
			exampleBox,
			layout.NewSpacer(),
			solutionButton, 
			hintBut,  
		)
		themeContent = container.NewBorder(nil, nil, nil, nil, themeContent)

		themeWindow.SetContent(themeContent)
		themeWindow.Show()
	})

	combo.SetSelected("Выберите нужную тему")

	centered := container.NewCenter(combo)
	content := container.NewVBox(
		title,
		welcomeText,
		layout.NewSpacer(),
		centered,
		layout.NewSpacer(),
	)

	w.SetContent(content)
	w.ShowAndRun()
}
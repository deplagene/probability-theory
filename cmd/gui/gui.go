package gui

import (
	"fmt"
	"homework/probability/types"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	w := a.NewWindow("Теория вероятностей")
	w.Resize(fyne.NewSize(800, 600))

	data := types.ProbabilityData{
		Theory:      "Вероятность события — это мера возможности наступления события. Она выражается числом от 0 до 1.",
		Formula:     "P(A) = m / n, где m — число благоприятных исходов, n — общее число исходов.",
		ExampleText: "Пример: В урне 10 шаров, из них 4 красных. Какова вероятность вытащить красный шар?",
	}

	title := widget.NewLabel("Теория Вероятностей")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	mainTheoryText := canvas.NewText("Теория с определением:", color.White)
	mainTheoryText.TextSize = 16
	mainTheoryText.TextStyle = fyne.TextStyle{Bold: true}
	theoryLabel := widget.NewLabel(data.Theory)
	theoryLabel.Wrapping = fyne.TextWrapWord
	theoryBox := container.NewVBox(
		mainTheoryText,
		theoryLabel,
	)
	theoryBox = container.NewBorder(nil, nil, nil, nil, theoryBox)

	mainFormulaText := canvas.NewText("Формула", color.White)
	mainFormulaText.TextStyle = fyne.TextStyle{Bold: true}
	mainFormulaText.TextSize = 16
	mainFormulaText.TextStyle = fyne.TextStyle{Bold: true}
	imageFormula := canvas.NewImageFromFile("images/formula1-1.png") 
	imageFormula.FillMode = canvas.ImageFillOriginal
	// formulaLabel := widget.NewLabel(data.Formula)
	// formulaLabel.Wrapping = fyne.TextWrapWord
	formulaBox := container.NewVBox(
		mainFormulaText,
		imageFormula,
	) 
	formulaBox = container.NewBorder(nil, nil, nil, nil, formulaBox)

	exampleLabel := widget.NewLabel(data.ExampleText)
	exampleLabel.Wrapping = fyne.TextWrapWord
	nEntry := widget.NewEntry()
	nEntry.SetPlaceHolder("Введите m (благоприятные исходы)")
	mEntry := widget.NewEntry()
	mEntry.SetPlaceHolder("Введите n (общие исходы)")

	// ! Короче, через виджеты создавай динамический текст, который можно будет изменить
	resultLabel := widget.NewLabel("Результат: ")
	resultLabel.TextStyle = fyne.TextStyle{Bold: true}

	calculateBtn := widget.NewButton("Рассчитать", func() {
		n, err1 := strconv.Atoi(nEntry.Text)
		m, err2 := strconv.Atoi(mEntry.Text)
		if err1 != nil || err2 != nil {
			resultLabel.SetText("Ошибка: введите корректные числа")
			return
		}
		p, err := g.service.Calculate(n, m)
		if err != nil {
			resultLabel.SetText(fmt.Sprintf("Ошибка: %v", err))
		} else {
			resultLabel.SetText(fmt.Sprintf("Результат: %.2f", p))
		}
	})

	mainExampleText := canvas.NewText("Пример задачи:", color.White)
	mainExampleText.TextSize = 16
	mainExampleText.TextStyle = fyne.TextStyle{Bold: true}
	exampleBox := container.NewVBox(
		mainExampleText,
		exampleLabel,
		nEntry,
		mEntry,
		calculateBtn,
		resultLabel,
	)
	exampleBox = container.NewBorder(nil, nil, nil, nil, exampleBox)

	rightSide := container.NewVBox(
		formulaBox,
		exampleBox,
	)

	mainContent := container.NewHBox(
		theoryBox,
		layout.NewSpacer(),
		rightSide,
	)

	content := container.NewVBox(
		title,
		mainContent,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
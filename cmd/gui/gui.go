package gui

import (
	"fmt"
	"homework/probability/types"
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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

func(g *Gui) Run() {
	a := app.New()
	w := a.NewWindow("Теория вероятностей")
	w.Resize(fyne.NewSize(600, 400))

	data := types.ProbabilityData {
		Title: "Комбинаторика",
		Theory: "Раздел математики, в котором рассматриваются задачи связанные с подсчетом количества комбинаций при определенных условиях",
		Formula: "n * m",
		ExampleText: "Имеется 100 рублей, мы хотим купить пирожное - 50р и мороженное - 50р. Пирожных 7 видов, а мороженных 4 вида. Сколькими способами мы это можем сделать?",
	}

	theoryLabel := widget.NewLabel(data.Theory)
	theoryLabel.Wrapping = fyne.TextWrapWord
	formulaLabel := widget.NewLabel(data.Formula)
	exampleLabel := widget.NewLabel(data.ExampleText)

	nEntry := widget.NewEntry()
    nEntry.SetPlaceHolder("Введите n")
    mEntry := widget.NewEntry()
    mEntry.SetPlaceHolder("Введите m")
    resultLabel := widget.NewLabel("Результат: ")

	calculateBtn := widget.NewButton("Рассчитать", func() {
        n, err1 := strconv.Atoi(nEntry.Text)
        m, err2 := strconv.Atoi(mEntry.Text)
        if err1 != nil || err2 != nil {
            resultLabel.SetText("Ошибка: введите корректные числа")
            return
        }
        p, err := g.service.Calculate(n, m)
        if err != nil {
            resultLabel.SetText(fmt.Sprintf("ошибка: %w", err))
        } else {
            resultLabel.SetText(fmt.Sprintf("результат: %v", p))
        }
    })

	content := container.NewVBox(
        widget.NewLabel("Теория:"),
        theoryLabel,
        widget.NewLabel("Формула:"),
        formulaLabel,
        widget.NewLabel("Пример задачи:"),
        exampleLabel,
        nEntry,
        mEntry,
        calculateBtn,
        resultLabel,
    )

	w.SetContent(content)
	w.ShowAndRun()
}
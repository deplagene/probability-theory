package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func reshator(a fyne.App) {
	win := a.NewWindow("Решатор задач")
	win.Resize(fyne.NewSize(700, 650))

	themes := []string{
		"Множества",
		"Элементы теории высказываний",
		"Комбинаторика",
		"Теория графов",
		"Изоморфизм графов",
		"Планарные графы",
		"Деревья",
		"Алгоритм Дейкстры",
	}

	description := widget.NewSelect(themes, nil)
	description.PlaceHolder = "Выберите тему"

	taskselect := widget.NewSelect([]string{}, nil)
	taskselect.PlaceHolder = "Выберите задание"
	taskselect.Disable()

	input := widget.NewEntry()
	input.SetPlaceHolder("Введите ответ...")

	task := widget.NewLabel("")
	task.Wrapping = fyne.TextWrapWord

	result := widget.NewLabel("")
	result.TextStyle = fyne.TextStyle{Bold: true}

	hint := widget.NewLabel("")
	hint.Wrapping = fyne.TextWrapWord
	hint.Hide()

	solut := widget.NewButton("Решение", nil)
	solut.Hide()

	hintB := widget.NewButton("Подсказка", nil)
	hintB.Hide()

	checkB := widget.NewButton("Проверить", nil)
	checkB.Disable()

	tasks := map[string][]struct {
		question    string
		answer      string
		explanation string
	}{
		"Множества": {
			{
				"В группе из 30 студентов: 18 изучают математику, 12 - физику, 5 - оба предмета. Сколько студентов не изучают ни один предмет?",
				"5",
				"Ответ: 5\nРешение:\n|A ∪ B| = |A| + |B| - |A ∩ B| = 18 + 12 - 5 = 25\n30 - 25 = 5 студентов",
			},
			{
				"В классе 35 учеников. 20 любят математику, 15 - информатику, 8 - оба предмета. Сколько учеников любят только математику?",
				"12",
				"Ответ: 12\nРешение:\nТолько математику: 20 - 8 = 12 учеников",
			},
			{
				"Из 50 сотрудников 30 знают английский, 25 - немецкий, 10 - оба языка. Сколько не знают ни одного языка?",
				"5",
				"Ответ: 5\nРешение:\n|A ∪ B| = 30 + 25 - 10 = 45\n50 - 45 = 5 сотрудников",
			},
		},
		"Элементы теории высказываний": {
			{
				"При каких значениях A и B выражение (A ∧ ¬B) → A будет ложным?",
				"A=0,B=1",
				"Ответ: A=0,B=1\nРешение:\nВыражение ложно только когда (A ∧ ¬B)=1 и A=0, что невозможно",
			},
			{
				"Упростите выражение: ¬(A ∨ B) ∧ (A ∨ ¬B)",
				"¬B ∧ A",
				"Ответ: ¬B ∧ A\nРешение:\nПрименяем законы де Моргана и дистрибутивности",
			},
			{
				"Сколько различных таблиц истинности существует для формул с 2 переменными?",
				"16",
				"Ответ: 16\nРешение:\nДля n переменных существует 2^(2^n) различных функций",
			},
		},
		"Комбинаторика": {
			{
				"Сколько способов расставить 5 книг на полке?",
				"120",
				"Ответ: 120\nРешение:\n5! = 120 перестановок",
			},
			{
				"Сколькими способами можно выбрать 3 дежурных из 10 человек?",
				"120",
				"Ответ: 120\nРешение:\nC(10,3) = 10!/(3!*7!) = 120",
			},
			{
				"Сколько 4-значных чисел можно составить из цифр 1,2,3,4 без повторений?",
				"24",
				"Ответ: 24\nРешение:\nP(4) = 4! = 24",
			},
		},
		"Теория графов": {
			{
				"Сколько ребер в полном графе с 5 вершинами?",
				"10",
				"Ответ: 10\nРешение:\nK(5) имеет 5*4/2 = 10 ребер",
			},
			{
				"Сколько ребер в дереве с 7 вершинами?",
				"6",
				"Ответ: 6\nРешение:\nДерево с n вершинами имеет n-1 ребро",
			},
			{
				"Может ли граф с 6 вершинами иметь степени: 5,4,3,3,2,1?",
				"нет",
				"Ответ: нет\nРешение:\nСумма степеней нечетная (18), что невозможно",
			},
		},
		"Изоморфизм графов": {
			{
				"Что значит, что два графа изоморфны?",
				"имеют одинаковую структуру",
				"Ответ: имеют одинаковую структуру\nРешение:\nИзоморфные графы — это графы с одинаковым числом вершин, рёбер и одинаковыми связями, но возможно с разными названиями вершин.",
			},
			{
				"Два графа имеют по 3 вершины и 2 ребра. Могут ли они быть изоморфны?",
				"да",
				"Ответ: да\nРешение:\nЕсли структура соединений одинакова, такие графы могут быть изоморфны независимо от названий вершин.",
			},
			{
				"Могут ли быть изоморфны графы, если один из них имеет 5 рёбер, а другой 6?",
				"нет",
				"Ответ: нет\nРешение:\nИзоморфные графы должны иметь одинаковое количество рёбер и вершин.",
			},
		},

		"Планарные графы": {
			{
				"Что означает, что граф планарен?",
				"его можно нарисовать без пересечений",
				"Ответ: его можно нарисовать без пересечений\nРешение:\nПланарный граф — это граф, который можно изобразить на плоскости так, чтобы рёбра не пересекались.",
			},
			{
				"Сколько рёбер максимум может быть в планарном графе с 5 вершинами?",
				"9",
				"Ответ: 9\nРешение:\nДля простого планарного графа с v вершинами: e ≤ 3v - 6.\nПодставим: 3×5 - 6 = 9.",
			},
			{
				"Является ли граф в виде квадрата с одной диагональю планарным?",
				"да",
				"Ответ: да\nРешение:\nТакой граф можно нарисовать без пересечений, он планарный.",
			},
		},

		"Деревья": {
			{
				"Сколько рёбер в дереве из 4 вершин?",
				"3",
				"Ответ: 3\nРешение:\nВ дереве с n вершинами всегда n - 1 ребро.\n4 - 1 = 3.",
			},
			{
				"Сколько вершин у дерева с 5 рёбрами?",
				"6",
				"Ответ: 6\nРешение:\nВ дереве количество вершин на 1 больше числа рёбер.\n5 + 1 = 6.",
			},
			{
				"Сколько листьев у дерева, где только одна вершина соединена со всеми остальными?",
				"4",
				"Ответ: 4\nРешение:\nЭто звёздчатое дерево: 1 центральная вершина и 4 листа.",
			},
		},

		"Алгоритм Дейкстры": {
			{
				"Для чего используется алгоритм Дейкстры?",
				"поиск кратчайшего пути",
				"Ответ: поиск кратчайшего пути\nРешение:\nАлгоритм находит кратчайшие пути от одной вершины до всех остальных в графе с неотрицательными весами.",
			},
			{
				"Можно ли использовать алгоритм Дейкстры в графе с отрицательными рёбрами?",
				"нет",
				"Ответ: нет\nРешение:\nАлгоритм Дейкстры не работает с отрицательными весами, так как может дать неправильные результаты.",
			},
			{
				"Сколько вершин обрабатывает алгоритм Дейкстры в графе из 3 вершин?",
				"3",
				"Ответ: 3\nРешение:\nАлгоритм обходит каждую вершину по одному разу — всего 3.",
			},
		},

	}	

	var currentTask struct {
		answer      string
		explanation string
		attempts    int
	}

	description.OnChanged = func(theme string) {
		taskselect.Enable()
		taskselect.Options = []string{"Задание 1", "Задание 2", "Задание 3"}
		taskselect.Selected = ""
		task.SetText("")
		input.SetText("")
		result.SetText("")
		solut.Hide()
		checkB.Disable()
	}

	taskselect.OnChanged = func(selected string) {
		if selected == "" {
			return
		}

		theme := description.Selected
		taskNum := map[string]int{"Задание 1": 0, "Задание 2": 1, "Задание 3": 2}[selected]

		t := tasks[theme][taskNum]
		currentTask.answer = t.answer
		currentTask.explanation = t.explanation
		currentTask.attempts = 0

		task.SetText(t.question)
		input.SetText("")
		result.SetText("")
		solut.Hide()
		checkB.Enable()
	}

	checkB.OnTapped = func() {
		user := input.Text
		if user == currentTask.answer {
			result.SetText("Верно! Ответ: " + currentTask.answer)
			solut.Hide()
		} else {
			currentTask.attempts++
			remaining := 3 - currentTask.attempts
			if remaining > 0 {
				result.SetText(fmt.Sprintf("Неверно! Осталось попыток: %d", remaining))
			} else {
				result.SetText("Неверно! Попытки закончились")
				solut.Show()
			}
		}
	}

	solut.OnTapped = func() {
		show(a, "Решение", currentTask.explanation)
	}

	form := container.NewVBox(
		widget.NewLabel("Решатор задач по дискретной математике"),
		description,
		taskselect,
		task,
		input,
		container.NewHBox(checkB, solut),
		result,
	)

	bg := canvas.NewRectangle(color.RGBA{45, 45, 45, 255})
	bg.CornerRadius = 12
	content := container.NewStack(bg, container.NewPadded(form))

	win.SetContent(content)
	win.Show()
}
func show(a fyne.App, title, content string) {
	win := a.NewWindow(title)
	label := widget.NewLabel(content)
	label.Wrapping = fyne.TextWrapWord

	box := container.NewVBox(label)
	bg := canvas.NewRectangle(color.RGBA{45, 45, 45, 255})
	bg.CornerRadius = 12

	win.SetContent(container.NewStack(bg, container.NewPadded(box)))
	win.Resize(fyne.NewSize(400, 300))
	win.Show()
}
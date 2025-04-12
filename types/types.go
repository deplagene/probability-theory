package types

import (
	"homework/probability/utils"

	"fyne.io/fyne/v2"
)

var ThemeIcons = map[string]fyne.Resource {
	"Множества": utils.ThemeIconsMustLoad("asserts/icons/icons8-lock-100.png"),
	"Элементы теории высказываний": utils.ThemeIconsMustLoad("asserts/icons/icons8-epistemology-64.png"),
	"Комбинаторика": utils.ThemeIconsMustLoad("asserts/icons/icons8-color-filter-96.png"),
	"Теория графов": utils.ThemeIconsMustLoad("asserts/icons/icons8-graphs-64.png"),
	"Изоморфизм графов": utils.ThemeIconsMustLoad("asserts/icons/icons8-graphs-64.png"),
	"Планарные графы": utils.ThemeIconsMustLoad("asserts/icons/icons8-graphs-64.png"),
	"Деревья": utils.ThemeIconsMustLoad("asserts/icons/icons8-tree-64.png"),
	"Алгоритм Дейкстры": utils.ThemeIconsMustLoad("asserts/icons/icons8-algorithm-100.png"),
}

var Themes = []string{
	"Множества",
	"Элементы теории высказываний",
	"Комбинаторика",
	"Теория графов",
	"Изоморфизм графов",
	"Планарные графы",
	"Деревья",
	"Алгоритм Дейкстры",
}

type ProbabilityData struct {
	Title              string
	Theory             string
	FormulaPath        string
	FormulaDescription string
	ExampleText        string
	ExampleImage       string
	Hint               string
	SolutionText       string
}

type MainService interface {
	Calculate(n, m int) (float64, error)
}

func ThemeSwitcher(theme string) ProbabilityData {
	switch theme {
	case "Множества":
		return ProbabilityData{
			Title:              "Множества",
			Theory:             "Понятие множеств ввел Георг Кантор...",
			SolutionText:       "Решение для множества: ...", 
		}

	case "Элементы теории высказываний":
		return ProbabilityData{
			Title:              "Элементы теории высказываний",
			Theory:             "Высказывание — это утверждение, которое может быть либо истинным, либо ложным.",
			FormulaDescription: "Основные логические операции:",
			ExampleImage:       "asserts/images/элементы_теории_выск.png",
			ExampleText: "Найти все наборы значений A, B, C ∈ {0, 1}, при которых логическое выражение (A ∨ B) → (¬C ∧ A) принимает значение 1 (истина).",
			Hint: "Порядок выполнения операций: 1. Скобки, 2. Отрицание (¬), 3. Конъюнкция (∧), 4. Дизъюнкция (∨), 5. Импликация (→)",
			SolutionText: "Решение примера:\n" +
				"1. Перебираем все возможные комбинации A, B, C ∈ {0,1}.\n" +
				"2. Рассматриваем выражение (A ∨ B) → (¬C ∧ A).\n" +
				"3. Получаем следующие результаты:\n" +
				"- A = 1, B = 0, C = 0: 1 → (1 ∧ 1) = 1\n" +
				"- A = 0, B = 0, C = 1: 0 → (1 ∧ 0) = 1\n" +
				"- A = 1, B = 1, C = 1: 1 → (0 ∧ 1) = 0 (ложь)\n" +
				"Итог: выбираем все комбинации, где выражение истинно.",
		}

	case "Комбинаторика":
		return ProbabilityData{}

	case "Теория графов":
		return ProbabilityData{}

	case "Изоморфизм графов":
		return ProbabilityData{
			Title:  "Изоморфизм графов",
			Theory: "",
		}

	case "Планарные графы":
		return ProbabilityData{
			Title:  "Планарные графы",
			Theory: "Граф называется планарным, если существует изоморфный ему плоский график (т.е график, расположенный на плоскости).",
		}

	case "Деревья":
		return ProbabilityData{
			Title:        "Деревья",
			Theory:       "Граф G называется ациклическим или лесом, если в нем нет циклов. Ациклический связный граф называется деревом.",
			ExampleImage: "asserts/images/деревья-схема 1.png", 
			ExampleText:  "Данный граф состоит из трех компонентов связанности, каждая из которых является деревом. Множество вершин образуют компоненту связанности графа, если для любых его вершин существует путь связываний их и никакая другая вершина не связана путем ни с какой вершиной этого множества.",
		}

	case "Алгоритм Дейкстры":
		return ProbabilityData{
			Title:  "Алгоритм Дейкстры",
			Theory: "Алгоритм Дейкстры используется для нахождения кратчайшего пути от одной вершины до всех остальных в взвешенном графе с неотрицательными весами.",
		}
		
	default:
		return ProbabilityData{
			Title:  "Ничего не найдено",
			Theory: "Теория для этой темы пока не добавлена.",
		}
	}
}

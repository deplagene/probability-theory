package types

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
}

type MainService interface {
	Calculate(n, m int) (float64, error)
}

func ThemeSwitcher(theme string) ProbabilityData {
	switch theme {
	case "Множества":
		return ProbabilityData{
			Title:  "Множества",
			Theory: "Понятие множеств ввел Георг Кантор, хоть и понятие множеств является неопределенным, обозначим, что множества - совокупность элементов любой природы, рассматриваемых как единое целое",
		}
	case "Элементы теории высказываний":
		return ProbabilityData{}
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
			Theory: "Граф называется планарным, если существует изоморфный ему плоский график ( т.е график, расположенный на плоскости ).",
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

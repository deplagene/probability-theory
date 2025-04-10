package themes

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type customTheme struct {
	fyne.Theme
	regularFont fyne.Resource
	boldFont    fyne.Resource
	italicFont  fyne.Resource
}

func NewCustomTheme() fyne.Theme {
	regularFont, err := os.ReadFile("asserts/fonts/JetBrainsMono-Regular.ttf")
	if err != nil {
		fmt.Println("Ошибка загрузки шрифта (Regular):", err)
	}

	boldFont, err := os.ReadFile("asserts/fonts/JetBrainsMono-Bold.ttf")
	if err != nil {
		fmt.Println("Ошибка загрузки шрифта (Bold):", err)
	}

	italicFont, err := os.ReadFile("asserts/fonts/JetBrainsMono-Italic.ttf")
	if err != nil {
		fmt.Println("Ошибка загрузки шрифта (Italic):", err)
	}

	var regularResource, boldResource, italicResource fyne.Resource
	if regularFont != nil {
		regularResource = fyne.NewStaticResource("JetBrainsMono-Regular.ttf", regularFont)
	}
	if boldFont != nil {
		boldResource = fyne.NewStaticResource("JetBrainsMono-Bold.ttf", boldFont)
	}
	if italicFont != nil {
		italicResource = fyne.NewStaticResource("JetBrainsMono-Italic.ttf", italicFont)
	}

	return &customTheme{
		Theme:       theme.DefaultTheme(),
		regularFont: regularResource,
		boldFont:    boldResource,
		italicFont:  italicResource,
	}
}

func (t *customTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		if t.boldFont != nil {
			return t.boldFont
		}
	}
	if style.Italic {
		if t.italicFont != nil {
			return t.italicFont
		}
	}

	if t.regularFont != nil {
		return t.regularFont
	}

	return theme.DefaultTheme().Font(style)
}

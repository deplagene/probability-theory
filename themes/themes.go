package themes

import (
	"fmt"
	"image/color"
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
		fmt.Println("Ошибка загрузки шрифта (Regular):", err)// Размер текста
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

func (t *customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 0x1E, G: 0x1E, B: 0x1E, A: 0xFF} // Тёмный фон (#1E1E1E)
	case theme.ColorNameForeground:
		return color.RGBA{R: 0xE0, G: 0xE0, B: 0xE0, A: 0xFF} // Светлый текст (#E0E0E0)
	case theme.ColorNameButton:
		return color.RGBA{R: 0xBB, G: 0x86, B: 0xFC, A: 0xFF} // Светло-фиолетовый для кнопок (#BB86FC)
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xFF} // Серый для placeholder (#808080)
	case theme.ColorNameShadow:
		return color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x33} // Тёмная тень
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 0x2D, G: 0x2D, B: 0x2D, A: 0xFF} // Тёмный фон для полей ввода (#2D2D2D)
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0xBB, G: 0x86, B: 0xFC, A: 0xFF} // Акцентный цвет (#BB86FC)
	case theme.ColorNameDisabled:
		return color.RGBA{R: 0x66, G: 0x66, B: 0x66, A: 0xFF} // Серый для отключённых элементов (#666666)
	default:
		return t.Theme.Color(name, variant)
	}
}

func (t *customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return t.Theme.Icon(name)
}

func (t *customTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 14 // Размер текста
	case theme.SizeNameHeadingText:
		return 24 // Размер заголовка
	case theme.SizeNameSubHeadingText:
		return 18 // Размер подзаголовка
	case theme.SizeNamePadding:
		return 16 // Отступы
	case theme.SizeNameInnerPadding:
		return 8 // Внутренние отступы
	default:
		return t.Theme.Size(name)
	}
}
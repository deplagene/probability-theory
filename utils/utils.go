package utils

import (
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/theme"
)

// todo: добавить иконки
func themeIconsMustLoad(path string) fyne.Resource {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Println("Icon load error:", err)
		return theme.FyneLogo()
	}
	return fyne.NewStaticResource(filepath.Base(path), data)
}
package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenDirectoryDialog() {
	runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "test",
	})
}

func (a *App) OpenFileDialog() {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "test",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "txt data",
				Pattern:     "*.txt;*.json",
			},
		},
	})
	if err != nil {
		fmt.Printf("err %s!", err)
	}
	fmt.Println(selection)
}

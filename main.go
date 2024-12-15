package main

import (
	"PackMe/app"
	"PackMe/internal/execution"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := app.NewApp()

	err := wails.Run(&options.App{
		Title:  "PackMe",
		Width:  760,
		Height: 400,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Mac: &mac.Options{
			OnFileOpen: func(filePath string) {
				execution.Unpack(filePath)
				runtime.Quit(context.Background())
			},
		},
		Bind: []interface{}{
			app,
		},
		OnDomReady: func(ctx context.Context) {
			directory, err := runtime.OpenDirectoryDialog(ctx, runtime.OpenDialogOptions{
				Title: "Select Directory to Pack",
			})
			if err != nil {
				println("Error:", err.Error())
				return
			}

			runtime.EventsEmit(ctx, "app:StartProcessing")
			resultPath := execution.Pack(ctx, directory)
			runtime.EventsEmit(ctx, "app:directorySelected", resultPath)
		},
	})

	if err != nil {
		panic(err)
	}
}

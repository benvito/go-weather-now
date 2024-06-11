package main

import (
	"embed"
	"go-weather-now/parser"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	image := parser.NewImage(parser.Url(""))

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "Go Weather Now",
		Width:             300,
		Height:            450,
		HideWindowOnClose: false,
		Frameless:         true,
		DisableResize:     true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 12, G: 12, B: 12, A: 200},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			image,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

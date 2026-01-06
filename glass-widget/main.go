package main

import (
	"embed"

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

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "glass-widget",
		Width:  450,
		Height: 380,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
        Frameless:        true,
        AlwaysOnTop:      true,
		Bind: []interface{}{
			app,
		},
        Windows: &windows.Options{
            WebviewIsTransparent: true,
            WindowIsTranslucent:  true,
            BackdropType:         windows.Acrylic,
        },
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

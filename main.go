/*
 * @Author: "xieyuhua" "1130
 * @Date: 2023-12-06 17:18:57
 * @LastEditors: "xieyuhua" "1130
 * @LastEditTime: 2023-12-13 10:40:22
 * @FilePath: \myproject\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	// "github.com/wailsapp/wails/v2/pkg/options/linux"
	// "github.com/wailsapp/wails/v2/pkg/options/mac"
	// "github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// AppMenu := menu.NewMenu()
	// FileMenu := AppMenu.AddSubmenu("File")
	// FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
	// 	// runtime.Quit(app.ctx)
	// })
	// FileMenu.AddSeparator()
	// FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
	// 	runtime.Quit(app.ctx)
	// })

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "数据流向",
		Width:  1224,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// Menu: AppMenu, // reference the menu above
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	// err := wails.Run(&options.App{
	//     Title:              "Menus Demo",
	//     Width:              800,
	//     Height:             600,
	//     DisableResize:      false,
	//     Fullscreen:         false,
	//     WindowStartState:   options.Maximised,
	//     Frameless:          true,
	//     MinWidth:           400,
	//     MinHeight:          400,
	//     MaxWidth:           1280,
	//     MaxHeight:          1024,
	//     StartHidden:        false,
	//     HideWindowOnClose:  false,
	//     BackgroundColour:   &options.RGBA{R: 0, G: 0, B: 0, A: 255},
	//     AlwaysOnTop:        false,
	//     AssetServer: &assetserver.Options{
	//         Assets:     assets,
	//         Handler:    assetsHandler,
	//         Middleware: assetsMidldeware,
	//     },
	//     Menu:               app.applicationMenu(),
	//     Logger:             nil,
	//     LogLevel:           logger.DEBUG,
	//     LogLevelProduction: logger.ERROR,
	//     OnStartup:          app.startup,
	//     OnDomReady:         app.domready,
	//     OnShutdown:         app.shutdown,
	//     OnBeforeClose:      app.beforeClose,
	//     CSSDragProperty:   "--wails-draggable",
	//     CSSDragValue:      "drag",
	//     EnableDefaultContextMenu: false,
	//     EnableFraudulentWebsiteDetection: false,
	//     ZoomFactor:           1.0,
	//     IsZoomControlEnabled: false,
	//     Bind: []interface{}{
	//         app,
	//     },
	//     ErrorFormatter: func(err error) any { return err.Error() },
	//     Windows: &windows.Options{
	//         WebviewIsTransparent:              false,
	//         WindowIsTranslucent:               false,
	//         BackdropType:                      windows.Mica,
	//         DisableWindowIcon:                 false,
	//         DisableFramelessWindowDecorations: false,
	//         WebviewUserDataPath:               "",
	//         WebviewBrowserPath:                "",
	//         Theme:                             windows.SystemDefault,
	//         CustomTheme: &windows.ThemeSettings{
	//             DarkModeTitleBar:   windows.RGB(20, 20, 20),
	//             DarkModeTitleText:  windows.RGB(200, 200, 200),
	//             DarkModeBorder:     windows.RGB(20, 0, 20),
	//             LightModeTitleBar:  windows.RGB(200, 200, 200),
	//             LightModeTitleText: windows.RGB(20, 20, 20),
	//             LightModeBorder:    windows.RGB(200, 200, 200),
	//         },
	//         // User messages that can be customised
	//         Messages *windows.Messages
	//         // OnSuspend is called when Windows enters low power mode
	//         OnSuspend func()
	//         // OnResume is called when Windows resumes from low power mode
	//         OnResume func(),
	//         WebviewGpuDisabled: false,
	//     },
	//     Mac: &mac.Options{
	//         TitleBar: &mac.TitleBar{
	//             TitlebarAppearsTransparent: true,
	//             HideTitle:                  false,
	//             HideTitleBar:               false,
	//             FullSizeContent:            false,
	//             UseToolbar:                 false,
	//             HideToolbarSeparator:       true,
	//         },
	//         Appearance:           mac.NSAppearanceNameDarkAqua,
	//         WebviewIsTransparent: true,
	//         WindowIsTranslucent:  false,
	//         About: &mac.AboutInfo{
	//             Title:   "My Application",
	//             Message: "© 2021 Me",
	//             Icon:    icon,
	//         },
	//     },
	//     Linux: &linux.Options{
	//         Icon: icon,
	//         WindowIsTranslucent: false,
	//         WebviewGpuPolicy: linux.WebviewGpuPolicyAlways,
	//         ProgramName: "wails"
	//     },
	//     Debug: options.Debug{
	//         OpenInspectorOnStartup: false,
	//     },
	// })

	if err != nil {
		println("Error:", err.Error())
	}
}

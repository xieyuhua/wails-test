/*
 * @Author: "xieyuhua" "1130
 * @Date: 2023-12-28 08:43:27
 * @LastEditors: "xieyuhua" "1130
 * @LastEditTime: 2023-12-28 08:59:51
 * @FilePath: \h5\wails_pf\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"embed"
	"log"
	// "os"
	"github.com/wailsapp/wails/v2"
	// "github.com/wailsapp/wails/v2/pkg/menu"
	// "github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()


    // // 创建菜单
    // appMenu := menu.NewMenu()
    // fileMenu := appMenu.AddSubmenu("文件")
    // fileMenu.AddText("退出", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
    //     os.Exit(0)
    // })

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "爆品订单管理系统",
		Width:             1500,
		Height:            900,
		MinWidth:          1500,
		MinHeight:         900,
		DisableResize:     true,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 16, G: 12, B: 42, A: 255},
		Assets:            assets,
		Menu:              nil,
		AlwaysOnTop:       false,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            true,
				UseToolbar:                 true,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "爆品订单管理系统",
				Message: "爆品订单管理系统",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
            Icon: icon,
        },
	})

	if err != nil {
		log.Fatal(err)
	}
}

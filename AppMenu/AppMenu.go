package AppMenu

import (
	"changeme/settings"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func NewMenu() *menu.Menu {
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("设置")
	FileMenu.AddText("设置存储路径", keys.CmdOrCtrl("s"), settingFileUrl)
	FileMenu.AddSeparator()
	return AppMenu
}

func settingFileUrl(_ *menu.CallbackData) {
	settings.ReadJson()
}

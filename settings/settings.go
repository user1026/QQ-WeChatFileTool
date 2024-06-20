package settings

import (
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

type JsonUrl struct {
	FileUrl string `json:"FileUrl"`
}

const settingsJsonUrl string = "settings.json"

func ReadJson() {
	file, err := os.Open(settingsJsonUrl)
	if err != nil {
		runtime.MessageDialog(C, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "",
			Message: "读取配置路径失败",
		})
		fmt.Println(err)
		return
	}
	var data []byte
	_, e := file.Read(data)
	if e != nil {
		runtime.MessageDialog(C, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "",
			Message: "读取配置文件失败",
		})
		return
	}
	var url JsonUrl
	json.Unmarshal(data, url)
	fmt.Println(string(data), url)
}

func SetJson() {

}

// gopherjs-electron

/* prepare
npm -g install electron-prebuilt
npm init -y
go get -u github.com/gopherjs/gopherjs
*/

/* build, run
gopherjs build index.go
electron .
*/

package main

import (
	"fmt"
	"path"
	"runtime"
	"encoding/json"
	"github.com/gopherjs/gopherjs/js"
)


func main() {
	fmt.Println("gopherjs+electronテスト")
	fmt.Printf("js: %v\n", js.Global.Get("process").Get("version"))
	fmt.Printf("electron: %v\n", js.Global.Get("process").Get("versions").Get("electron"))
	
	window_options := []byte(`
   		{
   		    "width": 1280
   		    , "height": 720
			, "web-preferences": {
				}
   		}
	`)
	var window_options_json map[string]interface{}
	err := json.Unmarshal(window_options, &window_options_json)
	if err != nil {
		panic(err)
	}

	//init electron app
	app := js.Global.Call("require", "app")
	js.Global.Call("require", "crash-reporter").Call("start")
	browserWindow := js.Global.Call("require", "browser-window")


	var mainWindow *js.Object

	app.Call("on", "window-all-closed", func() {
		if js.Global.Get("process").Get("platform").String() != "darwin" {
			app.Call("quit")
		}
	})

	app.Call("on", "ready", func() {
		mainWindow = browserWindow.New(window_options_json)
		
		_, filename, _, _ := runtime.Caller(1)
		url := "file:///" + path.Join(path.Dir(filename), "index.html")
		mainWindow.Call("loadUrl", url)
			
		mainWindow.Call("on", "closed", func() {
			mainWindow = nil
		})
	})
}

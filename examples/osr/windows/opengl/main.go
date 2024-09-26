package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/examples/osr/windows/opengl/form"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS
var app *cef.TCEFApplication

// CGO_ENABLED=1
// {$APPTYPE CONSOLE}
func main() {
	//libname.LibName = filepath.Join(os.Getenv("ENERGY_HOME"), "liblcl.dll")
	cef.GlobalInit(nil, nil)
	var window = &form.WindowForm{}
	//创建应用
	app = cef.NewApplication(true)
	// OSR 离屏渲染
	app.SetWindowlessRenderingEnabled(true)
	app.SetEnableGPU(true)
	// 指定消息模式
	app.SetExternalMessagePump(true)
	app.SetMultiThreadedMessageLoop(false)
	// create work schedule
	global := cef.GlobalWorkSchedulerCreate(nil)
	global.SetDefaultInterval(10)
	app.SetOnScheduleMessagePumpWork(nil)
	// 启动主进程, 执行后，二进制执行程序会被CEF多次执行创建子进程
	if app.StartMainProcess() {
		form.App = app
		fmt.Println("app run")
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources"
		server.Assets = resources
		server.StartHttpServer()
		lcl.RunApp(&window)
	}
}

package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/logger"
)

/*
子进程
这个示例演示了 主进程和 子进程相互独立出来，
子进程需要先编译好,提供给主进程(SetBrowseSubprocessPath)配置
*/
func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.CefLog_Debug)
	//全局配置初始化
	cef.GlobalInit(nil, nil)
	//创建Cef应用
	cefApp := cef.NewApplication()
	//主进程和子进程的变量绑定函数定义
	//cef.VariableBind.VariableCreateCallback(vars.VariableBind)
	//IPC通信
	IPCInit()
	//启动子进程
	cefApp.StartSubProcess()
	cefApp.Free()
}

// 渲染进程 IPC事件
func IPCInit() {
	fmt.Println("渲染进程IPC事件注册")
	//渲染进程监听的事件
	ipc.On("sub-process-on-event", func(context context.IContext) {
		fmt.Println("sub-process-on-event")
		//渲染进程处理程序....
		context.Result("返回结果")
	})
}

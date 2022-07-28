package main

import "C"
import (
	"bot_app/utils"
	"fmt"
	"syscall"
)

// 普通传参
func GoCallDll(a, b string) uintptr {
	dllFileName := "api.dll"
	// funName := Utf8ToGbkString("你好_世界")
	funName := utils.GBString("你好_世界")
	api := syscall.NewLazyDLL(dllFileName)
	fmt.Println("dll:", api.Name)
	hello_world := api.NewProc(funName)

	ret, _, err := hello_world.Call(utils.Str2ptr(a), utils.Str2ptr(b))
	if err != nil {
		fmt.Println(dllFileName, ret, err)
	}
	return ret
}
func GoCallDll2(a, b string) uintptr {
	dllFileName := "api.dll"
	// funName := Utf8ToGbkString("你好_世界")
	funName := utils.GBString("你好_世界")
	dllFile, _ := syscall.LoadLibrary(dllFileName)
	fmt.Println("+++++++syscall.LoadLibrary:", dllFile, "+++++++")
	defer syscall.FreeLibrary(dllFile)
	hello_world, err := syscall.GetProcAddress(dllFile, funName)

	fmt.Println("GetProcAddress", hello_world)
	ret, _, err := syscall.Syscall(hello_world,
		2,
		utils.Str2ptr(a),
		utils.Str2ptr(b),
		0)
	if err != nil {
		fmt.Println(dllFileName, ret, err)
	}
	return ret
}
func GoCallDll3(a, b string) uintptr {
	dllFileName := "api.dll"
	// funName := Utf8ToGbkString("你好_世界")
	// funName := "hello_world"
	funName := utils.GBString("你好_世界")
	DllTestDef := syscall.MustLoadDLL(dllFileName)
	hello_world := DllTestDef.MustFindProc(funName)

	fmt.Println("+++++++MustFindProc：", hello_world, "+++++++")
	ret, _, err := hello_world.Call(utils.Str2ptr(a), utils.Str2ptr(b))
	if err != nil {
		fmt.Println(dllFileName, ret, err)
	}
	return ret
}
func main() {
	fmt.Println("main")
	ret := utils.Ptr2str(GoCallDll("Hi,", "World!"))
	fmt.Println("main:", ret)
	// fmt.Println(Utf8ToGbkString("你好_世界"))
}

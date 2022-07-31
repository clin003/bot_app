package ui

import (
	"syscall"
	"unsafe"
)

// 默认UI

// MessageBox of Win32 API.
func MessageBox(hwnd uintptr, caption, title string, flags uint) int {
	p1, _ := syscall.UTF16PtrFromString(caption)
	p2, _ := syscall.UTF16PtrFromString(title)
	p1p := unsafe.Pointer(p1)
	p2p := unsafe.Pointer(p2)
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(p1p),
		uintptr(p2p),
		uintptr(flags))

	return int(ret)
}

// MessageBoxPlain of Win32 API.
func MessageBoxPlain(title, caption string) int {
	const (
		NULL  = 0
		MB_OK = 0
	)
	return MessageBox(NULL, caption, title, MB_OK)
}

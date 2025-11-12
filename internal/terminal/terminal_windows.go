//go:build windows
// +build windows

package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32                   = syscall.NewLazyDLL("kernel32.dll")
	getConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
)

type coord struct {
	X int16
	Y int16
}

type smallRect struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}

type consoleScreenBufferInfo struct {
	Size              coord
	CursorPosition    coord
	Attributes        uint16
	Window            smallRect
	MaximumWindowSize coord
}

func GetTerminalSize() (int, int) {
	var info consoleScreenBufferInfo
	handle := syscall.Handle(os.Stdout.Fd())

	ret, _, _ := getConsoleScreenBufferInfo.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&info)),
	)

	if ret == 0 {
		return 80, 24
	}

	return int(info.Window.Right - info.Window.Left + 1), int(info.Window.Top - info.Window.Bottom + 1)
}

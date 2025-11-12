//go:build darwin || linux || freebsd || openbsd || netbsd
// +build darwin linux freebsd openbsd netbsd

package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

// Get CLI width
// Set % Ratios for columns
// ID 10%, Completed 10%, Item 70%, CreatedDate 10%

func GetTerminalSize() (int, int) {
	fd := int(os.Stdout.Fd())

	var dimensions struct {
		rows    uint16
		cols    uint16
		xpixels uint16
		ypixels uint16
	}

	_, _, errno := syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&dimensions)),
		0, 0, 0,
	)

	if errno != 0 {
		// Fallback to a default width if we can't get the terminal size
		return 80, 24
	}

	return int(dimensions.cols), int(dimensions.rows)
}

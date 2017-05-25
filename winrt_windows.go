// +build windows
package winrt

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	modcombase, _ = syscall.LoadDLL("combase.dll")

	procRoInitialize, _        = modcombase.FindProc("RoInitialize")
	procRoUninitialize, _      = modcombase.FindProc("RoUninitialize")
	procWindowsCreateString, _ = modcombase.FindProc("WindowsCreateString")
	procWindowsDeleteString, _ = modcombase.FindProc("WindowsDeleteString")
)

func RoInitialize(initType uint32) (err error) {
	hr, _, _ := procRoInitialize.Call(uintptr(initType))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func RoUninitialize() {
	_, _, _ = procRoUninitialize.Call()
}

func WindowsCreateString(input uintptr, length uint32, output HSTRING) (err error) {
	hr, _, _ := procWindowsCreateString.Call(input, uintptr(length), uintptr(unsafe.Pointer(output)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func WindowsDeleteString(input HSTRING) (err error) {
	hr, _, _ := procWindowsDeleteString.Call(uintptr(unsafe.Pointer(input)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

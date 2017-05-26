// +build windows
package winrt

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	modcombase, _ = syscall.LoadDLL("combase.dll")

	procRoInitialize, _        = modcombase.FindProc("RoInitialize")
	procRoUninitialize, _      = modcombase.FindProc("RoUninitialize")
	procRoActivateInstance, _  = modcombase.FindProc("RoActivateInstance")
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

func RoActivateInstance(hstr *HSTRING, ins interface{}) (err error) {
	insValue := reflect.ValueOf(ins).Elem()
	hr, _, _ := procRoActivateInstance.Call(uintptr(unsafe.Pointer(hstr)), insValue.Addr().Pointer())
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func WindowsCreateString(input *uint16, length uint32, hstr **HSTRING) (err error) {
	hr, _, _ := procWindowsCreateString.Call(uintptr(unsafe.Pointer(input)), uintptr(length), uintptr(unsafe.Pointer(hstr)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func WindowsDeleteString(hstr **HSTRING) (err error) {
	hr, _, _ := procWindowsDeleteString.Call(uintptr(unsafe.Pointer(hstr)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

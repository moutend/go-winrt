// +build windows

package winrt

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func insGetIids(ins *IInspectable, iidCount uint32, iids **ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		ins.VTable().GetIids,
		3,
		uintptr(unsafe.Pointer(ins)),
		uintptr(iidCount),
		uintptr(unsafe.Pointer(&iids)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func insGetRuntimeClassName(ins *IInspectable, className *HSTRING) (err error) {
	hr, _, _ := syscall.Syscall(
		ins.VTable().GetRuntimeClassName,
		2,
		uintptr(unsafe.Pointer(ins)),
		uintptr(unsafe.Pointer(className)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func insGetTrustLevel(ins *IInspectable, trustLevel *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		ins.VTable().GetTrustLevel,
		2,
		uintptr(unsafe.Pointer(ins)),
		uintptr(unsafe.Pointer(trustLevel)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

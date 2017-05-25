package winrt

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IInspectable struct {
	ole.IUnknown
}

type IInspectableVtbl struct {
	ole.IUnknownVtbl
	GetIids             uintptr
	GetRuntimeClassName uintptr
	GetTrustLevel       uintptr
}

func (v *IInspectable) VTable() *IInspectableVtbl {
	return (*IInspectableVtbl)(unsafe.Pointer(v.RawVTable))
}

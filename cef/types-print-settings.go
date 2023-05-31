//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// PDF 打印回调 PdfPrintCallbackRef.New
package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// PrintSettingsRef -> ICefPdfPrintCallback
var PrintSettingsRef printSettings

type printSettings uintptr

func (*printSettings) New() *ICefPrintSettings {
	var result uintptr
	imports.Proc(internale_CefPrintSettingsRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefPrintSettings{instance: unsafe.Pointer(result)}
}

// Instance 实例
func (m *ICefPrintSettings) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPrintSettings) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefPrintSettings) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefPrintSettings) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefPrintSettings) SetOrientation(landscape bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetOrientation).Call(m.Instance(), api.PascalBool(landscape))
}

func (m *ICefPrintSettings) IsLandscape() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_IsLandscape).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefPrintSettings) SetPrinterPrintableArea(physicalSizeDeviceUnits *TCefSize, printableAreaDeviceUnits *TCefRect, landscapeNeedsFlip bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetPrinterPrintableArea).Call(m.Instance(), uintptr(unsafe.Pointer(physicalSizeDeviceUnits)), uintptr(unsafe.Pointer(printableAreaDeviceUnits)), api.PascalBool(landscapeNeedsFlip))
}

func (m *ICefPrintSettings) SetDeviceName(name string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetDeviceName).Call(m.Instance(), api.PascalStr(name))
}

func (m *ICefPrintSettings) GetDeviceName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_GetDeviceName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefPrintSettings) SetDpi(dpi int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetDpi).Call(m.Instance(), uintptr(dpi))
}

func (m *ICefPrintSettings) GetDpi() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_GetDpi).Call(m.Instance())
	return int32(r1)
}

func (m *ICefPrintSettings) SetPageRanges(ranges []TCefRange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetPageRanges).Call(m.Instance(), uintptr(unsafe.Pointer(&ranges[0])), uintptr(int32(len(ranges))))
}

func (m *ICefPrintSettings) GetPageRangesCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_GetPageRangesCount).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefPrintSettings) GetPageRanges() []TCefRange {
	if !m.IsValid() {
		return nil
	}
	count := m.GetPageRangesCount()
	if count > 0 {
		var resultRangesPtr uintptr
		imports.Proc(internale_CefPrintSettings_GetPageRanges).Call(m.Instance(), uintptr(unsafe.Pointer(&resultRangesPtr)))
		if resultRangesPtr != 0 {
			var rangeSize = unsafe.Sizeof(TCefRange{})
			resultRanges := make([]TCefRange, count, count)
			for i := 0; i < int(count); i++ {
				resultRanges[i] = *(*TCefRange)(common.GetParamPtr(resultRangesPtr, i*int(rangeSize)))
			}
			return resultRanges
		}
	}
	return nil
}

func (m *ICefPrintSettings) SetSelectionOnly(selectionOnly bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetSelectionOnly).Call(m.Instance(), api.PascalBool(selectionOnly))
}

func (m *ICefPrintSettings) IsSelectionOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_IsSelectionOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefPrintSettings) SetCollate(collate bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetCollate).Call(m.Instance(), api.PascalBool(collate))
}

func (m *ICefPrintSettings) WillCollate() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_WillCollate).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefPrintSettings) SetColorModel(model consts.TCefColorModel) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetColorModel).Call(m.Instance(), model.ToPtr())
}

func (m *ICefPrintSettings) GetColorModel() consts.TCefColorModel {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_GetColorModel).Call(m.Instance())
	return consts.TCefColorModel(r1)
}

func (m *ICefPrintSettings) SetCopies(copies int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetCopies).Call(m.Instance(), uintptr(copies))
}

func (m *ICefPrintSettings) GetCopies() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_GetCopies).Call(m.Instance())
	return int32(r1)
}

func (m *ICefPrintSettings) SetDuplexMode(mode consts.TCefDuplexMode) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefPrintSettings_SetDuplexMode).Call(m.Instance(), mode.ToPtr())
}

func (m *ICefPrintSettings) GetDuplexMode() consts.TCefDuplexMode {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_CefPrintSettings_GetDuplexMode).Call(m.Instance())
	return consts.TCefDuplexMode(r1)
}

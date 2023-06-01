//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// PostDataElementRef -> ICefPostDataElement
var PostDataElementRef postDataElement

// postDataElement
type postDataElement uintptr

func (m *postDataElement) New() *ICefPostDataElement {
	var result uintptr
	imports.Proc(internale_PostDataElementRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefPostDataElement{instance: unsafe.Pointer(result)}
}

func (m *postDataElement) UnWrap(data *ICefPostDataElement) *ICefPostDataElement {
	var result uintptr
	imports.Proc(internale_PostDataElementRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	data.instance = unsafe.Pointer(result)
	return data
}

// Instance 实例
func (m *ICefPostDataElement) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPostDataElement) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *ICefPostDataElement) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_PostDataElement_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefPostDataElement) SetToEmpty() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_PostDataElement_SetToEmpty).Call(m.Instance())
}

func (m *ICefPostDataElement) SetToFile(fileName string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_PostDataElement_SetToFile).Call(m.Instance(), api.PascalStr(fileName))
}

func (m *ICefPostDataElement) SetToBytes(bytes []byte) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_PostDataElement_SetToBytes).Call(m.Instance(), uintptr(uint32(len(bytes))), uintptr(unsafe.Pointer(&bytes[0])))
}

func (m *ICefPostDataElement) GetType() consts.TCefPostDataElementType {
	if !m.IsValid() {
		return -1
	}
	r1, _, _ := imports.Proc(internale_PostDataElement_GetType).Call(m.Instance())
	return consts.TCefPostDataElementType(r1)
}

func (m *ICefPostDataElement) GetFile() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_PostDataElement_GetFile).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefPostDataElement) GetBytesCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_PostDataElement_GetBytesCount).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefPostDataElement) GetBytes() (result []byte, count uint32) {
	if !m.IsValid() {
		return nil, 0
	}
	size := m.GetBytesCount()
	if size <= 0 {
		return nil, 0
	}
	result = make([]byte, size)
	r1, _, _ := imports.Proc(internale_PostDataElement_GetBytes).Call(m.Instance(), uintptr(size), uintptr(unsafe.Pointer(&result)))
	count = uint32(r1)
	return result, count
}

func (m *TCefPostDataElementArray) Get(index uint32) *ICefPostDataElement {
	if index < m.postDataElementLength {
		return &ICefPostDataElement{instance: unsafe.Pointer(common.GetParamOf(int(index), m.postDataElement))}
	}
	return nil
}

func (m *TCefPostDataElementArray) Size() uint32 {
	return m.postDataElementLength
}

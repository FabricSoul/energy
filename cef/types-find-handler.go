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
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ************************** creates ************************** //

// FindHandlerRef -> ICefFindHandler
var FindHandlerRef findHandler

type findHandler uintptr

func (*findHandler) New() *ICefFindHandler {
	var result uintptr
	imports.Proc(internale_CefFindHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFindHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// ************************** impl ************************** //

// Instance 实例
func (m *ICefFindHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefFindHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefFindHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefFindHandler) SetOnFindResult(fn onFindResult) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefFindHandler_OnFindResult).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onFindResult func(browser *ICefBrowser, count int32, selectionRect *TCefRect, activeMatchOrdinal int32, finalUpdate bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onFindResult:
			browse := &ICefBrowser{instance: getPtr(0)}
			count := int32(getVal(1))
			selectionRect := (*TCefRect)(getPtr(2))
			activeMatchOrdinal := int32(getVal(3))
			finalUpdate := api.GoBool(getVal(4))
			fn.(onFindResult)(browse, count, selectionRect, activeMatchOrdinal, finalUpdate)
		default:
			return false
		}
		return true
	})
}

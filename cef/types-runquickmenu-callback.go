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
	"github.com/energye/energy/consts"
)

// Instance 实例
func (m *ICefRunQuickMenuCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefRunQuickMenuCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefRunQuickMenuCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefRunQuickMenuCallback) Cont(commandId int32, eventFlags consts.TCefEventFlags) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RunContextMenuCallback_Cont).Call(m.Instance(), uintptr(commandId), eventFlags.ToPtr())
}

func (m *ICefRunQuickMenuCallback) Cancel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_RunContextMenuCallback_Cancel).Call(m.Instance())
}

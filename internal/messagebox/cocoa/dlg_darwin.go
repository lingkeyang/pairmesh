package cocoa

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include <sys/syslimits.h>
// #include "dlg.h"
import "C"

import (
	"unsafe"
)

type AlertParams struct {
	p C.AlertDlgParams
}

func mkAlertParams(msg, title string, style C.AlertStyle) *AlertParams {
	a := AlertParams{C.AlertDlgParams{msg: C.CString(msg), style: style}}
	if title != "" {
		a.p.title = C.CString(title)
	}
	return &a
}

func (a *AlertParams) run() C.DlgResult {
	return C.alertDlg(&a.p)
}

func (a *AlertParams) free() {
	C.free(unsafe.Pointer(a.p.msg))
	if a.p.title != nil {
		C.free(unsafe.Pointer(a.p.title))
	}
}

func Confirm(msg, title string) bool {
	a := mkAlertParams(msg, title, C.MSG_YESNO)
	defer a.free()
	return a.run() == C.DLG_OK
}

func Info(msg, title string) {
	a := mkAlertParams(msg, title, C.MSG_INFO)
	defer a.free()
	a.run()
}

func Error(msg, title string) {
	a := mkAlertParams(msg, title, C.MSG_ERROR)
	defer a.free()
	a.run()
}

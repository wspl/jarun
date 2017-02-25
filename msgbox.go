package main

import (
	"syscall"
)

import (
	"github.com/lxn/win"
)

type MsgBoxStyle uint

const (
	MsgBoxOK                MsgBoxStyle = win.MB_OK
	MsgBoxOKCancel          MsgBoxStyle = win.MB_OKCANCEL
	MsgBoxAbortRetryIgnore  MsgBoxStyle = win.MB_ABORTRETRYIGNORE
	MsgBoxYesNoCancel       MsgBoxStyle = win.MB_YESNOCANCEL
	MsgBoxYesNo             MsgBoxStyle = win.MB_YESNO
	MsgBoxRetryCancel       MsgBoxStyle = win.MB_RETRYCANCEL
	MsgBoxCancelTryContinue MsgBoxStyle = win.MB_CANCELTRYCONTINUE
	MsgBoxIconHand          MsgBoxStyle = win.MB_ICONHAND
	MsgBoxIconQuestion      MsgBoxStyle = win.MB_ICONQUESTION
	MsgBoxIconExclamation   MsgBoxStyle = win.MB_ICONEXCLAMATION
	MsgBoxIconAsterisk      MsgBoxStyle = win.MB_ICONASTERISK
	MsgBoxUserIcon          MsgBoxStyle = win.MB_USERICON
	MsgBoxIconWarning       MsgBoxStyle = win.MB_ICONWARNING
	MsgBoxIconError         MsgBoxStyle = win.MB_ICONERROR
	MsgBoxIconInformation   MsgBoxStyle = win.MB_ICONINFORMATION
	MsgBoxIconStop          MsgBoxStyle = win.MB_ICONSTOP
	MsgBoxDefButton1        MsgBoxStyle = win.MB_DEFBUTTON1
	MsgBoxDefButton2        MsgBoxStyle = win.MB_DEFBUTTON2
	MsgBoxDefButton3        MsgBoxStyle = win.MB_DEFBUTTON3
	MsgBoxDefButton4        MsgBoxStyle = win.MB_DEFBUTTON4
)

func MsgBox(title, message string, style MsgBoxStyle) int {
	msg, _ := syscall.UTF16PtrFromString(message)
	tit, _ := syscall.UTF16PtrFromString(title)
	return int(win.MessageBox(win.HWND(0), msg, tit, uint32(style)))
}

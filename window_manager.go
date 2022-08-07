package main

import (
    "fmt"
    "github.com/TheTitanrain/w32"
    "errors"
    "syscall"
)

var (
    user32                 = syscall.NewLazyDLL("user32.dll")
    procEnumWindows        = user32.NewProc("EnumWindows")
)
 
func ModifyWindow(name string, focused bool) {
    window, err := FindWindow(name)
    if err != nil {
        panic(err)
    }
    var flags uintptr = w32.WS_EX_TOPMOST|w32.WS_EX_LAYERED
    if !focused {
        flags = flags|w32.WS_EX_TRANSPARENT|w32.WS_EX_NOACTIVATE
    }
    w32.SetWindowLongPtr(w32.HWND(window), -20, flags)
}
 
func FindWindow(title string) (w32.HANDLE, error) {
    var hwnd syscall.Handle
    cb := func(h syscall.Handle, p uintptr) uintptr {
        windowText := w32.GetWindowText(w32.HWND(h))
        if windowText == title {
            hwnd = h
            return 0 // stop enumaration
        }
        return 1 // continue enumeration
    }
    EnumWindows(cb, 0)
    if hwnd == 0 {
        return 0, errors.New(fmt.Sprintf("Could not find window with title '%s'", title))
    }
    return w32.HANDLE(hwnd), nil
}
 
func EnumWindows(fn interface{}, lparam uintptr) uintptr {
    ret, _, _ := procEnumWindows.Call(
        syscall.NewCallback(fn),
        lparam,
    )
    return ret
}
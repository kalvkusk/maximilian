package main

import (
    "log"
    "strings"
    "syscall"

    "github.com/AllenDang/w32"
    "github.com/hnakamur/w32syscall"
)

func main() {
    err := w32syscall.EnumWindows(func(hwnd syscall.Handle, lparam uintptr) bool {
        h := w32.HWND(hwnd)
        text := w32.GetWindowText(h)
        if strings.Contains(text, "Rainforest VM") {
            w32.MoveWindow(h, 0, 0, 1440, 900, true)
        }
        return true
    }, 0)
    if err != nil {
        log.Fatalln(err)
    }
}

package main

import (
    "unsafe"
    "syscall"
    "strings"
    windows "golang.org/x/sys/windows"
)

const TH32CS_SNAPPROCESS = 0x00000002

type WindowsProcess struct {
    ProcessID       int
    ParentProcessID int
    Exe             string
}

func processes() ([]WindowsProcess, error) {
    handle, err := windows.CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)
    if err != nil {
        return nil, err
    }
    defer windows.CloseHandle(handle)

    var entry windows.ProcessEntry32
    entry.Size = uint32(unsafe.Sizeof(entry))
    // get the first process
    err = windows.Process32First(handle, &entry)
    if err != nil {
        return nil, err
    }

    results := make([]WindowsProcess, 0, 50)
    for {
        results = append(results, newWindowsProcess(&entry))

        err = windows.Process32Next(handle, &entry)
        if err != nil {
            // windows sends ERROR_NO_MORE_FILES on last process
            if err == syscall.ERROR_NO_MORE_FILES {
                return results, nil
            }
            return nil, err
        }
    }
}

func findProcessByName(processes []WindowsProcess, name string) *WindowsProcess {
    for _, p := range processes {
        if strings.ToLower(p.Exe) == strings.ToLower(name) {
            return &p
        }
    }
    return nil
}

func newWindowsProcess(e *windows.ProcessEntry32) WindowsProcess {
    // Find when the string ends for decoding
    end := 0
    for {
        if e.ExeFile[end] == 0 {
            break
        }
        end++
    }

    return WindowsProcess{
        ProcessID:       int(e.ProcessID),
        ParentProcessID: int(e.ParentProcessID),
        Exe:             syscall.UTF16ToString(e.ExeFile[:end]),
    }
}

const defaultName string = "UE4Game-Win64-Shipping.exe"

func bindDefaultProcess() (uint32, bool) {
    procs, err := processes()
	if err != nil {
		return 0, false
	}

	explorer := findProcessByName(procs, defaultName)
	if explorer == nil {
		return 0, false
	}
    
	pid := uint32(explorer.ProcessID)
    return pid, true
}
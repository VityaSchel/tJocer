package main

import (
	"encoding/binary"
	"math"
	"unsafe"
	// "fmt"
	// "path/filepath"
	// "strconv"

	// "github.com/0xrawsec/golang-win32/win32"
	// kernel32 "github.com/0xrawsec/golang-win32/win32/kernel32"
	windows "golang.org/x/sys/windows"
)

var handle windows.Handle
var procReadProcessMemory *windows.Proc

func memoryReadInit(pid uint32) {
  handle, _ = windows.OpenProcess(0x0010 | windows.PROCESS_VM_READ | windows.PROCESS_QUERY_INFORMATION, false, pid)
  procReadProcessMemory = windows.MustLoadDLL("kernel32.dll").MustFindProc("ReadProcessMemory")
}

// func memoryReadInit(pid uint32) {
// 	win32handle, _ := kernel32.OpenProcess(0x0010 | windows.PROCESS_VM_READ | windows.PROCESS_QUERY_INFORMATION, win32.BOOL(0), win32.DWORD(pid))
// 	moduleHandles, _ := kernel32.EnumProcessModules(win32handle)
// 	for _, moduleHandle := range moduleHandles {
// 		s, _ := kernel32.GetModuleFilenameExW(win32handle, moduleHandle)
// 		targetModuleFilename := "UE4Game-Win64-Shipping.exe"
// 		if(filepath.Base(s) == targetModuleFilename) {
// 			fmt.Println(strconv.FormatInt(int64(moduleHandle), 16))
// 			break
// 		}
// 	}
// 	// windows.GetModuleInformation(handle, )
// }

func memoryReadClose() {
  windows.CloseHandle(handle)
}

func readMemoryAt(address int) float32 {
  var (
		data [4]byte
		length uint32
	)

	// BOOL ReadProcessMemory(HANDLE hProcess, LPCVOID lpBaseAddress, LPVOID lpBuffer, DWORD nSize, LPDWORD lpNumberOfBytesRead)
	procReadProcessMemory.Call(
		uintptr(handle), 
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)), 
		uintptr(unsafe.Pointer(&length)),
	)

	// println(a, b, c)
	
	bits := binary.LittleEndian.Uint32(data[:])
	float := math.Float32frombits(bits)

	return float
}

type staticPointer struct {
	pointer uint32
	offsets []string
}


// func getAddresses() {
// 	xPositionPointer := staticPointer{2518790, []string{"2E4", "10", "8", "8", "8", "78", "5E0"}}
// 	print(calculateAddress(xPositionPointer))
// }

// func calculateAddress(pointer staticPointer) int64 {
// 	var baseAddress int64 = "UE4Game-Win64-Shipping.exe"+pointer.pointer // convert it to decimal
// 	for _, offset := range pointer.offsets {
// 		result, _ := strconv.ParseInt(offset, 16, 64)
// 		baseAddress += result
// 		//sumHex(baseAddress, offset)
// 	}
// 	return baseAddress
// }

// func sumHex(a string, b string) string {
// 	start10, _ := strconv.ParseInt(a, 16, 0)
// 	sum10, _ := strconv.ParseInt(b, 16, 0)
// 	result10 := start10 + sum10
// 	result16 := strconv.FormatInt(result10, 16)
// 	return result16
// }
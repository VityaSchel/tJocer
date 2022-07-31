package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"unsafe"

	// "log"
	// "fmt"
	"path/filepath"

	"github.com/0xrawsec/golang-win32/win32"
	kernel32 "github.com/0xrawsec/golang-win32/win32/kernel32"
	windows "golang.org/x/sys/windows"
	// kiwi "github.com/Andoryuuta/kiwi"
	// "github.com/Andoryuuta/kiwi"
	// w32 "github.com/Andoryuuta/kiwi/w32"
)

var handle windows.Handle
var procReadProcessMemory *windows.Proc
var baseAddress int64

func memoryReadInit(pid uint32) (int64, string) {
  handle, _ = windows.OpenProcess(0x0010 | windows.PROCESS_VM_READ | windows.PROCESS_QUERY_INFORMATION, false, pid)
	if(handle == 0) {
		return -1, "NO_HANDLE"
	}

  procReadProcessMemory = windows.MustLoadDLL("kernel32.dll").MustFindProc("ReadProcessMemory")
	
	win32handle, _ := kernel32.OpenProcess(0x0010 | windows.PROCESS_VM_READ | windows.PROCESS_QUERY_INFORMATION, win32.BOOL(0), win32.DWORD(pid))
	moduleHandles, _ := kernel32.EnumProcessModules(win32handle)
	for _, moduleHandle := range moduleHandles {
		s, _ := kernel32.GetModuleFilenameExW(win32handle, moduleHandle)
		targetModuleFilename := "UE4Game-Win64-Shipping.exe"
		if(filepath.Base(s) == targetModuleFilename) {
			info, _ := kernel32.GetModuleInformation(win32handle, moduleHandle)
			baseAddress = int64(info.LpBaseOfDll)
			return baseAddress, ""
		}
	}

	return -1, "BASE_ADDRESS_NOT_FOUND"
}

func memoryReadClose() {
  windows.CloseHandle(handle)
}

func readMemoryAtByte4(address int64) (value uint32, err bool) {
	var (
		data [4]byte
		length uint32
	)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			value, err = 0, true
		}
	}()

	procReadProcessMemory.Call(
		uintptr(handle), 
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)), 
		uintptr(unsafe.Pointer(&length)),
	)
	
  byte4 := binary.LittleEndian.Uint32(data[:])
	return byte4, false
}


func readMemoryAtFloat32(address int64) (value float32, err bool) {
	result, err := readMemoryAtByte4(address)
	if(err) {
		return 0, true
	}

	float := math.Float32frombits(result)
	return float, err
}

func readMemoryAtByte8(address int64) (value uint64, err bool) {
	var (
		data [8]byte
		length uint32
	)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			value, err = 0, true
		}
	}()

	procReadProcessMemory.Call(
		uintptr(handle), 
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)), 
		uintptr(unsafe.Pointer(&length)),
	)
	
  byte8 := binary.LittleEndian.Uint64(data[:])
	return byte8, false
}

type staticPointer struct {
	baseOffset int64
	offsets []string
}


func GetAddresses() (int64, int64, int64) {
	xPositionPointer := staticPointer{0x0250C5B0, []string{"E80", "F0", "80", "60", "3D8", "160", "F8"}}
	zPositionPointer := staticPointer{0x0250C5B0, []string{"E84", "F0", "80", "60", "3D8", "160", "F8"}}
	
	xPositionAddress := calculateAddress(xPositionPointer)
	zPositionAddress := calculateAddress(zPositionPointer)

	return xPositionAddress, zPositionAddress, baseAddress
}

// * are constants
// 1. get base hex address of process
// 2. add *base offset* hex to base address
// 3. go to loop in array of offsets (from end to start):
// 3.1. take *offset*
// 3.2. add this offset to [value]
// 3.3. now you found the address that you have to read in order to get next address
// 3.4. !!! read [value] address and convert it with 8 bytes
// 3.5. !!! the address is the new [value]
// the "!!!" signs indicating that they should not be executed in last iteration of loop
// 4. return [value] this is the final address

func calculateAddress(pointer staticPointer) int64 {
	startingPointer := baseAddress + pointer.baseOffset
	startingAddress, _ := readMemoryAtByte8(startingPointer)
	var value int64 = int64(startingAddress)

	for i := len(pointer.offsets)-1; i >= 0; i-- {
		offset := pointer.offsets[i]
		addressPointer := sumHexSII(offset, int64(value))

		if(i > 0) {
			nextAddressDecimal, _ := readMemoryAtByte8(addressPointer)			
			value = int64(nextAddressDecimal)
		} else {
			value = addressPointer
		}
	}
	return value
}
package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"unsafe"

	// "log"
	// "fmt"
	// "path/filepath"
	"strconv"

	// "github.com/0xrawsec/golang-win32/win32"
	// kernel32 "github.com/0xrawsec/golang-win32/win32/kernel32"
	windows "golang.org/x/sys/windows"
	// kiwi "github.com/Andoryuuta/kiwi"
	// "github.com/Andoryuuta/kiwi"
	// w32 "github.com/Andoryuuta/kiwi/w32"
)

var handle windows.Handle
var procReadProcessMemory *windows.Proc

func memoryReadInit(pid uint32) {
  handle, _ = windows.OpenProcess(0x0010 | windows.PROCESS_VM_READ | windows.PROCESS_QUERY_INFORMATION, false, pid)
  procReadProcessMemory = windows.MustLoadDLL("kernel32.dll").MustFindProc("ReadProcessMemory")
	memoryReadInit2(pid)
}

func memoryReadInit2(pid uint32) {
	// fmt.Println(pid)
	// win32handle, _ := kernel32.OpenProcess(0x0010 | windows.PROCESS_VM_READ | windows.PROCESS_QUERY_INFORMATION, win32.BOOL(0), win32.DWORD(pid))
	// moduleHandles, _ := kernel32.EnumProcessModules(win32handle)
	// for _, moduleHandle := range moduleHandles {
	// 	s, _ := kernel32.GetModuleFilenameExW(win32handle, moduleHandle)
	// 	// kernel32.GetModule
	// 	targetModuleFilename := "UE4Game-Win64-Shipping.exe"
	// 	if(filepath.Base(s) == targetModuleFilename) {
	// 		// info, _ := kernel32.GetModuleInformation(win32handle, moduleHandle)
	// 		// // kernel32.MODULEINFO
	// 		// getModuleInformation := windows.MustLoadDLL("kernel32.dll").MustFindProc("GetModuleInformation")
	// 		// getModuleInformation.Call(

	// 		// )
	// 		// kernel32.GetModuleHandleW(targetModuleFilename)
	// 		// kernel32.GetModuleInformation(win32handle, moduleHandle)
	// 		proc, _ := kiwi.GetProcessByPID(int(pid))
	// 		base, _ := proc.GetModuleBase(targetModuleFilename)

	// 		fmt.Println("!! 1   ", base)
	// 		// w32.EnumProcesses()
	// 		// fmt.Println("!! 2   ", int64(info.LpBaseOfDll))
	// 		// fmt.Println("!! 3   ", strconv.FormatInt(int64(info.LpBaseOfDll), 16))
	// 		break
	// 	}
	// }
	// // windows.GetModuleInformation(handle, )

	// // proc, _ := kiwi.GetProcessByFileName("UE4Game-Win64-Shipping.exe")
	// // proc2, _ := kiwi.GetProcessByPID(int(pid))

	// // fmt.Println(proc.GetModuleBase("UE4Game-Win64-Shipping.exe"))
	// // fmt.Println(proc2.GetModuleBase(""))
	
	// // w32.EnumProcesses([pid], )
	// // proc.ReadFloat32()

	// // var me32 w32.MODULEENTRY32;

	// // hModuleSnap, _ := w32.CreateToolhelp32Snapshot(0x0010 | w32.PROCESS_VM_READ | w32.PROCESS_QUERY_INFORMATION, pid)

	// // if(!w32.Module32First( hModuleSnap, &me32 )) {
  // //   fmt.Println("Module32First failed")
  // //   w32.CloseHandle( hModuleSnap )
  // // }

  // // Now walk the module list of the process,
  // // and display information about each module
  
  // // for ok := true; ok; ok = w32.Module32Next( hModuleSnap, &me32 ) {
  // //   fmt.Println("MODULE NAME",  me32.SzModule )
  // //   // _tprintf( TEXT("\n     Executable     = %s"),     me32.szExePath );
  // //   // _tprintf( TEXT("\n     Process ID     = 0x%08X"),         me32.th32ProcessID );
  // //   // _tprintf( TEXT("\n     Ref count (g)  = 0x%04X"),     me32.GlblcntUsage );
  // //   // _tprintf( TEXT("\n     Ref count (p)  = 0x%04X"),     me32.ProccntUsage );
	// // 	fmt.Println("Base address",  me32.ModBaseAddr )
  // //   // _tprintf( TEXT("\n     Base size      = %d"),             me32.modBaseSize );
  // // }

  // // w32.CloseHandle( hModuleSnap );
}

func memoryReadClose() {
  windows.CloseHandle(handle)
}

func readMemoryAt(address int64) float32 {
	var (
		data [4]byte
		length uint32
	)

	procReadProcessMemory.Call(
		uintptr(handle), 
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)), 
		uintptr(unsafe.Pointer(&length)),
	)

  bits := binary.LittleEndian.Uint32(data[:])
	float := math.Float32frombits(bits)
	return float
}

func readMemoryAtByte8(address int64) uint64 {
	var (
		data [8]byte
		length uint32
	)

	procReadProcessMemory.Call(
		uintptr(handle), 
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)), 
		uintptr(unsafe.Pointer(&length)),
	)
	
  byte8 := binary.LittleEndian.Uint64(data[:])
	return byte8
}

type staticPointer struct {
	pointer string
	offsets []string
}


func GetAddresses() (int64, int64) {
	xPositionPointer := staticPointer{"2518790", []string{"2E4", "10", "8", "8", "8", "78", "5E0"}}
	zPositionPointer := staticPointer{"2518790", []string{"2E8", "10", "8", "8", "8", "78", "5E0"}}
	xPositionAddress := calculateAddress(xPositionPointer)
	zPositionAddress := calculateAddress(zPositionPointer)
	xPositionAddressInt, _ := strconv.ParseInt(xPositionAddress, 16, 0)
	zPositionAddressInt, _ := strconv.ParseInt(zPositionAddress, 16, 0)
	return xPositionAddressInt, zPositionAddressInt
}

// * are constants
// 1. get base hex address of process
// 2. add *staticpointer* hex to base address
// 3. go to loop in array of offsets (from end to start):
// 3.1. take *offset*
// 3.2. add this offset to [value]
// 3.3. now you found the address that you have to read in order to get next address
// 3.4. !!! read [value] address and convert it with 8 bytes
// 3.5. !!! the address is the new [value]
// the "!!!" signs indicating that they should not be executed in last iteration of loop
// 4. return [value] this is the final address

func calculateAddress(pointer staticPointer) string {
	// baseAddress := "300905A4D"
	// var value string = sumHex(baseAddress, pointer.pointer)
	var value string = "20877B33300"
	for i := len(pointer.offsets)-1; i >= 0; i-- {
		offset := pointer.offsets[i]
		addressPointer := sumHex(value, offset)

		if(i > 0) {
			addressInt, _ := strconv.ParseInt(addressPointer, 16, 64)
			nextAddressDecimal := readMemoryAtByte8(addressInt)			
			value = strconv.FormatInt(int64(nextAddressDecimal), 16)
		} else {
			value = addressPointer
		}
	}
	return value
}

func sumHex(aHex string, bHex string) string {
	aDecimal, _ := strconv.ParseInt(aHex, 16, 0)
	bDecimal, _ := strconv.ParseInt(bHex, 16, 0)
	resultDecimal := aDecimal + bDecimal
	resultHex := strconv.FormatInt(resultDecimal, 16)
	return resultHex
}
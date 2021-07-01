package main

import (
	"unsafe"

	"github.com/JamesHovious/w32"
)

/*
#cgo CFLAGS: -DWIN32_LEAN_AND_MEAN
#include <Windows.h>
*/
import "C"

const (
	DllProcessAttach uint32 = 1
	DllProcessDetach uint32 = 0

	DllThreadAttach uint32 = 2
	DllThreadDetach uint32 = 3
)

//export DllMain
func DllMain(hMod unsafe.Pointer, fdwReason uint32, lpReserved unsafe.Pointer) bool {
	switch fdwReason {
	case DllProcessAttach:
		{
			// Disable thread attach / detach call.
			C.DisableThreadLibraryCalls(C.HMODULE(hMod))

			// Display a message box.
			// This will be blocking, so ideally you'd want to run it in a thread.
			w32.MessageBox(0, "Process is being attached!", "Success!", w32.MB_OK|w32.MB_ICONINFORMATION)
		}

	case DllProcessDetach:
		{
			// Display a message box.
			// This will be blocking, so ideally you'd want to run it in a thread.
			w32.MessageBox(0, "Process is being detached!", "Success!", w32.MB_OK|w32.MB_ICONINFORMATION)
		}
	}

	return true
}

func main() {}

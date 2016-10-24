package equihash

// #cgo CFLAGS: -std=c11 -D_GNU_SOURCE -Wno-pointer-sign
// #cgo LDFLAGS: -L${SRCDIR} -lsodium
// #include "src/equi/equi.h"
import "C"
import "unsafe"

func Verify(header, solution []byte) bool {
	h := C.CString(string(header))
	defer C.free(unsafe.Pointer(h))

	sol := C.CString(string(solution))
	defer C.free(unsafe.Pointer(sol))

	return (bool)(C.verify(h, sol))
}

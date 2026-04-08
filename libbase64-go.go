package main

// #include <stdbool.h>
// #include "constants.h"
// #include <stdlib.h>
import "C"
import (
	"encoding/base64"
	"runtime"
	"unsafe"
)

const ERRNO_NO_ERROR = 0
const ERRNO_GENERIC_ERROR = 1
const ERRNO_DECODE_ERROR = 2

var VERSION string
var lastErrorCode int = ERRNO_NO_ERROR

//export libbase64_go_nts__getLastErrorCode
func libbase64_go_nts__getLastErrorCode() C.int {
	result := C.int(lastErrorCode)
	lastErrorCode = ERRNO_NO_ERROR
	return result
}

//export libbase64_go_nts__getErrorDescription
func libbase64_go_nts__getErrorDescription(errno C.int) *C.char {
	result := C.CString("No errors happened")
	if (errno == ERRNO_NO_ERROR){
		// pass
	}
	if (errno == ERRNO_GENERIC_ERROR){
		result = C.CString("Some generic error happend")
	}
	if (errno == ERRNO_DECODE_ERROR){
		result = C.CString("Some decoding error happend")
	}
	return result
}

//export libbase64_go_nts__BASE64_encode_std
func libbase64_go_nts__BASE64_encode_std(inputText *C.char) *C.char {
	goText := C.GoString(inputText)
	encodedString := base64.StdEncoding.EncodeToString([]byte(goText))
	return C.CString(encodedString)
}

//export libbase64_go_nts__BASE64_decode_std
func libbase64_go_nts__BASE64_decode_std(inputText *C.char) *C.char {
	goText := C.GoString(inputText)
	decodedBytes, err := base64.StdEncoding.DecodeString(goText)
	if err != nil {
		lastErrorCode = ERRNO_DECODE_ERROR
// @todo extract error string
//		log.Fatal("decode error:", err)
		return nil
	}
	return C.CString(string(decodedBytes))
}

//export libbase64_go_nts__BASE64_encode_url
func libbase64_go_nts__BASE64_encode_url(inputText *C.char) *C.char {
	goText := C.GoString(inputText)
	encodedString := base64.URLEncoding.EncodeToString([]byte(goText))
	return C.CString(encodedString)
}

//export libbase64_go_nts__BASE64_encode_raw
func libbase64_go_nts__BASE64_encode_raw(inputText *C.char) *C.char {
	goText := C.GoString(inputText)
	encodedString := base64.RawStdEncoding.EncodeToString([]byte(goText))
	return C.CString(encodedString)
}

//export libbase64_go__FreeResult
func libbase64_go__FreeResult(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

//export libbase64_go__Version
func libbase64_go__Version() *C.char {
	return C.CString(VERSION)
}

//export libbase64_go__Version_go
func libbase64_go__Version_go() *C.char {
	return C.CString(runtime.Version())
}

func main() {}

package main

import (
	"C"
)

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"unsafe"
)

func main() {}

//export gzip_decode
func gzip_decode(crypted *C.char, length C.int) *C.char {
	buf := bytes.NewBuffer(C.GoBytes(unsafe.Pointer(crypted), length))
	gzipReader, err := gzip.NewReader(ioutil.NopCloser(buf))
	if err != nil {
		return nil
	}
	origBytes, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil
	}
	return C.CString(string(origBytes))
}

func gzip_encode(origData *C.char, length C.int) *C.char {
	buf := bytes.NewBuffer(nil)
	gzipWriter := gzip.NewWriter(buf)
	if _, err := gzipWriter.Write(C.GoBytes(unsafe.Pointer(origData), length)); err != nil {
		return nil
	}
	if err := gzipWriter.Close(); err != nil {
		return nil
	}
	return C.CString(string(buf.Bytes()))
}

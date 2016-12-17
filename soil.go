package main

/*
#cgo darwin LDFLAGS: -lSOIL -framework OpenGL -framework CoreFoundation

#include <SOIL.h>

 */
import "C"
import (
	"unsafe"
)

func SoilLoadImage(filename string) ([]byte, int, int, int) {
	image := C.CString(filename)
	var width, height, channels C.int

	chars := C.SOIL_load_image(image, &width, &height, &channels, C.SOIL_LOAD_RGB)
	n := width * height * channels

	rawdata := unsafe.Pointer(chars)
	data := C.GoBytes(rawdata, C.int(n))

	return data, int(width), int(height), int(channels)
}
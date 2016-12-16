package main

/*
#cgo darwin LDFLAGS: -lSOIL -framework OpenGL -framework CoreFoundation

#include <SOIL.h>

 */
import "C"
import (
	"fmt"
	"unsafe"
)

func SoilLoadImage() ([]byte, int, int) {
	image := C.CString("container.jpg")
	var w, h, p C.int

	chars := C.SOIL_load_image(image, &w, &h, &p, C.SOIL_LOAD_RGB)
	n := w * h * p
	fmt.Printf("w: %d, h: %d, p: %d\n", w, h, p)

	rawdata := unsafe.Pointer(chars)
	data := C.GoBytes(rawdata, C.int(n))

	fmt.Println(len(data))
	fmt.Println(chars == nil)


	return data, int(w), int(h)
}
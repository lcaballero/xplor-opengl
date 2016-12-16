package main

/*
#cgo LDFLAGS: -lglew

#include <GL/glew.h>

void TurnOnExperimental() {
	glewExperimental = GL_TRUE;
}

void InitGlew() {
	glewInit();
}
*/
import "C"

func TurnOnExperimental() {
	C.TurnOnExperimental()
}
func InitGlew() {
	C.InitGlew()
}
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

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const DefaultTitle = "Go :: OS X - OpenGL"

type Window struct {
	window *glfw.Window
	Title string
}

/*
NVIDIA GeForce GT 750M 2048 MB graphics
*/
func NewWindow() *Window {
	w := &Window{}
	return w
}

func (win *Window) Create() error {
	fmt.Println("Starting GLFW context, OpenGL 4.1")

	err := glfw.Init()
	if err != nil {
		fmt.Println("Failed to init GLFW")
		return err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	if win.Title == "" {
		win.Title = DefaultTitle
	}

	win.window, err = glfw.CreateWindow(640, 480, win.Title, nil, nil)
	if err != nil {
		return err
	}

	err = gl.Init()
	if err != nil {
		return err
	}
	win.window.MakeContextCurrent()
	win.window.SetKeyCallback(closeWindow)

	C.TurnOnExperimental()
	C.InitGlew()

	fmt.Printf("opengl version: '%s'\n", gl.GoStr(gl.GetString(gl.VERSION)))

	return nil
}

func (win *Window) ViewportToFramebufferSize() {
	w, h := win.FramebufferSize()
	fmt.Printf("width: %d, height: %d\n", w, h)
	gl.Viewport(0, 0, w, h)
}

func (w *Window) FramebufferSize() (int32, int32) {
	width, height := w.window.GetFramebufferSize()
	return int32(width), int32(height)
}

func (w *Window) ShouldClose() bool {
	return w.window.ShouldClose()
}

func (w *Window) SwapBuffers() {
	w.window.SwapBuffers()
}

func (w *Window) TerminateWindowing() {
	glfw.Terminate()
}

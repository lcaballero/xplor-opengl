package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"
	"runtime"
	"fmt"
)

func init() {
	runtime.LockOSThread()
}

const title = "Go :: OS X - OpenGL"

/*
NVIDIA GeForce GT 750M 2048 MB graphics

 */


func main() {
	err := glfw.Init()
	if err != nil {
		fmt.Println("Failed to init GLFW")
		panic(err)
	}

	defer glfw.Terminate()

	fmt.Println("here")

	glfw.WindowHint(glfw.Resizable, glfw.False)
	fmt.Println("resize")
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	fmt.Println("major")
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	fmt.Println("minor")
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	fmt.Println("gl profile")
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	fmt.Println("gl forward compat")

	window, err := glfw.CreateWindow(640, 480, title, nil, nil)
	if err != nil {
		panic(err)
	}

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	fmt.Printf("opengl version: '%s'\n", gl.GoStr(gl.GetString(gl.VERSION)))

	window.MakeContextCurrent()
	window.Show()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

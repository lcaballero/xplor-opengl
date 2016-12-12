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
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"
	"runtime"
	"fmt"
	"strings"
)


func init() {
	runtime.LockOSThread()
}

const title = "Go :: OS X - OpenGL"

func closeWindow(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if (key == glfw.KeyEscape && action == glfw.Press) {
		w.SetShouldClose(true)
	}
}

/*
NVIDIA GeForce GT 750M 2048 MB graphics

 */
func main() {
	fmt.Println("Starting GLFW context, OpenGL 4.1")

	err := glfw.Init()
	if err != nil {
		fmt.Println("Failed to init GLFW")
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window := createWindow()
	window.MakeContextCurrent()
	window.SetKeyCallback(closeWindow)

	C.TurnOnExperimental()
	C.InitGlew()

	fmt.Printf("opengl version: '%s'\n", gl.GoStr(gl.GetString(gl.VERSION)))

	width, height := window.GetFramebufferSize()
	fmt.Printf("width: %d, height: %d\n", width, height)

	gl.Viewport(0, 0, int32(width), int32(height))

	vs := NewVertexShaderCompiler(vertex_shader_source)
	err = vs.CreateAndCompile()
	if err != nil {
		panic(err)
	}

	fs := NewFragmentShaderCompiler(fragment_shader_source)
	err = fs.CreateAndCompile()
	if err != nil {
		panic(err)
	}

	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vs.ID)
	gl.AttachShader(shaderProgram, fs.ID)
	gl.LinkProgram(shaderProgram)
	checkLinking(shaderProgram, "shader program")

	vs.Delete()
	fs.Delete()

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	var vbo, vao uint32
	gl.GenVertexArrays(1, &vao)

	gl.GenBuffers(1, &vbo)
	defer gl.DeleteBuffers(1, &vbo)

	gl.BindVertexArray(vao)
	defer gl.DeleteVertexArrays(1, &vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3 * 4, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	for !window.ShouldClose() {
		glfw.PollEvents()

		gl.ClearColor(0.0, 0.0, 0.4, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
 		gl.BindVertexArray(0)

		window.SwapBuffers()
	}
}

func createWindow() *glfw.Window {
	window, err := glfw.CreateWindow(640, 480, title, nil, nil)
	if err != nil {
		panic(err)
	}

	err = gl.Init()
	if err != nil {
		panic(err)
	}
	return window
}

func checkLinking(program uint32, msg string) {
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		logLength := int32(512)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))
		fmt.Println(log)
		panic("didn't compile shader")
	} else {
		fmt.Println("successfully linked", msg)
	}
}

var vertex_shader_source = `
#version 410 core
layout (location = 0) in vec3 position;

void main()
{
	gl_Position = vec4(position.x, position.y, position.z, 1.0);
}
` + "\x00"


var fragment_shader_source = `
#version 410 core

out vec4 color;

void main()
{
	color = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
` + "\x00"
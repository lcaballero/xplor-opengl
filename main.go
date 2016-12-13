package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
	"math"
)

func init() {
	runtime.LockOSThread()
}

func closeWindow(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func main() {
	window := NewWindow()
	err := window.Create()
	if err != nil {
		panic(err)
	}
	defer window.TerminateWindowing()

	window.ViewportToFramebufferSize()

	vs := NewVertexCompiler("shaders/shader.vert")
	fs := NewFragmentCompiler("shaders/shader.frag")

	p := NewShaderProgram()
	err = p.Attach(vs, fs)
	if err != nil {
		panic(err)
	}
	p.DeleteShaders()

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	var vbo, vao uint32
	gl.GenBuffers(1, &vbo)
	defer gl.DeleteBuffers(1, &vbo)

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	defer gl.DeleteVertexArrays(1, &vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	for !window.ShouldClose() {
		glfw.PollEvents()

		gl.ClearColor(0.0, 0.0, 0.4, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		t := glfw.GetTime()
		g := (math.Sin(t) / 2.0) + 0.5
		greenValue := float32(g)
		varName := gl.Str("ourColor\x00")
		colorLocation := gl.GetUniformLocation(p.GetID(), varName)

		p.UseProgram()
		gl.Uniform4f(colorLocation, 0.0, greenValue, 0.0, 1.0)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.BindVertexArray(0)

		window.SwapBuffers()
	}
}

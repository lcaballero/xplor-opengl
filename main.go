package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
	"math"
	"time"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

func transform(rads, scaleXYZ float32) mgl32.Mat4 {
	rot := mgl32.HomogRotate3DZ(rads)
	scale := mgl32.Scale3D(scaleXYZ, scaleXYZ, scaleXYZ)
	trans := scale.Mul4(rot)
	return trans
}

func main() {
	run()
}

func run() {
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
		// Positions      // Colors        // Texture Coords
		0.5,   0.5, 0.0,  1.0, 0.0, 0.0,   1.0, 1.0,  // Top Right
		0.5,  -0.5, 0.0,  0.0, 1.0, 0.0,   1.0, 0.0,  // Bottom Right
		-0.5, -0.5, 0.0,  0.0, 0.0, 1.0,   0.0, 0.0,  // Bottom Left
		-0.5,  0.5, 0.0,  1.0, 1.0, 0.0,   0.0, 1.0,  // Top Left
	}

	indices := []uint32{
		0, 1, 3,  // First Triangle
		1, 2, 3,  // Second Triangle
	}

	texture := LoadTexture("container.jpg")

	var vbo, vao, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	defer gl.DeleteBuffers(1, &vbo)
	defer gl.DeleteBuffers(1, &ebo)
	defer gl.DeleteVertexArrays(1, &vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)

	gl.BindVertexArray(0)

	tic := time.NewTicker(1 * time.Second)
	frames := 0
	var frameRot float32 = math.Pi / 60.0
	var currRot float32 = 0.0

	for !window.ShouldClose() {
		select {
		case <- tic.C:
			fmt.Printf("frames %d / sec\n", frames)
			frames = 0
		default:
			frames++
			glfw.PollEvents()

			gl.ClearColor(0.2, 0.3, 0.3, 1.0)
			gl.Clear(gl.COLOR_BUFFER_BIT)

			p.UseProgram()

			gl.ActiveTexture(gl.TEXTURE0)
			gl.BindTexture(gl.TEXTURE_2D, texture.Id)
			gl.Uniform1i(gl.GetUniformLocation(p.GetID(), gl.Str("ourTexture\x00")), 0)

			trans := transform(currRot, 0.5)
			currRot += frameRot
			loc := gl.GetUniformLocation(p.GetID(), gl.Str("transform\x00"))
			gl.UniformMatrix4fv(loc, 1, false, &trans[0])

			gl.BindVertexArray(vao)
			gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
			gl.BindVertexArray(0)

			window.SwapBuffers()
		}
	}
}

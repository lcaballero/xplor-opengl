package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
	"time"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

const Pi = math.Pi

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

	p, err := NewProgram("shaders/shader.vert", "shaders/shader.frag")
	if err != nil {
		panic(err)
	}

	var vbo, vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)

	defer gl.DeleteBuffers(1, &vbo)
	defer gl.DeleteVertexArrays(1, &vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(2)

	gl.BindVertexArray(0)

	texture, err := LoadTexture("container.jpg")
	if err != nil {
		panic(err)
	}

	fmt.Println(texture.String())

	tic := time.NewTicker(1 * time.Second)
	frames := 0
	//cubePos := cubePositions[0]

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
			gl.Uniform1i(p.UniformLocation("ourTexture"), 0)

			transformCube(p)

//			angle := 45.0
			model := mgl32.Ident4()
//				Mul4(mgl32.Translate3D(0.0, 0.0, 0.0)).
//				Mul4(mgl32.HomogRotate3DZ(angle))
			modelLoc := p.UniformLocation("model")
			gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])
			gl.DrawArrays(gl.TRIANGLES, 0, 36)

			gl.BindVertexArray(vao)
			gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
			gl.BindVertexArray(0)

			window.SwapBuffers()
		}
	}
}

func transformCube(p *ShaderProgram) {
	view := mgl32.Ident4()
	projection := mgl32.Perspective(45.0, float32(Width)/float32(Height), 0.1, 100.0)
	view = view.Mul4(mgl32.Translate3D(0.0, 0.0, -0.1))

	viewLoc := p.UniformLocation("view")
	projLoc := p.UniformLocation("projection")

	gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])
	gl.UniformMatrix4fv(projLoc, 1, false, &projection[0])
}

package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
	"math"
	"time"
	"fmt"
	"os"
	"golang.org/x/image/bmp"
	"bytes"
	"bufio"
	"io/ioutil"
	"image/jpeg"
	"image/color"
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
		// Positions      // Colors        // Texture Coords
		 0.5,  0.5, 0.0,  1.0, 0.0, 0.0,   1.0, 1.0,  // Top Right
		 0.5, -0.5, 0.0,  0.0, 1.0, 0.0,   1.0, 0.0,  // Bottom Right
		-0.5, -0.5, 0.0,  0.0, 0.0, 1.0,   0.0, 0.0,  // Bottom Left
		-0.5,  0.5, 0.0,  1.0, 1.0, 0.0,   0.0, 1.0,  // Top Left
	}

	indices := []uint32{
		0, 1, 3,  // First Triangle
		1, 2, 3,  // Second Triangle
	}

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

	imageBytes, w, h := SoilLoadImage()
//	imageBytes, w, h := LoadImage("container.bmp")
	fmt.Printf("texture w: %d, h: %d, len: %d\n", w, h, len(imageBytes))

	var texture0 uint32
	gl.GenTextures(1, &texture0)
	gl.BindTexture(gl.TEXTURE_2D, texture0)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, int32(w), int32(h), 0, gl.RGB, gl.UNSIGNED_BYTE, gl.Ptr(imageBytes))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	tic := time.NewTicker(1 * time.Second)
	frames := 0

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
			gl.BindTexture(gl.TEXTURE_2D, texture0)
			gl.Uniform1i(gl.GetUniformLocation(p.GetID(), gl.Str("ourTexture\x00")), 0)

			gl.BindVertexArray(vao)
			gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
			gl.BindVertexArray(0)

			window.SwapBuffers()
		}
	}
}

func GreenValue() float32 {
	t := glfw.GetTime()
	g := (math.Sin(t) / 2.0) + 0.5
	greenValue := float32(g)
	return greenValue
}

func LoadImage(filename string) ([]byte, int, int) {
	bin, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bin, 512, 512
}

func LoadTextureMap(filename string) ([]byte, int, int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("open")
		panic(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		fmt.Println("decode")
		panic(err)
	}

	m := img.Bounds()
	x, y, w, h := m.Min.X, m.Min.Y, m.Max.X, m.Max.Y
	size := w * h
	dim := size * 4 * 3
	fmt.Println(x, y, w, h, size, dim)
	data := make([]byte, dim)

	for i := 0; i < dim; i+=12 {
		if len(data) > dim {
			fmt.Printf("x: %d, y: %d", x, y)
		}
		pixColor := img.At(x, y)
		r, g, b, _ := color.NRGBA64Model.Convert(pixColor).RGBA()
		r1, r2, r3, r4 := r >> 12, r >> 8, r >> 4, r & 0x000f
		g1, g2, g3, g4 := g >> 12, g >> 8, g >> 4, g & 0x000f
		b1, b2, b3, b4 := b >> 12, b >> 8, b >> 4, b & 0x000f

//		fmt.Printf("<%x, %x, %x, %x, %x, %x>\n", r1, r2, g1, g2, b1, b2)
//		fmt.Printf("red is greater than 255: (%d, %d, %d, %d), <%x, %x, %x, %x>\n", r, g, b, a, r, g, b, a)

		data[i+0] = byte(r1 & 0x00f)
		data[i+1] = byte(r2 & 0x00f)
		data[i+2] = byte(r3 & 0x00f)
		data[i+3] = byte(r4 & 0x00f)

		data[i+4] = byte(g1 & 0x00f)
		data[i+5] = byte(g2 & 0x00f)
		data[i+6] = byte(g3 & 0x00f)
		data[i+7] = byte(g4 & 0x00f)

		data[i+8] = byte(b1 & 0x00f)
		data[i+9] = byte(b2 & 0x00f)
		data[i+10] = byte(b3 & 0x00f)
		data[i+11] = byte(b4 & 0x00f)

		x++
		if x >= w {
			y++
			x = 0
		}
	}
	return data, w, h
}

func LoadTexture(filename string) ([]byte, int, int) {
	f, err := os.Open(filename + ".jpg")
	if err != nil {
		fmt.Println("open")
		panic(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		fmt.Println("decode")
		panic(err)
	}

	w, h := img.Bounds().Max.X, img.Bounds().Max.Y

	bb := bytes.Buffer{}
	buf := bufio.NewWriter(&bb)

	err = bmp.Encode(buf, img)
	if err != nil {
		fmt.Println("encode")
		panic(err)
	}

	imageBytes := bb.Bytes()

	return imageBytes, w, h
}
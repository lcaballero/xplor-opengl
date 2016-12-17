package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
)


type Texture struct {
	Filename string
	Id uint32
	width int
	height int
	channels int
	data []byte
}

func LoadTexture(filename string) *Texture {
	imageBytes, w, h, c := SoilLoadImage(filename)
	fmt.Printf("texture w: %d, h: %d, len: %d\n", w, h, len(imageBytes))

	var textureID uint32
	gl.GenTextures(1, &textureID)
	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGB,
		int32(w), int32(h),
		0, gl.RGB, gl.UNSIGNED_BYTE, gl.Ptr(imageBytes))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	t := &Texture{
		Filename: filename,
		Id: textureID,
		data: imageBytes,
		width: w,
		height: h,
		channels: c,
	}

	return t
}

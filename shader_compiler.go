package main

import (
	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
	"fmt"
)

var ErrNotImplementedYet = errors.New("err not implemented yet")

type ShaderCompiler struct {
	ID uint32
	kind uint32
	source string
	msg string
}

func NewVertexShaderCompiler(source string) *ShaderCompiler {
	sc := &ShaderCompiler{
		source: source,
		kind: gl.VERTEX_SHADER,
		msg: "vertex shader",
	}
	return sc
}

func NewFragmentShaderCompiler(source string) *ShaderCompiler {
	sc := &ShaderCompiler{
		source: source,
		kind: gl.FRAGMENT_SHADER,
		msg: "fragment shader",
	}
	return sc
}

func (sc *ShaderCompiler) Delete() {
	gl.DeleteShader(sc.ID)
}

func (sc *ShaderCompiler) CreateAndCompile() error {
	sc.ID = gl.CreateShader(sc.kind)
	shaderBytes, free := gl.Strs(sc.source)
	defer free()

	gl.ShaderSource(sc.ID, 1, shaderBytes, nil)
	gl.CompileShader(sc.ID)
	return sc.checkCompilation(sc.ID, sc.msg)
}

func (sc *ShaderCompiler) checkCompilation(id uint32, msg string) error {
	var status int32
	gl.GetShaderiv(id, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		logLength := int32(512)
		space := strings.Repeat("\x00", int(logLength+1))
		logPtr := gl.Str(space)
		gl.GetShaderInfoLog(id, logLength, nil, logPtr)
		log := gl.GoStr(logPtr)
		return fmt.Errorf("shader compilation error: %s", log)
	} else {
		fmt.Printf("successfully compiled: %s\n", msg)
		return nil
	}
}


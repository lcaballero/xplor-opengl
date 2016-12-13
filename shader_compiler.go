package main

import (
	"errors"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
	"io/ioutil"
)

var ErrNotImplementedYet = errors.New("err not implemented yet")

type Shader interface {
	GetID() uint32
	Compile() error
	Delete()
}

type shader struct {
	id uint32
	source string
	msg    string
}

func NewVertexCompiler(filename string) Shader {
	sc := &shader{
		id:     gl.CreateShader(gl.VERTEX_SHADER),
		source: MustRead(filename),
		msg:    "vertex shader",
	}
	return sc
}

func NewFragmentCompiler(filename string) Shader {
	sc := &shader{
		id: gl.CreateShader(gl.FRAGMENT_SHADER),
		source: MustRead(filename),
		msg:    "fragment shader",
	}
	return sc
}

func MustRead(filename string) string {
	bin, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	bin = append(bin, byte(0))
	return string(bin)
}


func (sc *shader) GetID() uint32 {
	return sc.id
}

func (sc *shader) Delete() {
	gl.DeleteShader(sc.id)
}

func (sc *shader) Compile() error {
	shaderBytes, free := gl.Strs(sc.source)
	defer free()

	gl.ShaderSource(sc.id, 1, shaderBytes, nil)
	gl.CompileShader(sc.id)
	return sc.checkCompilation(sc.id, sc.msg)
}

func (sc *shader) checkCompilation(id uint32, msg string) error {
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

package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type ShaderProgram struct {
	id uint32
	shaders []Shader
}

func NewShaderProgram() *ShaderProgram {
	return &ShaderProgram{}
}

func (p *ShaderProgram) GetID() uint32 {
	return p.id
}

func (p *ShaderProgram) Attach(shaders ...Shader) error {
	p.shaders = shaders
	for _, s := range shaders {
		err := s.Compile()
		if err != nil {
			return err
		}
	}
	p.id = gl.CreateProgram()
	for _, s := range shaders {
		gl.AttachShader(p.id, s.GetID())
	}
	gl.LinkProgram(p.id)

	return p.checkLinking()
}

func (p *ShaderProgram) UseProgram() {
	gl.UseProgram(p.id)
}

func (p *ShaderProgram) DeleteShaders() {
	for _, s := range p.shaders {
		s.Delete()
	}
	p.shaders = make([]Shader, 0)
}

func (p *ShaderProgram) checkLinking() error {
	var status int32
	gl.GetProgramiv(p.id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		logLength := int32(512)
		space := strings.Repeat("\x00", int(logLength+1))
		logPtr := gl.Str(space)
		gl.GetProgramInfoLog(p.id, logLength, nil, logPtr)
		log := gl.GoStr(logPtr)
		return fmt.Errorf("program linking error: %s", log)
	} else {
		fmt.Println("successfully linked: shader programs")
		return nil
	}
}


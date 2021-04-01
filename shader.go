package OctaForce

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"io/ioutil"
	"log"
	"strings"
)

const (
	ShaderFragment = 1
	ShaderVertex   = 2
)

type Shader struct {
	path   string
	typ    uint32
	shader uint32
}

func NewShader(path string, shaderTyp int) *Shader {
	var typ uint32
	switch shaderTyp {
	case ShaderFragment:
		typ = gl.FRAGMENT_SHADER
	case ShaderVertex:
		typ = gl.VERTEX_SHADER
	}
	shader := &Shader{
		path: path,
		typ:  typ,
	}

	s, err := compileShader(shader.path, shader.typ)
	if err != nil {
		panic(err)
	}
	shader.shader = s
	return shader
}

func compileShader(path string, typ uint32) (uint32, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	shader := gl.CreateShader(typ)

	sources := string(content) + "\x00"
	cSources, free := gl.Strs(sources)
	gl.ShaderSource(shader, 1, cSources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logString := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logString))

		return 0, fmt.Errorf("failed to compile \n %v \n%v", string(content), logString)
	}

	return shader, nil
}

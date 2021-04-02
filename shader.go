package OctaForce

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"io/ioutil"
	"log"
	"strings"
)

var shaders []uint32

func initShaders() {
	shaders = []uint32{
		newShader("/shader/guiVertexShader.shader", gl.VERTEX_SHADER),
		newShader("/shader/guiFragmentShader.shader", gl.FRAGMENT_SHADER),

		newShader("/shader/vertexShader.shader", gl.VERTEX_SHADER),
		newShader("/shader/fragmentShader.shader", gl.FRAGMENT_SHADER),
	}
}

func newShader(localPath string, shaderTyp uint32) uint32 {
	content, err := ioutil.ReadFile(absPath + localPath)
	if err != nil {
		log.Fatal(err)
	}

	shader := gl.CreateShader(shaderTyp)

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

		panic(fmt.Errorf("failed to compile \n %v \n%v", string(content), logString))
	}
	return shader
}

func deleteShaders() {
	for _, shader := range shaders {
		gl.DeleteShader(shader)
	}
}

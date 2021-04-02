package OctaForce

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func initOpenGL() {
	err := gl.Init()
	if err != nil {
		panic(err)
	}

	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.DEPTH_TEST)

}

func preRender(clearColor [3]float32) {
	gl.ClearColor(clearColor[0], clearColor[1], clearColor[2], 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func render3D() {
	view := activeCamera.Transform.getMatrix().Inv()
	gl.UniformMatrix4fv(1, 1, false, &view[0])
	gl.UniformMatrix4fv(0, 1, false, &activeCamera.projection[0])

	for _, base := range materialBases {
		base.renderBase()
	}
}

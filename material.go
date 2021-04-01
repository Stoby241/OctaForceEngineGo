package OctaForce

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type IMaterial interface {
	render()
}

// Material is a Struct with is needed by the Mesh Component to set the Color of an Mesh.
type materialBase struct {
	vertexShader   *Shader
	fragmentShader *Shader
	programm       uint32

	meshes []*Mesh
}

type MaterialSimple struct {
	materialBase
	Color mgl32.Vec4
}

func NewMaterialSimple(color mgl32.Vec4) *MaterialSimple {
	return &MaterialSimple{Color: color}
}

func (m *MaterialSimple) render() {
	gl.UseProgram(m.programm)

	gl.Uniform3f(2, m.Color[0], m.Color[1], m.Color[2])

	for _, mesh := range globalActiveMeshesData.meshes {
		gl.BindVertexArray(mesh.vao)

		if mesh.needsVertexUpdate {
			pushVertexData(mesh)
		}

		if len(mesh.instances) == 0 {
			pushTransformData(mesh)
			gl.DrawElements(gl.TRIANGLES, int32(len(mesh.Indices)), gl.UNSIGNED_INT, nil)

		} else {
			pushTransformDataInstanced(mesh)
			gl.DrawElementsInstanced(gl.TRIANGLES, int32(len(mesh.Indices)), gl.UNSIGNED_INT, nil, int32(len(mesh.instances)+1))
		}
	}
}

const vertexStride int = 3 * 4
const transformStride int = 16 * 4

func pushVertexData(mesh *Mesh) {
	var vertexData []float32
	for _, vertex := range mesh.Vertices {
		vertexData = append(vertexData, []float32{
			vertex.Position.X(),
			vertex.Position.Y(),
			vertex.Position.Z(),
		}...)
	}

	// vertex VBO
	gl.GenBuffers(1, &mesh.vertexVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.vertexVBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData)*4, gl.Ptr(vertexData), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, int32(vertexStride), gl.PtrOffset(0))

	// EBO
	gl.GenBuffers(1, &mesh.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, mesh.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(mesh.Indices)*4, gl.Ptr(mesh.Indices), gl.STATIC_DRAW)

	mesh.needsVertexUpdate = false
}

func pushTransformData(mesh *Mesh) {
	if mesh.needsTransformUpdate {
		gl.GenBuffers(1, &mesh.transformVBO)
		gl.BindBuffer(gl.ARRAY_BUFFER, mesh.transformVBO)
		gl.BufferData(gl.ARRAY_BUFFER, transformStride, gl.Ptr(nil), gl.DYNAMIC_DRAW)

		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(1, 4, gl.FLOAT, false, int32(transformStride), gl.PtrOffset(0))
		gl.VertexAttribDivisor(1, 1)

		gl.EnableVertexAttribArray(2)
		gl.VertexAttribPointer(2, 4, gl.FLOAT, false, int32(transformStride), gl.PtrOffset(3*4))
		gl.VertexAttribDivisor(2, 1)

		gl.EnableVertexAttribArray(3)
		gl.VertexAttribPointer(3, 4, gl.FLOAT, false, int32(transformStride), gl.PtrOffset(7*4))
		gl.VertexAttribDivisor(3, 1)

		gl.EnableVertexAttribArray(4)
		gl.VertexAttribPointer(4, 4, gl.FLOAT, false, int32(transformStride), gl.PtrOffset(11*4))
		gl.VertexAttribDivisor(4, 1)

		mesh.needsTransformUpdate = false
	}
	gl.BufferData(gl.ARRAY_BUFFER, transformStride, gl.Ptr(mesh.Transform.matrix), gl.DYNAMIC_DRAW)
}
func pushTransformDataInstanced(mesh *Mesh) {
	if mesh.needsTransformUpdate {
		gl.GenBuffers(1, &mesh.transformVBO)
		gl.BindBuffer(gl.ARRAY_BUFFER, mesh.transformVBO)
		gl.BufferData(gl.ARRAY_BUFFER, (1+len(mesh.instances))*transformStride, gl.Ptr(nil), gl.DYNAMIC_DRAW)

		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(1, 4, gl.FLOAT, false, int32(transformStride), gl.PtrOffset(0))
		gl.VertexAttribDivisor(1, 1)

		gl.EnableVertexAttribArray(2)
		gl.VertexAttribPointer(2, 4, gl.FLOAT, false, int32(transformStride), gl.PtrOffset(3*4))
		gl.VertexAttribDivisor(2, 1)

		gl.EnableVertexAttribArray(3)
		gl.VertexAttribPointer(3, 4, gl.FLOAT, false, int32(transformStride), gl.PtrOffset(7*4))
		gl.VertexAttribDivisor(3, 1)

		gl.EnableVertexAttribArray(4)
		gl.VertexAttribPointer(4, 4, gl.FLOAT, false, int32(transformStride), gl.PtrOffset(11*4))
		gl.VertexAttribDivisor(4, 1)

		mesh.instanceTransforms = make([]float32, (1+len(mesh.instances))*16)

		mesh.needsTransformUpdate = false
	}

	mesh.instanceTransforms[0] = mesh.Transform.matrix[0]
	mesh.instanceTransforms[1] = mesh.Transform.matrix[1]
	mesh.instanceTransforms[2] = mesh.Transform.matrix[2]
	mesh.instanceTransforms[3] = mesh.Transform.matrix[3]

	mesh.instanceTransforms[4] = mesh.Transform.matrix[4]
	mesh.instanceTransforms[5] = mesh.Transform.matrix[5]
	mesh.instanceTransforms[6] = mesh.Transform.matrix[6]
	mesh.instanceTransforms[7] = mesh.Transform.matrix[7]

	mesh.instanceTransforms[8] = mesh.Transform.matrix[8]
	mesh.instanceTransforms[9] = mesh.Transform.matrix[9]
	mesh.instanceTransforms[10] = mesh.Transform.matrix[10]
	mesh.instanceTransforms[11] = mesh.Transform.matrix[11]

	mesh.instanceTransforms[12] = mesh.Transform.matrix[12]
	mesh.instanceTransforms[13] = mesh.Transform.matrix[13]
	mesh.instanceTransforms[14] = mesh.Transform.matrix[14]
	mesh.instanceTransforms[15] = mesh.Transform.matrix[15]

	for i, meshInstant := range mesh.instances {
		index := (1 + i) * 16

		mesh.instanceTransforms[index+0] = meshInstant.Transform.matrix[0]
		mesh.instanceTransforms[index+1] = meshInstant.Transform.matrix[1]
		mesh.instanceTransforms[index+2] = meshInstant.Transform.matrix[2]
		mesh.instanceTransforms[index+3] = meshInstant.Transform.matrix[3]

		mesh.instanceTransforms[index+4] = meshInstant.Transform.matrix[4]
		mesh.instanceTransforms[index+5] = meshInstant.Transform.matrix[5]
		mesh.instanceTransforms[index+6] = meshInstant.Transform.matrix[6]
		mesh.instanceTransforms[index+7] = meshInstant.Transform.matrix[7]

		mesh.instanceTransforms[index+8] = meshInstant.Transform.matrix[8]
		mesh.instanceTransforms[index+9] = meshInstant.Transform.matrix[9]
		mesh.instanceTransforms[index+10] = meshInstant.Transform.matrix[10]
		mesh.instanceTransforms[index+11] = meshInstant.Transform.matrix[11]

		mesh.instanceTransforms[index+12] = meshInstant.Transform.matrix[12]
		mesh.instanceTransforms[index+13] = meshInstant.Transform.matrix[13]
		mesh.instanceTransforms[index+14] = meshInstant.Transform.matrix[14]
		mesh.instanceTransforms[index+15] = meshInstant.Transform.matrix[15]
	}

	gl.BufferData(gl.ARRAY_BUFFER, (len(mesh.instances)+1)*transformStride, gl.Ptr(mesh.instanceTransforms), gl.DYNAMIC_DRAW)
}

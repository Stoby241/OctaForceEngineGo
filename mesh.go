package OctaForce

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type vertex struct {
	Position mgl32.Vec3
	Normal   mgl32.Vec3
	UVCord   mgl32.Vec2
}

// Mesh holds all data needed to render a 3D Object.
// When you change Vertices or Indices buy your self don't forget to set needsVertexUpdate to true. Otherwise the changes
// will not be applied.
type Mesh struct {
	Vertices          []vertex
	Indices           []uint32
	vao               uint32
	vertexVBO         uint32
	needsVertexUpdate bool

	ebo                  uint32
	transformVBO         uint32
	needsTransformUpdate bool

	Material IMaterial

	instances          []*MeshInstant
	instanceTransforms []float32

	Transform *Transform
}

func NewMesh() *Mesh {
	return &Mesh{Transform: NewTransform()}
}

// LoadOBJ returns the Mesh struct of the given OBJ file.
func (m *Mesh) LoadOBJ(path string, loadMaterials bool) {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	var vertices []mgl32.Vec3
	var normals []mgl32.Vec3
	var uvCord []mgl32.Vec2
	var faces [][3][3]uint32
	for _, line := range lines {
		values := strings.Split(line, " ")
		values[len(values)-1] = strings.Replace(values[len(values)-1], "\r", "", 1)

		switch values[0] {
		case "mtllib":
			break
		case "v":
			vertices = append(vertices, mgl32.Vec3{ParseFloat(values[1]), ParseFloat(values[2]), ParseFloat(values[3])})
			break
		case "vn":
			normals = append(normals, mgl32.Vec3{ParseFloat(values[1]), ParseFloat(values[2]), ParseFloat(values[3])})
			break
		case "vt":
			uvCord = append(uvCord, mgl32.Vec2{ParseFloat(values[1]), ParseFloat(values[2])})
			break
		case "f":
			var face [3][3]uint32
			for j, value := range values {
				if j == 0 {
					continue
				}

				number := strings.Split(value, "/")
				face[j-1][0] = uint32(ParseInt(number[0]))
				face[j-1][1] = uint32(ParseInt(number[1]))
				face[j-1][2] = uint32(ParseInt(number[2]))
			}
			faces = append(faces, face)
			break
		}
	}

	m.Vertices = make([]vertex, len(vertices))
	m.Material = NewMaterialSimple(mgl32.Vec4{1, 1, 1, 1})
	for _, face := range faces {
		for _, values := range face {
			vertexIndex := values[0] - 1
			m.Indices = append(m.Indices, vertexIndex)
			//goland:noinspection GoNilness
			m.Vertices[vertexIndex].Position = vertices[vertexIndex]
			//Mesh.Vertices[vertexIndex].UVCord = uvCord[values[1] -1]
			//Mesh.Vertices[vertexIndex].Normal = normals[values[2] -1]
		}
	}

	m.needsVertexUpdate = true
}

type activeMeshesData struct {
	meshes []*Mesh
}

func initActiveMeshesData() {
	globalActiveMeshesData = &activeMeshesData{}
}

var globalActiveMeshesData *activeMeshesData

func GetActiveMeshes() *activeMeshesData {
	return globalActiveMeshesData
}

func (a *activeMeshesData) AddMesh(mesh *Mesh) {
	gl.GenVertexArrays(1, &mesh.vao)
	a.meshes = append(a.meshes, mesh)
}
func (a *activeMeshesData) RemoveMesh(mesh *Mesh) {

	for i := len(a.meshes) - 1; i >= 0; i-- {
		if a.meshes[i] == mesh {
			a.meshes = append(a.meshes[:i], a.meshes[i+1:]...)
		}
	}

	//unUsedVAOs = append(unUsedVAOs, mesh.vao)
}

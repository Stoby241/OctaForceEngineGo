package test

import (
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"math"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)
import of "OctaForceEngineGo"

var absPath string

func init() {
	_, b, _, _ := runtime.Caller(0)
	absPath = filepath.Dir(b)
}

func TestOctaForce(t *testing.T) {
	of.Init(start)
}

const (
	movementSpeed float32 = 100
	mouseSpeed    float32 = 3
)

func start() {
	camera := of.NewCamera()

	camera.Transform = of.NewTransform()
	camera.Transform.SetPosition(mgl32.Vec3{0, 0, 200})

	of.SetActiveCamera(camera)

	task := of.NewTask(func() {

		deltaTime := float32(of.GetDeltaTime())
		if of.KeyPressed(of.KeyW) {
			camera.Transform.MoveRelative(mgl32.Vec3{0, 0, -1}.Mul(deltaTime * movementSpeed))
		}
		if of.KeyPressed(of.KeyS) {
			camera.Transform.MoveRelative(mgl32.Vec3{0, 0, 1}.Mul(deltaTime * movementSpeed))
		}
		if of.KeyPressed(of.KeyA) {
			camera.Transform.MoveRelative(mgl32.Vec3{-1, 0, 0}.Mul(deltaTime * movementSpeed))
		}
		if of.KeyPressed(of.KeyD) {
			camera.Transform.MoveRelative(mgl32.Vec3{1, 0, 0}.Mul(deltaTime * movementSpeed))
		}
		if of.MouseButtonPressed(of.MouseButtonLeft) {
			mouseMovement := of.GetMouseMovement()
			camera.Transform.Rotate(mgl32.Vec3{-1, 0, 0}.Mul(mouseMovement.Y() * deltaTime * mouseSpeed))
			camera.Transform.Rotate(mgl32.Vec3{0, -1, 0}.Mul(mouseMovement.X() * deltaTime * mouseSpeed))
		}
	})
	task.SetRepeating(true)
	task.SetRaceTask(of.GetEngineTask(of.WindowUpdateTask), of.GetEngineTask(of.RenderTask))
	of.AddTask(task)

	mesh := of.NewMesh()
	mesh.LoadOBJ(absPath+"/mesh/LowPolySphere.obj", false)
	mesh.Material = of.Material{DiffuseColor: [3]float32{1, 1, 1}}
	of.GetActiveMeshes().AddMesh(mesh)

	for i := 0; i < 1000; i++ {
		var instants []*of.MeshInstant
		for j := 0; j < 1000; j++ {
			meshInstant := of.NewMeshInstant(mesh, &of.Material{DiffuseColor: [3]float32{1, 0, 1}})
			meshInstant.Transform.SetPosition(mgl32.Vec3{float32(i) * 10, float32(j) * 10, 0})
			instants = append(instants, meshInstant)
		}

		newTask(instants)
	}
}

func newTask(instants []*of.MeshInstant) {
	task := of.NewTask(func() {
		z := float32(math.Sin(float64(time.Now().Second())))
		for _, instant := range instants {
			instant.Transform.MoveRelative(mgl32.Vec3{0, 0, z})
		}
	})
	task.SetRepeating(true)
	dependices := make([]of.Data, len(instants))
	for i, instant := range instants {
		dependices[i] = instant
	}
	task.SetRaceTask(of.GetEngineTask(of.RenderTask))
	of.AddTask(task)
}

func stop() {

}

var globalInt int

func TestGoRotine(t *testing.T) {

	for i := 0; i < 100; i++ {
		go multiFunc(i)
	}
	multiFunc(101)

}
func multiFunc(id int) {
	for {
		globalInt++
		log.Printf("%d %d", id, globalInt)
	}
}

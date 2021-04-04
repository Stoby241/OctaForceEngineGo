// +build debug

package OctaForce

import (
	"time"
)

func initRender() {
	initGLFW()

	initOpenGL()
	initShaders()

	initMaterials()

	deleteShaders()
}

var (
	maxFPS           float64
	fps              float64
	renderFrameStart time.Time
	renderDeltaTime  float64
)

func runRender() {

	clearColor := [3]float32{0.0, 0.0, 0.0}

	wait := time.Duration(1.0 / maxFPS * 1000000000)
	for running {
		renderFrameStart = time.Now()

		processEvents()

		if window.ShouldClose() {
			running = false
		}

		preRender(clearColor)
		render3D()
		postRender()

		diff := time.Since(renderFrameStart)
		if diff > 0 {
			fps = (wait.Seconds() / diff.Seconds()) * maxFPS
		} else {
			fps = 10000
		}

		if diff < wait {
			renderDeltaTime = wait.Seconds()
			time.Sleep(wait - diff)
		} else {
			renderDeltaTime = diff.Seconds()
		}
	}
}

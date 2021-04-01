package OctaForce

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
	"log"
	"path/filepath"
	"runtime"
	"time"
)

var absPath string

func init() {
	_, b, _, _ := runtime.Caller(0)
	absPath = filepath.Dir(b)
}

func initRender() {
	initImGui()

	log.Print("test")

	initGLFW()
	log.Print("test")
	initImGuiKeyMapping()

	initOpenGL()
	initImGuiGLBuffers()
}

var (
	maxFPS           float64
	fps              float64
	renderFrameStart time.Time
	renderDeltaTime  float64
	imGuiFunc        func()
)

func SetImGuiFunc(function func()) {
	imGuiFunc = function
}

func runRender() {

	clearColor := [3]float32{0.0, 0.0, 0.0}

	wait := time.Duration(1.0 / maxFPS * 1000000000)
	for running {
		renderFrameStart = time.Now()

		processEvents()

		if window.ShouldClose() {
			running = false
		}

		newFrame()
		imgui.NewFrame()

		imgui.Text(fmt.Sprintf("FPS : %.0f", fps))
		imgui.Text(fmt.Sprintf("UPS : %.0f", ups))

		if imGuiFunc != nil {
			imGuiFunc()
		}

		imgui.Render()

		preRender(clearColor)

		render3D()
		renderImGui(DisplaySize(), FramebufferSize(), imgui.RenderedDrawData())
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

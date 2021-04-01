package OctaForce

import (
	"log"
	"runtime"
)

var stopFunc func()

func SetStopFunc(function func()) {
	stopFunc = function
}

func Init(gameStartFunc func()) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.LockOSThread()

	maxUPS = 60
	maxFPS = 30
	running = true

	initState()
	initActiveMeshesData()
	initActiveCamera()
	initWorker()
	initDispatcher()

	initRender()

	gameStartFunc()

	go runDispatcher()
	runRender()

	if stopFunc != nil {
		stopFunc()
	}
}

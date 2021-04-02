package OctaForce

import (
	"log"
	"path/filepath"
	"runtime"
)

var absPath string

func init() {
	_, b, _, _ := runtime.Caller(0)
	absPath = filepath.Dir(b)
}

var stopFunc func()

func SetStopFunc(function func()) {
	stopFunc = function
}

var debug bool

func Init(gameStartFunc func()) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.LockOSThread()

	maxUPS = 60
	maxFPS = 30
	running = true

	initState()
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

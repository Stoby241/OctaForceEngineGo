package OctaForce

var workerTasks chan *task

const workerAmmount = 12

func initWorker() {
	workerTasks = make(chan *task, workerAmmount)

	for i := 0; i < workerAmmount; i++ {
		go runWorker()
	}
}

func runWorker() {
	for task := range workerTasks {
		task.function()
		task.done <- true
	}
}

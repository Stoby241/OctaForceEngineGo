package OctaForce

import (
	"time"
)

var (
	running          bool
	ups              float64
	maxUPS           float64
	updateFrameStart time.Time
	updateDeltaTime  float64
)

func GetDeltaTime() float64 {
	return updateDeltaTime
}

func initDispatcher() {
	addTask = make(chan *task, 1)
	removeTask = make(chan *task, 1)

	updateTasksListSync = make(chan bool)
	updateTasksListRelease = make(chan bool)

	go updateTasksList()
}

func runDispatcher() {
	wait := time.Duration(1.0 / maxUPS * 1000000000)

	for running {
		updateFrameStart = time.Now()

		updateTasksListSync <- true

		copyTaskSlices()

		updateTasksListRelease <- true

		dispatchTasks()

		diff := time.Since(updateFrameStart)
		if diff > 0 {
			ups = (wait.Seconds() / diff.Seconds()) * maxUPS
		} else {
			ups = 10000
		}

		if diff < wait {
			updateDeltaTime = wait.Seconds()
			time.Sleep(wait - diff)
		} else {
			updateDeltaTime = diff.Seconds()
		}
	}
}

var addTask chan *task

func AddTask(task *task) {
	addTask <- task
}

var removeTask chan *task

func RemoveTask(task *task) {
	removeTask <- task
}

var repeatingTasks []*task
var oneTimeTasks []*task
var updateTasksListSync chan bool
var updateTasksListRelease chan bool

func updateTasksList() {
	for running {
		select {
		case <-updateTasksListSync:
			<-updateTasksListRelease

		case task := <-addTask:
			if task.repeating {
				repeatingTasks = append(repeatingTasks, task)
			} else {
				oneTimeTasks = append(oneTimeTasks, task)
			}

		case task := <-removeTask:
			if !task.repeating {
				break
			}

			for i, repeatingTask := range repeatingTasks {
				if repeatingTask == task {
					repeatingTasks = append(repeatingTasks[:i], repeatingTasks[i+1:]...)
					break
				}
			}
		}
	}
}

var tasks []*task

func copyTaskSlices() {
	tasks = nil

	for _, task := range repeatingTasks {
		tasks = append(tasks, task)
	}

	for _, task := range oneTimeTasks {
		tasks = append(tasks, task)
	}
	oneTimeTasks = nil
}

var done bool
var aktiveTasks []*task

func dispatchTasks() {
	done = false
	for !done {
		done = true

		for i := len(aktiveTasks) - 1; i >= 0; i-- {
			done = false
			select {
			case <-aktiveTasks[i].done:
				aktiveTasks = append(aktiveTasks[:i], aktiveTasks[i+1:]...)

			default:
			}
		}

		for i := len(tasks) - 1; i >= 0; i-- {
			done = false
			task := tasks[i]
			if canStartTask(task) {
				workerTasks <- task
				aktiveTasks = append(aktiveTasks, task)
			}
		}
	}
}

func canStartTask(task *task) bool {
	for _, aktiveTask := range aktiveTasks {
		for _, raceTask := range task.raceTasks {
			if aktiveTask == raceTask {
				return false
			}
		}

		for _, raceTask := range aktiveTask.raceTasks {
			if task == raceTask {
				return false
			}
		}
	}
	return true
}

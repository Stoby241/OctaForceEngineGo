package OctaForce

type Data interface{}

type taskTyp int

const (
	taskMax = 0
)

var engineTasks []*task

func initState() {
	engineTasks = make([]*task, taskMax)
	for i := range engineTasks {
		engineTasks[i] = NewTask(func() {})
	}
}

func GetEngineTask(id taskTyp) *task {
	return engineTasks[id]
}

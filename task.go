package OctaForce

type task struct {
	function  func()
	repeating bool
	done      chan bool
	raceTasks []*task
}

func NewTask(function func()) *task {
	return &task{
		function:  function,
		repeating: false,
		done:      make(chan bool, 1),
	}
}

func (t *task) SetRepeating(repeating bool) {
	t.repeating = repeating
}
func (t *task) SetRaceTask(tasks ...*task) {
	t.raceTasks = tasks
}

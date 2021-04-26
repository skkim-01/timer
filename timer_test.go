package Timer

import (
	"fmt"
	"testing"
	"time"

	Task "github.com/skkim-01/task"
)

// Create struct for Task
type TaskforTest struct{}

// Create Task Function
func (t *TaskforTest) Sum_value(x, y int) int {
	return x + y
}

// Create Task Callback Function
var LocalCallback Task.TaskCallback = func(cbResult []interface{}) {
	// Print callback result
	fmt.Println("### LOG ### task.callback.Result:")
	fmt.Println(cbResult...)
}

func Test_timer_done(t *testing.T) {
	fmt.Println("### LOG ### ------ Test_timer_done ------")
	// Create New Task
	t1 := Task.NewTask(&TaskforTest{}, "Sum_value", LocalCallback, 10, 20)
	// Close Task
	defer t1.Close()

	// Create New Timer
	timer := NewTimer(t1, 1000, 3)

	timer.Start()

	for {
		time.Sleep(time.Second)
		if timer.State() == STATE_DONE {
			break
		}
	}

	timer.Close()
}

func Test_timer_interrupt(t *testing.T) {
	fmt.Println("### LOG ### ------ Test_timer_interrupt ------")
	// Create New Task
	t1 := Task.NewTask(&TaskforTest{}, "Sum_value", LocalCallback, 10, 20)
	// Close Task
	defer t1.Close()

	// Create New Timer
	timer := NewTimer(t1, 1000, 0)

	timer.Start()

	// sleep 7 seconds
	for i := 0; i < 7; i++ {
		time.Sleep(time.Second)
	}

	fmt.Println("### LOG ### ------ stop timer ------")
	timer.Stop()

	fmt.Println("### LOG ### ------ close timer ------")
	timer.Close()
}

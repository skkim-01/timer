package Timer

import Task "github.com/skkim-01/task"

// NewTimer : Create new timer
func NewTimer(task *Task.Task, intv int64, cnt int) *Timer {
	t := Timer{
		timerTask: task,
		interval:  intv,
		count:     cnt,
		eventChan: make(chan int),
		state:     STATE_STOP,
	}
	return &t
}

// Start: Start timer object
func (t *Timer) Start() {
	t.state = STATE_IDLE
	go t.timerThread()
}

// Stop: Stop timer object
func (t *Timer) Stop() {
	if STATE_IDLE == t.state || STATE_RUN_TASK == t.state {
		t.state = STATE_STOP_WAITING
		t.eventChan <- 1
		<-t.eventChan
	}
	t.state = STATE_DONE
}

// State: return timer state - check TimerObject.go
func (t *Timer) State() int {
	return t.state
}

// Close: close timer object
func (t *Timer) Close() {
	t.timerTask.Close()
	close(t.eventChan)
}

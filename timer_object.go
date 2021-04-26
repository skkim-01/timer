package Timer

import (
	Task "github.com/skkim-01/task"
)

// enum : Timer state
const (
	STATE_STOP = 0 + iota
	STATE_IDLE
	STATE_RUN_TASK
	STATE_STOP_WAITING
	STATE_DONE
)

// Timer : struct for timer object
type Timer struct {
	timerTask *Task.Task // timer task
	interval  int64      // timer running interval - milli-second
	count     int        // timer running count. [0: infinity]
	eventChan chan int   // timer handling event channel
	state     int        // timer state.
}

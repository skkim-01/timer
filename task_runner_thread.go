package Timer

import (
	"time"
)

// timerThread
func (t *Timer) timerThread() {
	var invokeCount int = 0

	// _START_SECTION : timer thread loop start point
_START_SECTION:
	// invoke count check
	if 0 != t.count && invokeCount == t.count {
		goto _EXIT
	}

	if true == t.isFirstRunning && 0 == invokeCount {
		t._invokeTimer(&invokeCount)
	}

	// channel wait
	select {
	// case timeout
	case <-time.After(time.Millisecond * time.Duration(t.interval)):
		t._invokeTimer(&invokeCount)
		goto _START_SECTION

	// case exit event
	case <-t.eventChan:
		t.state = STATE_STOP_WAITING
		goto _EXIT_EVENT
	}

	// _EXIT_EVENT : when catch exit event, send channel response
_EXIT_EVENT:
	t.eventChan <- 1

	// _EXIT : when count is over, exit
_EXIT:
	t.state = STATE_DONE
}

// _invokeTimer : timer task invoke, update timer status and increase invoke count
func (t *Timer) _invokeTimer(invokeCount *int) {
	t.state = STATE_RUN_TASK
	t.timerTask.InvokeTask()
	*invokeCount++
	t.state = STATE_IDLE
}

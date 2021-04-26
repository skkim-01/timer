# timer

### Package Name : Timer

#### abstract
##### timer. using task.  [ https://github.com/skkim-01/task ]

#### type definitions
###### Timer struct
```go
  type Timer struct {
    timerTask *Task.Task // timer task
    interval  int64      // timer running interval: milli-second
    count     int        // timer running count. [0: infinity]
    eventChan chan int   // timer handling event channel
    state     int        // timer state.
  }
```
###### Timer state value
```go
  const (
    STATE_STOP = 0 + iota
    STATE_IDLE
    STATE_RUN_TASK
    STATE_STOP_WAITING
    STATE_DONE
  )
```

#### APIs
###### NewTimer(): Create new timer.
```go
  // intv: interval. wait milliseconds.
  // cnt: repeat count. 0: infinity
  func NewTimer(task *Task.Task, intv int64, cnt int) *Timer   
```

##### Timer
###### Start(): Start timer.
```go
  func (t *Timer) Start()
```

###### Stop(): Stop timer. (doesn't close thread msg channel)
```go
  func (t *Timer) Stop()
```

###### State(): Timer state.
```go
  func (t *Timer) State()
```

###### Close(): Close timer. (close thread msg channel)
```go
  func (t *Timer) Close()
```

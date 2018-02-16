package stopwatch

import (
	"bytes"
	"fmt"
	"math"
	"time"
)

type StopWatch struct {
	id          string
	running     bool
	runningTask string
	keepTasks   bool
	tasks       []*TaskInfo
	taskCount   int
	lastTask    *TaskInfo
	start       time.Duration
	total       time.Duration
}

func New() *StopWatch {
	return NewID("")
}

func NewID(id string) *StopWatch {
	sw := new(StopWatch)
	sw.id = id
	sw.keepTasks = true
	return sw
}

func (sw *StopWatch) ID() string {
	return sw.id
}

func (sw *StopWatch) SetKeepTasks(keepTasks bool) {
	sw.keepTasks = keepTasks
}

func (sw *StopWatch) Start() bool {
	return sw.StartTask("")
}

func (sw *StopWatch) StartTask(name string) bool {
	if sw.running {
		return false
	}
	sw.running = true
	sw.runningTask = name
	sw.start = Start()
	return true
}

func (sw *StopWatch) Stop() bool {
	if !sw.running {
		return false
	}
	last := Stop(sw.start)
	sw.total += last
	sw.lastTask = &TaskInfo{Name: sw.runningTask, Elapsed: last}
	if sw.keepTasks {
		sw.tasks = append(sw.tasks, sw.lastTask)
	}
	sw.taskCount++
	sw.running = false
	sw.runningTask = ""
	return true
}

func (sw *StopWatch) Running() bool {
	return sw.running
}

func (sw *StopWatch) RunningTaskName() string {
	return sw.runningTask
}

func (sw *StopWatch) LastTask() *TaskInfo {
	return sw.lastTask
}

func (sw *StopWatch) AllTasks() []*TaskInfo {
	return sw.tasks
}

func (sw *StopWatch) TaskCount() int {
	return sw.taskCount
}

func (sw *StopWatch) Elapsed() time.Duration {
	return sw.total
}

func (sw *StopWatch) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "StopWatch '%s': running time (millis) = %d", sw.id, sw.total/time.Millisecond)
	if sw.keepTasks {
		for _, t := range sw.tasks {
			percent := math.Floor((100*t.Elapsed.Seconds())+0.5) / sw.total.Seconds()
			fmt.Fprintf(w, "; [%s] took %d = %.3f%%", t.Name, t.Elapsed/time.Millisecond, percent)
		}
	}
	return w.String()
}

type TaskInfo struct {
	Name    string
	Elapsed time.Duration
}

package fsm

import (
	"context"
	"sync"
	"time"
)

// Queue 异步事件队列
type Queue struct {
	sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
	tasks  []func()
	dt     time.Duration
	closed bool
}

func NewQueue(dt time.Duration) *Queue {
	ctx, cancel := context.WithCancel(context.Background())

	ret := &Queue{
		ctx:    ctx,
		cancel: cancel,
		dt:     dt,
	}
	return ret
}

func (q *Queue) Start() {
	go func() {
		for {
			select {
			case <-q.ctx.Done():
				return
			default:
				q.parking()
			}
		}
	}()
}

// 等待dt后执行队列中任务
func (q *Queue) parking() {
	// 让出cpu
	stopTimer := time.After(q.dt)

	// 等待定时器到期或者其他条件
	select {
	case <-stopTimer:
		// 定时器到期，执行这里的代码
		q.executeAll()
		//这里放置需要在定时器到期后执行的操作
	}

}

func (q *Queue) executeAll() {
	q.Lock()
	var tasksToExecute []func()
	if len(q.tasks) > 0 {
		tasksToExecute = append(tasksToExecute, q.tasks...)
		q.tasks = nil
	}
	q.Unlock()

	for _, t := range tasksToExecute {
		t()
	}
}
func (q *Queue) Shutdown() {
	q.closeTasks()
	// 确保在退出前执行所有任务
	q.executeAll()
	q.cancel()
}

// stop task列表不再新增
func (q *Queue) closeTasks() {
	q.Lock()
	defer q.Unlock()
	q.closed = true
}
func (q *Queue) AddTask(f func()) {
	q.Lock()
	defer q.Unlock()
	if q.closed {
		return
	}
	q.tasks = append(q.tasks, f)
}

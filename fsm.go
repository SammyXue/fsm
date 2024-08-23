package fsm

import (
	"sync"
	"time"
)

const timeDuration = 10 * time.Millisecond

// StateMachine 有限状态机模型
type StateMachine struct {
	sync.Once
	sync.WaitGroup
	Now State
	q   *Queue
}

func NewStateMachine() *StateMachine {

	ret := &StateMachine{
		q: NewQueue(timeDuration),
	}
	ret.Add(1)
	return ret
}

func (sm *StateMachine) Start(state State) {
	sm.Now = state
	sm.q.Start()
}

func (sm *StateMachine) ProcessInput(input Input) {

	sm.q.AddTask(func() {
		if sm.Now == nil {
			sm.end()
		}

		sm.Now = sm.Now.HandleInput(input)
		if sm.Now == nil || sm.Now.IsEnd() {
			sm.end()
		}

	})
}

func (sm *StateMachine) end() {
	sm.Do(func() {
		sm.q.Shutdown()
		sm.Done()

	})

}

func (sm *StateMachine) WaitForEnd() {
	sm.Wait()
}

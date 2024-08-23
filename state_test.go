package fsm

import (
	"testing"
)

type MockStateStart struct {
	output string
}

func (m *MockStateStart) HandleInput(input Input) State {
	// 这里可以根据需要模拟不同的行为
	return &MockStateEnd{output: "end"}
}

func (m *MockStateStart) IsEnd() bool {
	return false
}

type MockStateEnd struct {
	output string
}

func (m *MockStateEnd) HandleInput(input Input) State {
	return &MockStateEnd{}
}

func (m *MockStateEnd) IsEnd() bool {
	return true
}

func TestStateMachine(t *testing.T) {
	fsm := NewStateMachine()
	fsm.Start(&MockStateStart{output: "start"})
	t.Logf("%+v", fsm.Now)
	input := Input{} // 假设 Input 是一个有效的输入
	fsm.ProcessInput(input)
	fsm.WaitForEnd()
	t.Logf("%+v", fsm.Now)

}

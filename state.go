package fsm

// State 定义状态接口，包含一个处理输入的方法 是否状态终结
type State interface {
	HandleInput(input Input) State
	IsEnd() bool
}

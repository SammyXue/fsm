package fsm

type Input struct {
	Action int
	Param  string
}

func NewInput(Action int) Input {
	return Input{Action: Action}
}

func (i Input) WithParam(s string) Input {
	i.Param = s
	return i
}

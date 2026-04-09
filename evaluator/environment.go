package evaluator

type Environment struct {
	store map[string]int64
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]int64),
	}
}

func (e *Environment) Get(name string) (int64, bool) {
	val, ok := e.store[name]
	return val, ok
}

func (e *Environment) Set(name string, val int64) {
	e.store[name] = val
}

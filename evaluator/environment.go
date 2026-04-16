package evaluator

type Environment struct {
	store map[string]interface{}
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]interface{}),
	}
}

func (e *Environment) Get(name string) (interface{}, bool) {
	val, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	return val, ok
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Set(name string, val interface{}) interface{} {
	e.store[name] = val
	return val
}

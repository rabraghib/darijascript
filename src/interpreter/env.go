package interpreter

import "github.com/rabraghib/darijascript/src/parser"

type Environment struct {
	store     map[string]interface{}
	outer     *Environment
	functions map[string]*parser.FunctionDeclaration
}

func NewEnvironment() *Environment {
	return &Environment{
		outer:     nil,
		store:     make(map[string]interface{}),
		functions: make(map[string]*parser.FunctionDeclaration),
	}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (env *Environment) Get(name string) (interface{}, bool) {
	val, ok := env.store[name]
	if !ok && env.outer != nil {
		return env.outer.Get(name)
	}
	return val, ok
}

func (env *Environment) Set(name string, value interface{}) {
	env.store[name] = value
}

func (env *Environment) Update(name string, value interface{}) {
	if _, ok := env.store[name]; ok {
		env.store[name] = value
	} else if env.outer != nil {
		env.outer.Update(name, value)
	}
}

func (env *Environment) GetFunction(name string) (*parser.FunctionDeclaration, bool) {
	fn, ok := env.functions[name]
	if !ok && env.outer != nil {
		return env.outer.GetFunction(name)
	}
	return fn, ok
}

func (env *Environment) SetFunction(name string, fn *parser.FunctionDeclaration) {
	env.functions[name] = fn
}

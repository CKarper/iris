package js

import (
	"errors"
	"fmt"

	"github.com/robertkrimen/otto"

	// Add underscore.js to otto
	"github.com/dop251/goja"
	_ "github.com/robertkrimen/otto/underscore"
)

const (
	filterNotDefinedError = "Filter not defined"
)

type (
	loader interface {
	}
	functionCaller interface {
	}

	Runner interface {
		Load(scripts []string, gloabl map[string]interface{}) (*otto.Value, error)
		CallFn(name string, input ...interface{}) (*otto.Value, error)
	}

	runner struct {
		vm  *otto.Otto
		nvm *goja.Runtime
	}
)

func NewJSEngine() Runner {
	return &runner{
		vm:  otto.New(),
		nvm: goja.New(),
	}
}

func (r *runner) Load(scripts []string, input map[string]interface{}) (*otto.Value, error) {
	global := ""
	for _, script := range scripts {
		c, err := r.vm.Compile("", string(script))
		if err != nil {
			return nil, err
		}
		global = fmt.Sprintf("%s\n%s", global, c.String())
	}
	for k, v := range input {
		err := r.vm.Set(k, v)
		if err != nil {
			return nil, err
		}
	}
	res, err := r.vm.Run(global)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *runner) CallFn(name string, input ...interface{}) (*otto.Value, error) {
	v, err := r.vm.Get(name)
	if err != nil {
		return nil, err
	}
	if v.IsDefined() {
		res, err := v.Call(otto.Value{}, input...)
		return &res, err
	}
	return nil, errors.New(filterNotDefinedError)
}

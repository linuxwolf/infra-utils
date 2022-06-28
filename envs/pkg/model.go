package pkg

import (
	"fmt"
	"sort"
	"strings"
)

type Env struct {
	vars map[string]string
}

func NewEnvWith(values map[string]string) *Env {
	result := &Env{vars: map[string]string{}}
	for k, v := range values {
		result.vars[k] = v
	}

	return result
}

func (e *Env) Variables() map[string]string {
	result := map[string]string{}
	for k, v := range e.vars {
		result[k] = v
	}
	return result
}

func (e *Env) Including(other *Env) *Env {
	result := NewEnvWith(other.vars)
	for k, v := range e.vars {
		result.vars[k] = v
	}

	return result
}

func (e *Env) Excluding(other *Env) *Env {
	result := NewEnvWith(e.vars)
	for k, _ := range other.vars {
		delete(result.vars, k)
	}

	return result
}

func (e *Env) String() string {
	var idx int

	keys := make([]string, len(e.vars))
	idx = 0
	for k, _ := range e.vars {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)

	output := make([]string, len(e.vars))
	idx = 0
	for _, k := range keys {
		v := e.vars[k]
		// TODO: escape newlines
		output[idx] = fmt.Sprintf("%s=%s", k, v)
		idx++
	}

	return strings.Join(output, "\n")
}

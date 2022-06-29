package pkg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestModel(t *testing.T) {
	suite.Run(t, new(ModelTestSuite))
}

type ModelTestSuite struct{ suite.Suite }

func (suite *ModelTestSuite) TestNewEnvWith() {
	T := suite.T()
	var actual *Env

	actual = NewEnvWith(nil)
	assert.Equal(T, len(actual.vars), 0)
	assert.Equal(T, actual.Variables(), map[string]string{})

	actual = NewEnvWith(map[string]string{
		"FOO": "foo value",
		"BAR": "bar value",
	})
	assert.Equal(T, actual.vars, map[string]string{
		"FOO": "foo value",
		"BAR": "bar value",
	})
	assert.Equal(T, actual.Variables(), map[string]string{
		"FOO": "foo value",
		"BAR": "bar value",
	})
	assert.Equal(T, actual.String(), strings.TrimSpace(`
BAR="bar value"
FOO="foo value"`))
}

func (suite *ModelTestSuite) TestIncluding() {
	T := suite.T()

	env1 := NewEnvWith(map[string]string{
		"FOO": "env1 foo",
		"BAR": "env1 bar",
	})
	env2 := NewEnvWith(map[string]string{
		"FOO": "env2 foo",
		"BAZ": "env2 baz",
	})

	var result *Env
	result = env1.Including(env2)
	assert.Equal(T, result.vars, map[string]string{
		"FOO": "env1 foo",
		"BAR": "env1 bar",
		"BAZ": "env2 baz",
	})
	assert.Equal(
		T,
		result.String(),
		strings.TrimSpace(`
BAR="env1 bar"
BAZ="env2 baz"
FOO="env1 foo"`),
	)

	result = env2.Including(env1)
	assert.Equal(T, result.vars, map[string]string{
		"FOO": "env2 foo",
		"BAR": "env1 bar",
		"BAZ": "env2 baz",
	})
	assert.Equal(
		T,
		result.String(),
		strings.TrimSpace(`
BAR="env1 bar"
BAZ="env2 baz"
FOO="env2 foo"`),
	)

	result = env1.Including(NewEnvWith(nil))
	assert.Equal(T, result.vars, map[string]string{
		"FOO": "env1 foo",
		"BAR": "env1 bar",
	})
	assert.Equal(
		T,
		result.String(),
		strings.TrimSpace(`
BAR="env1 bar"
FOO="env1 foo"`),
	)
}

func (suite *ModelTestSuite) TestExcluding() {
	T := suite.T()

	env1 := NewEnvWith(map[string]string{
		"FOO": "env1 foo",
		"BAR": "env1 bar",
	})
	env2 := NewEnvWith(map[string]string{
		"FOO": "env2 foo",
		"BAZ": "env2 baz",
	})

	var result *Env
	result = env1.Excluding(env2)
	assert.Equal(T, result.vars, map[string]string{
		"BAR": "env1 bar",
	})
	result = env2.Excluding(env1)
	assert.Equal(T, result.vars, map[string]string{
		"BAZ": "env2 baz",
	})
	result = env1.Excluding(NewEnvWith(nil))
	assert.Equal(T, result.vars, map[string]string{
		"FOO": "env1 foo",
		"BAR": "env1 bar",
	})
}

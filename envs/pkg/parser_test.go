package pkg

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestParser(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}

type ParserTestSuite struct {
	suite.Suite
	parser  *Parser
	environ []string
}

func (suite *ParserTestSuite) SetupTest() {
	suite.parser = NewParser(nil)
	suite.environ = os.Environ()

	os.Clearenv()
	os.Setenv("ENV_FOO", "environ foo")
	os.Setenv("ENV_BAR", "environ bar")
	os.Setenv("ENV_BAZ", "environ baz")
}
func (suite *ParserTestSuite) TearDownTest() {
	os.Clearenv()
	for _, e := range suite.environ {
		parts := strings.SplitN(e, "=", 2)
		os.Setenv(parts[0], parts[1])
	}
}

func (suite *ParserTestSuite) TestParseLine_Simple() {
	T := suite.T()
	p := suite.parser

	k, v, err := p.parseLine("FOO=foo value")
	assert.Equal(T, k, "FOO")
	assert.Equal(T, v, "foo value")
	assert.Equal(T, err, nil)
}

func (suite *ParserTestSuite) TestParseLine_Errors() {
	T := suite.T()
	var (
		k, v string
		err  error
	)
	p := suite.parser

	k, v, err = p.parseLine("")
	assert.Equal(T, k, "")
	assert.Equal(T, v, "")
	assert.Equal(T, err, ErrParseEmptyLine)

	k, v, err = p.parseLine("# this is a comment")
	assert.Equal(T, k, "")
	assert.Equal(T, v, "")
	assert.Equal(T, err, ErrParseCommentLine)

	k, v, err = p.parseLine("this value=nope")
	assert.Equal(T, k, "")
	assert.Equal(T, v, "")
	assert.Equal(T, err, ErrParseInvalid)
}

func (suite *ParserTestSuite) TestProcessReader() {
	T := suite.T()
	source := `
# main env file
FOO=foo value
BAR=bar value

# bogus tests
not value=bogus
BOGUS=not not valid
`
	reader := strings.NewReader(source)
	p := suite.parser
	result := p.ProcessReader(reader)
	assert.Equal(T, result, NewEnvWith(map[string]string{
		"FOO":   "foo value",
		"BAR":   "bar value",
		"BOGUS": "not not valid",
	}))
}

func (suite *ParserTestSuite) TestProcessArray() {
	T := suite.T()
	source := []string{
		"# main env file",
		"FOO=foo value",
		"BAR=bar value",
		"",
		"# bogus tests",
		"not value=bogus",
		"BOGUS=not not valid",
	}

	p := suite.parser
	result := p.ProcessArray(source)
	assert.Equal(T, result, NewEnvWith(map[string]string{
		"FOO":   "foo value",
		"BAR":   "bar value",
		"BOGUS": "not not valid",
	}))
}

func (suite *ParserTestSuite) TestNewEnvsFromEnviron() {
	T := suite.T()
	envs := NewEnvsFromEnviron()
	assert.Equal(T, envs, NewEnvWith(map[string]string{
		"ENV_FOO": "environ foo",
		"ENV_BAR": "environ bar",
		"ENV_BAZ": "environ baz",
	}))
	assert.Equal(
		T,
		envs.String(),
		"ENV_BAR=environ bar\nENV_BAZ=environ baz\nENV_FOO=environ foo",
	)
}

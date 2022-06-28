package pkg

import (
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
	parser *Parser
}

func (suite *ParserTestSuite) SetupTest() {
	suite.parser = NewParser(nil)
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
		"FOO=foo value",
		"BAR=bar value",
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

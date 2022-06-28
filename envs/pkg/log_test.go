package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestLog(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}

type LogTestSuite struct{ suite.Suite }

func (suite *LogTestSuite) TestSetupLogging() {
	T := suite.T()

	result := SetupLogging()
	assert.NotNil(T, result)
}

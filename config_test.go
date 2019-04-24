package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) TestNewConfig() {
	config := NewConfig()
	assert.Equal(s.T(), config.Format, DefaultFormat)
	assert.Equal(s.T(), config.Level, DefaultLevel)
	assert.Nil(s.T(), config.OutputTo)
	assert.False(s.T(), config.Debug)
}

func (s *ConfigTestSuite) TestNewConfigFromEnvironment() {
	os.Setenv("TEST_"+EnvKeyFormat, "FORMAT_JSON")
	os.Setenv("TEST_"+EnvKeyLevel, "LEVEL_INFO")
	os.Setenv("TEST_"+EnvKeyDebug, "1")
	config := NewConfigFromEnvironment("TEST")
	assert.Equal(s.T(), config.Format, FormatJSON)
	assert.Equal(s.T(), config.Level, LevelInfo)
	assert.Nil(s.T(), config.OutputTo)
	assert.True(s.T(), config.Debug)
}

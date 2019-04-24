package logger

import (
	"bytes"
	"regexp"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoggerTestSuite struct {
	suite.Suite
	logs   bytes.Buffer
	config *Config
}

func TestLogger(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}

func (s *LoggerTestSuite) SetupTest() {
	s.logs = bytes.Buffer{}
	s.config = NewConfig()
	s.config.OutputTo = &s.logs
}

func (s *LoggerTestSuite) TestDebug_displaysCurrentFilename() {
	s.config.Debug = true
	logger := New(s.config)
	logger.Trace("__TestDebug_displaysCurrentFilename")
	_, file, _, ok := runtime.Caller(0)
	assert.True(s.T(), ok)
	assert.Contains(s.T(), s.logs.String(), file)
}

func (s *LoggerTestSuite) TestDebug_displaysCurrentFunction() {
	s.config.Debug = true
	logger := New(s.config)
	logger.Trace("__TestDebug_displaysCurrentFunction")
	programCounter, _, _, ok := runtime.Caller(0)
	assert.True(s.T(), ok)
	assert.Contains(s.T(), s.logs.String(), runtime.FuncForPC(programCounter).Name())
}

func (s *LoggerTestSuite) TestFields_withDebug() {
	s.config.Debug = true
	s.config.Format = FormatJSON
	logger := New(s.config)
	logger.WithFields(map[string]interface{}{
		"random": "field",
	}).Trace("__TestFields_withDebug")
	assert.Contains(s.T(), s.logs.String(), `"@file": "`)
	assert.Contains(s.T(), s.logs.String(), `"@caller": "`)
	assert.Contains(s.T(), s.logs.String(), `"@level": "`)
	assert.Contains(s.T(), s.logs.String(), `"@message": "`)
	assert.Contains(s.T(), s.logs.String(), `"@timestamp": "`)
	assert.Contains(s.T(), s.logs.String(), `"@data": {`)
}

func (s *LoggerTestSuite) TestFields_withJSON() {
	s.config.Debug = true
	s.config.Format = FormatJSON
	s.config.Fields["test"] = "fields"
	logger := New(s.config)
	logger.Trace("__TestFields_withJSON")
	assert.Contains(s.T(), s.logs.String(), "__TestFields_withJSON")
	assert.Regexp(s.T(), regexp.MustCompile(`"test": "fields"`), s.logs.String())
}

func (s *LoggerTestSuite) TestFields_withText() {
	s.config.Format = FormatText
	s.config.Fields["test"] = "fields"
	logger := New(s.config)
	logger.Trace("__TestFields_withText")
	assert.Contains(s.T(), s.logs.String(), "__TestFields_withText")
	assert.Regexp(s.T(), regexp.MustCompile(`test.*=fields`), s.logs.String())
}

func (s *LoggerTestSuite) TestLevels() {
	s.config.Level = LevelDebug
	logger := New(s.config)
	logger.Trace("__TestLevels")
	assert.Len(s.T(), s.logs.String(), 0)
	logger.Debug("__TestLevels")
	assert.Contains(s.T(), s.logs.String(), "__TestLevels")
}

func (s *LoggerTestSuite) TestOutput() {
	logger := New(s.config)
	logger.Trace("__TestOutput")
	assert.Contains(s.T(), s.logs.String(), "__TestOutput")
}

package server

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigKeyTestSuite struct {
	suite.Suite
}

func TestConfigKey(t *testing.T) {
	suite.Run(t, new(ConfigKeyTestSuite))
}

func (s *ConfigKeyTestSuite) TestGetBool() {
	helper := viper.New()
	environmentKey := &EnvironmentKey{
		DefaultValue: true,
		Name:         "GET_BOOL",
		helper:       helper,
	}
	assert.Equal(s.T(), true, environmentKey.GetBool())
}

func (s *ConfigKeyTestSuite) TestGetDuration() {
	helper := viper.New()
	environmentKey := &EnvironmentKey{
		DefaultValue: time.Second * 123,
		Name:         "GET_DURATION",
		helper:       helper,
	}
	assert.Equal(s.T(), time.Second*123, environmentKey.GetDuration())
}

func (s *ConfigKeyTestSuite) TestGetInt() {
	helper := viper.New()
	environmentKey := &EnvironmentKey{
		DefaultValue: 12345,
		Name:         "GET_INT",
		helper:       helper,
	}
	assert.Equal(s.T(), 12345, environmentKey.GetInt())
}

func (s *ConfigKeyTestSuite) TestGetString() {
	helper := viper.New()
	environmentKey := &EnvironmentKey{
		DefaultValue: "some string",
		Name:         "GET_STRING",
		helper:       helper,
	}
	assert.Equal(s.T(), "some string", environmentKey.GetString())
}

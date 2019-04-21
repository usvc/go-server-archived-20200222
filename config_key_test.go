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

func (s *ConfigKeyTestSuite) TestHasValidTLS_valid() {
	config := &Config{
		TLSCertificatePath: "secrets/tls.crt",
		TLSKeyPath:         "secrets/tls.key",
	}
	assert.Equal(s.T(), true, config.HasValidTLS())
}

func (s *ConfigKeyTestSuite) TestHasValidTLS_noExists() {
	config := &Config{
		TLSCertificatePath: "secrets/tls3.crt",
		TLSKeyPath:         "secrets/tls.key",
	}
	assert.Equal(s.T(), false, config.HasValidTLS())
	config = &Config{
		TLSCertificatePath: "secrets/tls.crt",
		TLSKeyPath:         "secrets/tls3.key",
	}
	assert.Equal(s.T(), false, config.HasValidTLS())
}

func (s *ConfigKeyTestSuite) TestHasValidTLS_mismatch() {
	config := &Config{
		TLSCertificatePath: "secrets/tls.crt",
		TLSKeyPath:         "secrets/tls2.key",
	}
	assert.Equal(s.T(), false, config.HasValidTLS())
	config = &Config{
		TLSCertificatePath: "secrets/tls2.crt",
		TLSKeyPath:         "secrets/tls.key",
	}
	assert.Equal(s.T(), false, config.HasValidTLS())
}

func (s *ConfigKeyTestSuite) TestHasValidTLS_undefined() {
	config := &Config{
		TLSCertificatePath: "",
		TLSKeyPath:         "",
	}
	assert.Equal(s.T(), false, config.HasValidTLS())
	config = &Config{
		TLSCertificatePath: "secrets/tls.crt",
		TLSKeyPath:         "",
	}
	assert.Equal(s.T(), false, config.HasValidTLS())
	config = &Config{
		TLSCertificatePath: "",
		TLSKeyPath:         "secrets/tls.key",
	}
	assert.Equal(s.T(), false, config.HasValidTLS())
}

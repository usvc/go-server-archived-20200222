package server

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) TestNewConfigFromEnvironment() {
	config := NewConfigFromEnvironment()
	assert.Equal(s.T(), DefaultHost, config.Host)
	assert.Equal(s.T(), DefaultLivenessCheckInterval, config.LivenessCheckInterval)
	assert.Equal(s.T(), DefaultLivenessCheckMethod, config.LivenessCheckMethod)
	assert.Equal(s.T(), DefaultLivenessCheckPath, config.LivenessCheckPath)
	assert.Equal(s.T(), DefaultLivenessCheckStatusCode, config.LivenessCheckStatusCode)
	assert.Equal(s.T(), DefaultLivenessCheckTimeout, config.LivenessCheckTimeout)
	assert.Equal(s.T(), DefaultMaxHeaderBytes, config.MaxHeaderBytes)
	assert.Equal(s.T(), DefaultPort, config.Port)
	assert.Equal(s.T(), DefaultTimeout, config.TimeoutIdle)
	assert.Equal(s.T(), DefaultTimeout, config.TimeoutRead)
	assert.Equal(s.T(), DefaultTLSCertificatePath, config.TLSCertificatePath)
	assert.Equal(s.T(), DefaultTLSKeyPath, config.TLSKeyPath)
}

func (s *ConfigTestSuite) TestNewConfigFromEnvironment_withPrefix() {
	prefix := "PREFIX_TEST"
	expectedConfig := &Config{
		Host:                    "1.1.1.1",
		LivenessCheckInterval:   time.Second * 10000,
		LivenessCheckMethod:     "METHOD",
		LivenessCheckPath:       "/liveness/check/path",
		LivenessCheckStatusCode: 418,
		LivenessCheckTimeout:    time.Second * 20000,
		MaxHeaderBytes:          4096,
		Port:                    "99999",
		TimeoutIdle:             time.Second * 30000,
		TimeoutRead:             time.Second * 30000,
		TLSCertificatePath:      "/path/to/cert",
		TLSKeyPath:              "/path/to/key",
	}
	os.Setenv(prefix+"_HOST", expectedConfig.Host)
	os.Setenv(prefix+"_LIVENESS_CHECK_INTERVAL", "10000s")
	os.Setenv(prefix+"_LIVENESS_CHECK_METHOD", expectedConfig.LivenessCheckMethod)
	os.Setenv(prefix+"_LIVENESS_CHECK_PATH", expectedConfig.LivenessCheckPath)
	os.Setenv(prefix+"_LIVENESS_CHECK_STATUS_CODE", strconv.Itoa(expectedConfig.LivenessCheckStatusCode))
	os.Setenv(prefix+"_LIVENESS_CHECK_TIMEOUT", "20000s")
	os.Setenv(prefix+"_MAX_HEADER_BYTES", strconv.Itoa(expectedConfig.MaxHeaderBytes))
	os.Setenv(prefix+"_PORT", expectedConfig.Port)
	os.Setenv(prefix+"_TIMEOUT_IDLE", "30000s")
	os.Setenv(prefix+"_TIMEOUT_READ", "30000s")
	os.Setenv(prefix+"_TLS_CERTIFICATE_PATH", expectedConfig.TLSCertificatePath)
	os.Setenv(prefix+"_TLS_KEY_PATH", expectedConfig.TLSKeyPath)
	config := NewConfigFromEnvironment(prefix)
	assert.Equal(s.T(), expectedConfig.Host, config.Host)
	assert.Equal(s.T(), expectedConfig.LivenessCheckInterval, config.LivenessCheckInterval)
	assert.Equal(s.T(), expectedConfig.LivenessCheckMethod, config.LivenessCheckMethod)
	assert.Equal(s.T(), expectedConfig.LivenessCheckPath, config.LivenessCheckPath)
	assert.Equal(s.T(), expectedConfig.LivenessCheckStatusCode, config.LivenessCheckStatusCode)
	assert.Equal(s.T(), expectedConfig.LivenessCheckTimeout, config.LivenessCheckTimeout)
	assert.Equal(s.T(), expectedConfig.MaxHeaderBytes, config.MaxHeaderBytes)
	assert.Equal(s.T(), expectedConfig.Port, config.Port)
	assert.Equal(s.T(), expectedConfig.TimeoutIdle, config.TimeoutIdle)
	assert.Equal(s.T(), expectedConfig.TimeoutRead, config.TimeoutRead)
	assert.Equal(s.T(), expectedConfig.TLSCertificatePath, config.TLSCertificatePath)
	assert.Equal(s.T(), expectedConfig.TLSKeyPath, config.TLSKeyPath)
}

func (s *ConfigTestSuite) TestGetAddr() {
	config := &Config{
		Host: "host",
		Port: "111111",
	}
	assert.Equal(s.T(), "host:111111", config.GetAddr())
}

func (s *ConfigTestSuite) TestGetProtocol() {
	config := &Config{
		TLSCertificatePath: "secrets/tls.crt",
		TLSKeyPath:         "secrets/tls.key",
	}
	assert.Equal(s.T(), "https", config.GetProtocol())
}

func (s *ConfigTestSuite) TestGetProtocol_invalidCertKey() {
	config := &Config{
		TLSCertificatePath: "secrets/tls3.crt",
		TLSKeyPath:         "secrets/tls3.key",
	}
	assert.Equal(s.T(), "http", config.GetProtocol())
}

func (s *ConfigTestSuite) TestGetURL() {
	config := &Config{
		Host:               "host",
		Port:               "111111",
		TLSCertificatePath: "secrets/tls.crt",
		TLSKeyPath:         "secrets/tls.key",
	}
	assert.Equal(s.T(), "https://host:111111", config.GetURL())
}

func (s *ConfigTestSuite) TestHasValidTLS() {
	config := &Config{
		TLSCertificatePath: "secrets/tls.crt",
		TLSKeyPath:         "secrets/tls.key",
	}
	assert.Equal(s.T(), true, config.HasValidTLS())
}

func (s *ConfigTestSuite) TestHasValidTLS_invalidTLS() {
	config := &Config{
		TLSCertificatePath: "secrets/tls.crt",
		TLSKeyPath:         "secrets/tls2.key",
	}
	assert.Equal(s.T(), false, config.HasValidTLS())
}

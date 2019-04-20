package server

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// DefaultHost defines the default host address to bind to
const DefaultHost = "0.0.0.0"

// DefaultLivenessCheckInterval defines the default interval between responsiveness checks
const DefaultLivenessCheckInterval = time.Millisecond * 250

// DefaultLivenessCheckMethod defines the default method used for the responsiveness check
const DefaultLivenessCheckMethod = "get"

// DefaultLivenessCheckPath defines the default uri path used for the responsiveness check
const DefaultLivenessCheckPath = "/"

// DefaultLivenessCheckStatusCode defines the default expected status code for responsiveness check
const DefaultLivenessCheckStatusCode = 200

// DefaultLivenessCheckTimeout defines the default number of seconds the responsiveness check will be performed
const DefaultLivenessCheckTimeout = time.Second * 15

// DefaultMaxHeaderBytes defines the default number of bytes allowed for the header
const DefaultMaxHeaderBytes = 1024

// DefaultPort defines the default port to bind to
const DefaultPort = "2222"

// DefaultTimeout defines the default duration before requests time out
const DefaultTimeout = time.Second * 10

// DefaultTLSCertificatePath defines the default relative path to the server certificate
const DefaultTLSCertificatePath = "tls/server.crt"

// DefaultTLSKeyPath defines the default relative path to the certificate key
const DefaultTLSKeyPath = "tls/server.key"

func NewConfigFromEnvironment(environmentPrefix ...string) *Config {
	helper := viper.New()
	if len(environmentPrefix) > 0 {
		helper.SetEnvPrefix(environmentPrefix[0])
	}
	return &Config{
		Host:                    NewEnvironmentKey(helper, "HOST", DefaultHost).GetString(),
		LivenessCheckInterval:   NewEnvironmentKey(helper, "LIVENESS_CHECK_INTERVAL", DefaultLivenessCheckInterval).GetDuration(),
		LivenessCheckMethod:     NewEnvironmentKey(helper, "LIVENESS_CHECK_METHOD", DefaultLivenessCheckMethod).GetString(),
		LivenessCheckPath:       NewEnvironmentKey(helper, "LIVENESS_CHECK_PATH", DefaultLivenessCheckPath).GetString(),
		LivenessCheckStatusCode: NewEnvironmentKey(helper, "LIVENESS_CHECK_STATUS_CODE", DefaultLivenessCheckStatusCode).GetInt(),
		LivenessCheckTimeout:    NewEnvironmentKey(helper, "LIVENESS_CHECK_TIMEOUT", DefaultLivenessCheckTimeout).GetDuration(),
		MaxHeaderBytes:          NewEnvironmentKey(helper, "MAX_HEADER_BYTES", DefaultMaxHeaderBytes).GetInt(),
		Port:                    NewEnvironmentKey(helper, "PORT", DefaultPort).GetString(),
		TimeoutIdle:             NewEnvironmentKey(helper, "TIMEOUT_IDLE", DefaultTimeout).GetDuration(),
		TimeoutRead:             NewEnvironmentKey(helper, "TIMEOUT_READ", DefaultTimeout).GetDuration(),
		TLSCertificatePath:      NewEnvironmentKey(helper, "TLS_CERTIFICATE_PATH", DefaultTLSCertificatePath).GetString(),
		TLSKeyPath:              NewEnvironmentKey(helper, "TLS_KEY_PATH", DefaultTLSKeyPath).GetString(),
	}
}

type Config struct {
	Host                    string
	LivenessCheckPath       string
	LivenessCheckMethod     string
	LivenessCheckStatusCode int
	LivenessCheckTimeout    time.Duration
	LivenessCheckInterval   time.Duration
	MaxHeaderBytes          int
	OnError                 func(error)
	OnListening             func()
	OnListeningTimeout      func()
	OnShutdown              func()
	OnStarting              func()
	Port                    string
	TimeoutIdle             time.Duration
	TimeoutRead             time.Duration
	TLSCertificatePath      string
	TLSKeyPath              string
}

// GetAddr returns a string that's usable as the Addr property of the
// http.Server struct
func (config *Config) GetAddr() string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}

// GetProtocol checks if TLS is available and returns the appropriate
// protocol
func (config *Config) GetProtocol() string {
	if config.HasValidTLS() {
		return "https"
	}
	return "http"
}

// GetURL returns a string that's friendly to link finders
func (config *Config) GetURL() string {
	return fmt.Sprintf("%s://%s:%s", config.GetProtocol(), config.Host, config.Port)
}

// HasValidTLS returns true when the TLS configurations are valid and working
// as expected, false otherwise
func (config *Config) HasValidTLS() bool {
	return (stringDefined(config.TLSCertificatePath) == nil) &&
		(stringDefined(config.TLSKeyPath) == nil) &&
		(fileExists(config.TLSCertificatePath) == nil) &&
		(fileExists(config.TLSKeyPath) == nil) &&
		(tlsCertKeyMatches(config.TLSCertificatePath, config.TLSKeyPath) == nil)
}

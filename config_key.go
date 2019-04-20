package server

import (
	"time"

	"github.com/spf13/viper"
)

type Key interface {
	GetBool() bool
	GetDuration() time.Duration
	GetInt() int
	GetString() string
}

// NewEnvironmentKey is a helper function to create a new
// EnvironmentKey struct
func NewEnvironmentKey(helperInstance *viper.Viper, name string, defaultValue interface{}) *EnvironmentKey {
	return &EnvironmentKey{defaultValue, name, helperInstance}
}

// EnvironmentKey is used for retrieving values from the
// environment variables
type EnvironmentKey struct {
	DefaultValue interface{}
	Name         string
	helper       *viper.Viper
}

// GetBool retrieves the value as a duration
func (key *EnvironmentKey) GetBool() bool {
	key.registerBinding()
	return key.helper.GetBool(key.Name)
}

// GetDuration retrieves the value as a duration
func (key *EnvironmentKey) GetDuration() time.Duration {
	key.registerBinding()
	return key.helper.GetDuration(key.Name)
}

// GetInt retrieves the value as an integer
func (key *EnvironmentKey) GetInt() int {
	key.registerBinding()
	return key.helper.GetInt(key.Name)
}

// GetString retrieves the value as a string
func (key *EnvironmentKey) GetString() string {
	key.registerBinding()
	return key.helper.GetString(key.Name)
}

func (key *EnvironmentKey) registerBinding() {
	key.helper.SetDefault(key.Name, key.DefaultValue)
	key.helper.BindEnv(key.Name)
}

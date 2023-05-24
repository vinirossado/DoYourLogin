package configuration

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	value string
}

func MakeConfig(value string) *Config {
	return &Config{value: value}
}

func (cfg *Config) SetValue(value string) {
	cfg.value = value
}

func (cfg *Config) UpdateValueFromEnv(envname string) {
	val, exists := os.LookupEnv(envname)
	if !exists {
		log.Panicf("%s is not set", envname)
	}
	cfg.value = val
}

func (cfg *Config) ValueAsString() string {
	return cfg.value
}

func (cfg *Config) ValueAsByte() []byte {
	return []byte(cfg.value)
}

func (cfg *Config) ValueAsFloat() float64 {
	val, err := strconv.ParseFloat(cfg.value, 64)
	if err != nil {
		log.Panicf("'%s' is not a valid float64 type", cfg.value)
	}

	return val
}

func (cfg *Config) ValueAsInt() int {
	val, err := strconv.Atoi(cfg.value)
	if err != nil {
		log.Panicf("'%s' is not a valid int type", cfg.value)
	}

	return val
}

func (cfg *Config) ValueAsBool() bool {
	val, err := strconv.ParseBool(cfg.value)
	if err != nil {
		log.Panicf("'%s' is not a valid bool type", cfg.value)
	}

	return val
}

func (cfg *Config) ValueAsDuration() time.Duration {
	val, err := strconv.ParseInt(cfg.value, 10, 64)
	if err == nil {
		log.Panicf("'%s' is not a valid duration type", cfg.value)
	}

	return time.Duration(val)
}

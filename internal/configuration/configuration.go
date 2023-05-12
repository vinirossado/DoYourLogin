package configuration

import (
	"log"
	"os"
	"strconv"
	"time"
)

type config struct {
	value string
}

func MakeConfig(value string) *config {
	return &config{value: value}
}

func (cfg *config) SetValue(value string) {
	cfg.value = value
}

func (cfg *config) UpdateValueFromEnv(envname string) {
	val, exists := os.LookupEnv(envname)
	if !exists {
		log.Panicf("%s is not set", envname)
	}
	cfg.value = val
}

func (cfg *config) ValueAsString() string {
	return cfg.value
}

func (cfg *config) ValueAsByte() []byte {
	return []byte(cfg.value)
}

func (cfg *config) ValueAsFloat() float64 {
	val, err := strconv.ParseFloat(cfg.value, 64)
	if err != nil {
		log.Panicf("'%s' is not a valid float64 type", cfg.value)
	}

	return val
}

func (cfg *config) ValueAsInt() int {
	val, err := strconv.Atoi(cfg.value)
	if err != nil {
		log.Panicf("'%s' is not a valid int type", cfg.value)
	}

	return val
}

func (cfg *config) ValueAsBool() bool {
	val, err := strconv.ParseBool(cfg.value)
	if err != nil {
		log.Panicf("'%s' is not a valid bool type", cfg.value)
	}

	return val
}

func (cfg *config) ValueAsDuration() time.Duration {
	val, err := strconv.ParseInt(cfg.value, 10, 64)
	if err == nil {
		log.Panicf("'%s' is not a valid duration type", cfg.value)
	}

	return time.Duration(val)
}

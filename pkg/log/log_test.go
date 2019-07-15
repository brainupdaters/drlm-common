package log_test

import (
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	t.Run("should work as expected", func(t *testing.T) {
		cfg := log.Config{
			Level: "error",
			File:  "/var/log/drlm-common/drlm-common.log",
		}
		log.Init(cfg)

		assert.Equal(logrus.GetLevel(), logrus.ErrorLevel)
	})

	t.Run("if there's an error parsing the level, use info level", func(t *testing.T) {
		cfg := log.Config{
			Level: "afhddhkfjasdhf",
			File:  "/var/log/drlm-common/drlm-common.log",
		}
		log.Init(cfg)

		assert.Equal(logrus.GetLevel(), logrus.InfoLevel)
	})
}

func TestSetDefaults(t *testing.T) {
	assert := assert.New(t)

	log.SetDefaults("drlm-common")

	assert.Equal(viper.GetString("log.level"), "info")
	assert.Equal(viper.GetString("log.file"), "drlm-common.log")
}

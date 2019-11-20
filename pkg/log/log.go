// SPDX-License-Identifier: AGPL-3.0-only

/*
 * Copyright (C) 2019 DRLM Project
 * Authors: DRLM Common authors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package log

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config is the configuration of the log package
type Config struct {
	// Level is the level that the logger is going to log
	Level string

	// File is the file where the logger is going to write the logs
	File string
}

// Init initializes the log package
func Init(cfg Config) {
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	pathMap := lfshook.PathMap{
		logrus.TraceLevel: cfg.File,
		logrus.DebugLevel: cfg.File,
		logrus.InfoLevel:  cfg.File,
		logrus.WarnLevel:  cfg.File,
		logrus.ErrorLevel: cfg.File,
		logrus.FatalLevel: cfg.File,
		logrus.PanicLevel: cfg.File,
	}

	logrus.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.TextFormatter{},
	))
}

// SetDefaults sets the default configurations for Viper
func SetDefaults(app string) {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file", app+".log")
}

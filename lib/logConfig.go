package lib

import (
	"errors"
)

type ConfFileWriter struct {
	On              bool
	LogPath         string
	RotateLogPath   string
	WfLogPath       string
	RotateWfLogPath string
}

type ConfConsoleWriter struct {
	On    bool
	Color bool
}

type LogConfig struct {
	Level string
	FW    ConfFileWriter
	CW    ConfConsoleWriter
}

func SetupLogInstanceWithConf(lc LogConfig, logger *Logger) (err error) {
	if lc.FW.On {
		if len(lc.FW.LogPath) > 0 {
			w := NewFileWriter()
			w.SetFileName(lc.FW.LogPath)
			w.SetPathPattern(lc.FW.RotateLogPath)
			w.SetLogLevelFloor(TRACE)
			if len(lc.FW.WfLogPath) > 0 {
				w.SetLogLevelCeil(INFO)
			} else {
				w.SetLogLevelCeil(ERROR)
			}
			logger.Register(w)
		}

		if len(lc.FW.WfLogPath) > 0 {
			wfw := NewFileWriter()
			wfw.SetFileName(lc.FW.WfLogPath)
			wfw.SetPathPattern(lc.FW.RotateWfLogPath)
			wfw.SetLogLevelFloor(WARNING)
			wfw.SetLogLevelCeil(ERROR)
			logger.Register(wfw)
		}
	}

	if lc.CW.On {
		w := NewConsoleWriter()
		w.SetColor(lc.CW.Color)
		logger.Register(w)
	}
	switch lc.Level {
	case "trace":
		logger.SetLevel(TRACE)

	case "debug":
		logger.SetLevel(DEBUG)

	case "info":
		logger.SetLevel(INFO)

	case "warning":
		logger.SetLevel(WARNING)

	case "error":
		logger.SetLevel(ERROR)

	case "fatal":
		logger.SetLevel(FATAL)

	default:
		err = errors.New("Invalid log level")
	}
	return
}

func SetupDefaultLogWithConf(lc LogConfig) (err error) {
	defaultLoggerInit()
	return SetupLogInstanceWithConf(lc, logger_default)
}

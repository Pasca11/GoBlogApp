package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"path"
)

type logLevel string

const (
	LOCAL logLevel = "LOCAL"
	STAGE logLevel = "STAGE"
	PROD  logLevel = "PROD"
)

type Logger interface {
	DebugContext(ctx context.Context, msg string, args ...interface{})
	InfoContext(ctx context.Context, msg string, args ...interface{})
	WarnContext(ctx context.Context, msg string, args ...interface{})
	ErrorContext(ctx context.Context, msg string, args ...interface{})
}

type Config struct {
	Path    string   `yaml:"path"`
	Level   logLevel `yaml:"level"`
	Source  bool     `yaml:"source"`
	Graylog GrayLog  `yaml:"graylog"`
}

type GrayLog struct {
	Use      bool   `yaml:"use"`
	ConnType string `yaml:"conn_type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func New(cfg *Config) (Logger, error) {
	handlerOption := &slog.HandlerOptions{
		AddSource: cfg.Source,
	}
	var output io.Writer
	var err error

	switch cfg.Level {
	case LOCAL:
		handlerOption.Level = slog.LevelDebug
	case STAGE:
		handlerOption.Level = slog.LevelInfo
		output, err = loadFileOutput(cfg)
	case PROD:
		handlerOption.Level = slog.LevelError
		output, err = loadFileOutput(cfg)
	}
	if err != nil {
		panic(err)
	}

	if cfg.Graylog.Use {
		output = connectGrayLog(cfg)
	}

	log := slog.New(slog.NewJSONHandler(output, handlerOption))
	return log, nil
}

func connectGrayLog(cfg *Config) io.Writer {
	conn, err := net.Dial(cfg.Graylog.ConnType, fmt.Sprintf("%s:%s", cfg.Graylog.Host, cfg.Graylog.Port))
	if err != nil {
		panic(err)
	}
	return conn
}

func loadFileOutput(cfg *Config) (io.Writer, error) {
	err := os.MkdirAll(path.Dir(cfg.Path), 0755)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(cfg.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

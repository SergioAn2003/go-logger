// Package logger реализует логгер
package logger

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level кастомный уровень логирования
type Level int

// ToZapLevel преобразовывает Level в zapcore.Level
//
//	 var level logger.Level
//	 level = 0
//	 zapLevel, err := level.ToZapLevel()
//	 if err != nil{
//		panic(err)
//	 }

const (
	DebugLevel Level = -4 // DebugLevel уровень debug
	InfoLevel  Level = 0  // InfoLevel уровень info
	WarnLevel  Level = 4  // WarnLevel уровень warn
	ErrorLevel Level = 8  // ErrorLevel уровень error
	FatalLevel Level = 12 // FatalLevel уровень fatal
)

func (l Level) ToZapLevel() (zapcore.Level, error) {
	switch l {
	case DebugLevel:
		return zapcore.DebugLevel, nil
	case InfoLevel:
		return zapcore.InfoLevel, nil
	case WarnLevel:
		return zapcore.WarnLevel, nil
	case ErrorLevel:
		return zapcore.ErrorLevel, nil
	case FatalLevel:
		return zapcore.FatalLevel, nil
	default:
		return 0, ErrInvalidLevel
	}
}

// logger структура логгера которая реализует интерфейс Logger
type logger struct {
	log     *zap.Logger
	core    zapcore.Core
	encoder zapcore.Encoder
	outputs []zapcore.WriteSyncer
	fields  []zapcore.Field
	mode    string
}

// New функция конструктор,  принимает тип логгера.
// Тип может быть "dev", "prod" или "mock". Возвращает логгер и ошибку
//
//	log, err:= logger.New("dev")
func New(mode string) (Logger, error) {
	var (
		log logger
		err error
	)

	mode = strings.ToLower(mode)

	switch mode {
	case Dev:
		log, err = newDev()
	case Prod:
		log, err = newProd()
	case Mock:
		log = newMock()
	default:
		return nil, ErrInvalidMode
	}

	if err != nil {
		return nil, err
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

		for {
			s := <-c
			log.Debug(fmt.Sprintf("receive signal: %s", s.String()))
			if goErr := log.log.Sync(); goErr != nil {
				log.Error(goErr.Error())
				return
			}
		}
	}()

	return &log, nil
}

// newMock конструктор для тестового пустого логгера,
// возвращает логгер и ошибку
func newMock() logger {
	return logger{
		log:     zap.NewNop(),
		core:    nil,
		encoder: nil,
		outputs: nil,
		fields:  nil,
		mode:    Mock,
	}
}

// newDev конструктор для локального логгера,
// возвращает логгер для разработки и ошибку
func newDev() (logger, error) {
	output := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewDevelopmentEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.TimeKey = TimestampKey
	encoderConfig.LevelKey = LevelKey
	encoderConfig.MessageKey = MessageKey

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	level, err := DebugLevel.ToZapLevel()
	if err != nil {
		return logger{}, err
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, output, level),
	)

	return logger{
		log:     zap.New(core),
		core:    core,
		encoder: encoder,
		outputs: []zapcore.WriteSyncer{output},
		fields:  nil,
		mode:    Dev,
	}, nil
}

// newProd конструктор для продакшн логгера,
// возвращает продакшн логгер и ошибку
func newProd() (logger, error) {
	output := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = TimestampKey
	encoderConfig.LevelKey = LevelKey
	encoderConfig.MessageKey = MessageKey

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	level, err := InfoLevel.ToZapLevel()
	if err != nil {
		return logger{}, err
	}

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, output, level),
	)

	return logger{
		log:     zap.New(core),
		core:    core,
		encoder: jsonEncoder,
		outputs: []zapcore.WriteSyncer{output},
		fields:  nil,
		mode:    Prod,
	}, nil
}

// Debug пишет лог уровня debug
// принимает сообщение
//
//	log.Debug("hello world")
func (l *logger) Debug(msg string) {
	l.log.Debug(msg)
}

// DebugF пишет форматированный лог уровня debug
// принимает шаблон и аргументы
//
//	log.DebugF("hello %s", "world")
func (l *logger) DebugF(format string, args ...any) {
	l.log.Debug(fmt.Sprintf(format, args...))
}

// DebugW пишет лог уровня debug с аргументами
// принимает сообщение и аргументы в виде мапы
//
//	log.DebugW("hello world", map[string]any{"key": "value"})
func (l *logger) DebugW(msg string, args map[string]any) {
	l.log.Debug(msg, convertMapToFields(args)...)
}

// Info пишет лог уровня info
// принимает сообщение
//
//	log.Info("hello world")
func (l *logger) Info(msg string) {
	l.log.Info(msg)
}

// InfoF пишет форматированный лог уровня info
// принимает шаблон и аргументы
//
//	log.InfoF("hello %s", "world")
func (l *logger) InfoF(format string, args ...any) {
	l.log.Info(fmt.Sprintf(format, args...))
}

// InfoW пишет лог уровня info с аргументами
// принимает сообщение и аргументы в виде мапы
//
//	log.InfoW("hello world", map[string]any{"key": "value"})
func (l *logger) InfoW(msg string, args map[string]any) {
	l.log.Info(msg, convertMapToFields(args)...)
}

// Warn пишет лог уровня warn
// принимает сообщение
//
//	log.Warn("hello world")
func (l *logger) Warn(msg string) {
	l.log.Warn(msg)
}

// WarnF пишет форматированный лог уровня warn
// принимает шаблон и аргументы
//
//	log.WarnF("hello %s", "world")
func (l *logger) WarnF(format string, args ...any) {
	l.log.Warn(fmt.Sprintf(format, args...))
}

// WarnW пишет лог уровня warn с аргументами
// принимает сообщение и аргументы в виде мапы
//
//	log.WarnW("hello world", map[string]any{"key": "value"})
func (l *logger) WarnW(msg string, args map[string]any) {
	l.log.Warn(msg, convertMapToFields(args)...)
}

// Error пишет лог уровня error
// принимает сообщение
//
//	log.Error("hello world")
func (l *logger) Error(msg string) {
	l.log.Error(msg)
}

// ErrorF пишет форматированный лог уровня error
// принимает шаблон и аргументы
//
//	log.ErrorF("hello %s", "world")
func (l *logger) ErrorF(format string, args ...any) {
	l.log.Error(fmt.Sprintf(format, args...))
}

// ErrorW пишет лог уровня error с аргументами
// принимает сообщение и аргументы в виде мапы
//
//	log.ErrorW("hello world", map[string]any{"key": "value"})
func (l *logger) ErrorW(msg string, args map[string]any) {
	l.log.Error(msg, convertMapToFields(args)...)
}

// Fatal пишет лог уровня fatal и вызывает панику
// принимает сообщение
//
//	log.Fatal("hello world")
func (l *logger) Fatal(msg string) {
	if l.mode == Mock {
		return
	}

	ce := l.log.Check(zap.PanicLevel, msg)
	ce.Level = zap.FatalLevel
	ce.Write()
}

// FatalF пишет форматированный лог уровня fatal и вызывает панику
// принимает шаблон и аргументы
//
//	log.FatalF("hello %s", "world")
func (l *logger) FatalF(format string, args ...any) {
	if l.mode == Mock {
		return
	}

	ce := l.log.Check(zap.PanicLevel, fmt.Sprintf(format, args...))
	ce.Level = zap.FatalLevel
	ce.Write()
}

// FatalW пишет лог уровня fatal с аргументами и вызывает панику
// принимает сообщение и аргументы в виде мапы
//
//	log.FatalW("hello world", map[string]any{"key": "value"})
func (l *logger) FatalW(msg string, args map[string]any) {
	if l.mode == Mock {
		return
	}

	l.log = l.log.With(convertMapToFields(args)...)

	ce := l.log.Check(zap.PanicLevel, msg)
	ce.Level = zap.FatalLevel
	ce.Write()
}

// WithAttrs создает новый логгер с аргументами
// принимает аргументы в виде мапы
//
// log := log.WithAttrs(map[string]any{"key": "value"})
// log.Debug("hello world")
func (l *logger) WithAttrs(args map[string]any) Logger {
	fields := convertMapToFields(args)

	return &logger{
		log:     l.log.With(fields...),
		core:    l.core,
		encoder: l.encoder,
		outputs: l.outputs,
		fields:  fields,
	}
}

// SetLoggerLevel меняет уровень логирования
//
//	log := logger.New("prod")
//	log.SetLevel(logger.Debug)
//	log.Debug("hello world")
func (l *logger) SetLoggerLevel(level Level) error {
	zapLevel, err := level.ToZapLevel()
	if err != nil {
		return err
	}

	l.core = zapcore.NewCore(l.encoder, zapcore.NewMultiWriteSyncer(l.outputs...), zapLevel)
	l.log = zap.New(l.core)

	if l.fields != nil {
		l.log = l.log.With(l.fields...)
	}

	return nil
}

// convertMapToFields конвертирует map[string]any в []zap.Field
func convertMapToFields(args map[string]any) []zap.Field {
	fields := make([]zap.Field, 0, len(args))

	for k, v := range args {
		fields = append(fields, zap.Any(k, v))
	}

	return fields
}

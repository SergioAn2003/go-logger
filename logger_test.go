package logger_test

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/SergioAn2003/go-logger"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func BenchmarkProdDebug(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("hello world")
	}
}

func BenchmarkProdDebugF(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.DebugF("hello %s", "world")
	}
}

func BenchmarkProdDebugW(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.DebugW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkProdInfo(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info("hello world")
	}
}

func BenchmarkProdInfoF(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.InfoF("hello %s", "world")
	}
}

func BenchmarkProdInfoW(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.InfoW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkProdWarn(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Warn("hello world")
	}
}

func BenchmarkProdWarnF(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WarnF("hello %s", "world")
	}
}

func BenchmarkProdWarnW(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WarnW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkProdError(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Error("hello world")
	}
}

func BenchmarkProdErrorF(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		log.ErrorF("hello %s", "world")
	}
}

func BenchmarkProdErrorW(b *testing.B) {
	log, err := logger.New("prod")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkDevDebug1(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("hello world")
	}
}

func BenchmarkDevDebugF(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.DebugF("hello %s", "world")
	}
}

func BenchmarkDevDebugW(b *testing.B) {
	log, err := logger.New("dev")

	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.DebugW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkDevInfo1(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info("hello world")
	}
}

func BenchmarkDevInfoF(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.InfoF("hello %s", "world")
	}
}

func BenchmarkDevInfoW(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.InfoW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkDevWarn1(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Warn("hello world")
	}
}

func BenchmarkDevWarnF(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WarnF("hello %s", "world")
	}
}

func BenchmarkDevWarnW(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WarnW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkDevError1(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Error("hello world")
	}
}

func BenchmarkDevErrorF(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorF("hello %s", "world")
	}
}

func BenchmarkDevErrorW(b *testing.B) {
	log, err := logger.New("dev")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkMockDebug(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("hello world")
	}
}

func BenchmarkMockDebugF(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.DebugF("hello %s", "world")
	}
}

func BenchmarkMockDebugW(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.DebugW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkMockInfo(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info("hello world")
	}
}

func BenchmarkMockInfoF(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.InfoF("hello %s", "world")
	}
}

func BenchmarkMockInfoW(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.InfoW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkMockWarn(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Warn("hello world")
	}
}

func BenchmarkMockWarnF(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WarnF("hello %s", "world")
	}
}

func BenchmarkMockWarnW(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.WarnW("hello world", map[string]any{"key": "value"})
	}
}

func BenchmarkMockError(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Error("hello world")
	}
}

func BenchmarkMockErrorF(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorF("hello %s", "world")
	}
}

func BenchmarkMockErrorW(b *testing.B) {
	log, err := logger.New("mock")
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorW("hello world", map[string]any{"key": "value"})
	}
}

func TestDevLogger(t *testing.T) {
	r := require.New(t)

	stdout := os.Stdout

	const (
		debugLevel = "\033[35mDEBUG\033[0m"
		infoLevel  = "\033[34mINFO\033[0m"
		warnLevel  = "\033[33mWARN\033[0m"
		errorLevel = "\033[31mERROR\033[0m"
		fatalLevel = "\033[31mFATAL\033[0m"

		wantMsg = "hello world"
	)

	tests := []struct {
		name      string
		logFunc   func(logger.Logger, string)
		logFuncF  func(logger.Logger, string, ...any)
		logFuncW  func(logger.Logger, string, map[string]any)
		wantLevel string
		wantMsg   string
		wantArgs  map[string]any
	}{
		{
			name:      "Debug",
			logFunc:   logger.Logger.Debug,
			wantLevel: debugLevel,
			wantMsg:   wantMsg,
		},
		{
			name:      "DebugF",
			logFuncF:  logger.Logger.DebugF,
			wantLevel: debugLevel,
			wantMsg:   wantMsg,
		},
		{
			name:      "DebugW",
			logFuncW:  logger.Logger.DebugW,
			wantLevel: "\033[35mDEBUG\033[0m",
			wantMsg:   "hello world",
			wantArgs:  map[string]any{"key": "value"},
		},
		{
			name:      "Info",
			logFunc:   logger.Logger.Info,
			wantLevel: "\033[34mINFO\033[0m",
			wantMsg:   "hello world",
		},
		{
			name:      "InfoF",
			logFuncF:  logger.Logger.InfoF,
			wantLevel: "\033[34mINFO\033[0m",
			wantMsg:   "hello world",
		},
		{
			name:      "InfoW",
			logFuncW:  logger.Logger.InfoW,
			wantLevel: "\033[34mINFO\033[0m",
			wantMsg:   "hello world",
			wantArgs:  map[string]any{"key": "value"},
		},
		{
			name:      "Warn",
			logFunc:   logger.Logger.Warn,
			wantLevel: "\033[33mWARN\033[0m",
			wantMsg:   "hello world",
		},
		{
			name:      "WarnF",
			logFuncF:  logger.Logger.WarnF,
			wantLevel: "\033[33mWARN\033[0m",
			wantMsg:   "hello world",
		},
		{
			name:      "WarnW",
			logFuncW:  logger.Logger.WarnW,
			wantLevel: "\033[33mWARN\033[0m",
			wantMsg:   "hello world",
			wantArgs:  map[string]any{"key": "value"},
		},
		{
			name:      "Error",
			logFunc:   logger.Logger.Error,
			wantLevel: "\033[31mERROR\033[0m",
			wantMsg:   "hello world",
		},
		{
			name:      "ErrorF",
			logFuncF:  logger.Logger.ErrorF,
			wantLevel: "\033[31mERROR\033[0m",
			wantMsg:   "hello world",
		},
		{
			name:      "ErrorW",
			logFuncW:  logger.Logger.ErrorW,
			wantLevel: "\033[31mERROR\033[0m",
			wantMsg:   "hello world",
			wantArgs:  map[string]any{"key": "value"},
		},
		{
			name:      "Fatal",
			logFunc:   logger.Logger.Fatal,
			wantLevel: "\033[31mFATAL\033[0m",
			wantMsg:   "hello world",
		},
		{
			name:      "FatalF",
			logFuncF:  logger.Logger.FatalF,
			wantLevel: "\033[31mFATAL\033[0m",
			wantMsg:   "hello world",
		},
		{
			name:      "FatalW",
			logFuncW:  logger.Logger.FatalW,
			wantLevel: "\033[31mFATAL\033[0m",
			wantMsg:   "hello world",
			wantArgs:  map[string]any{"key": "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tempOut, err := os.CreateTemp("", "test*.log")
			if err != nil {
				r.NoError(err)
			}

			os.Stdout = tempOut

			defer func() {
				_ = tempOut.Close()
				_ = os.Remove(tempOut.Name())
				os.Stdout = stdout

				if tt.wantLevel == fatalLevel {
					var panicOccurred bool

					if rec := recover(); rec != nil {
						panicOccurred = true
					}

					r.True(panicOccurred)
				}
			}()

			log, err := logger.New("dev")
			r.NoError(err)

			switch {
			case tt.logFunc != nil:
				tt.logFunc(log, tt.wantMsg)
			case tt.logFuncF != nil:
				tt.logFuncF(log, "hello %s", "world")
			case tt.logFuncW != nil:
				tt.logFuncW(log, tt.wantMsg, map[string]any{"key": "value"})
			}

			out, err := os.ReadFile(tempOut.Name())
			r.NoError(err)

			r.Contains(string(out), tt.wantMsg)
			r.Contains(string(out), tt.wantLevel)

			if tt.logFuncW != nil {
				re := regexp.MustCompile(`{[^{}]*}`)
				jsonArgs := re.FindString(string(out))

				actualArgs := map[string]any{}

				err := json.Unmarshal([]byte(jsonArgs), &actualArgs)
				r.NoError(err)

				r.Equal(tt.wantArgs, actualArgs)
			}
		})
	}
}

func TestProdLogger(t *testing.T) {
	r := require.New(t)

	stdout := os.Stdout

	const (
		levelInfo  = "info"
		levelWarn  = "warn"
		levelError = "error"
		levelFatal = "fatal"

		wantMsg = "hello world"
	)

	tests := []struct {
		name     string
		logFunc  func(logger.Logger, string)
		logFuncF func(logger.Logger, string, ...any)
		logFuncW func(logger.Logger, string, map[string]any)
		wantData map[string]string
	}{
		{
			name:     "Debug",
			logFunc:  logger.Logger.Debug,
			wantData: map[string]string{},
		},
		{
			name:     "DebugF",
			logFuncF: logger.Logger.DebugF,
			wantData: map[string]string{},
		},
		{
			name:     "DebugW",
			logFuncW: logger.Logger.DebugW,
			wantData: map[string]string{},
		},
		{
			name:    "Info",
			logFunc: logger.Logger.Info,
			wantData: map[string]string{
				"level": levelInfo,
				"msg":   wantMsg,
			},
		},
		{
			name:     "InfoF",
			logFuncF: logger.Logger.InfoF,
			wantData: map[string]string{
				"level": levelInfo,
				"msg":   wantMsg,
			},
		},
		{
			name:     "InfoW",
			logFuncW: logger.Logger.InfoW,
			wantData: map[string]string{
				"level": levelInfo,
				"msg":   wantMsg,
				"key":   "value",
			},
		},
		{
			name:    "Warn",
			logFunc: logger.Logger.Warn,
			wantData: map[string]string{
				"level": levelWarn,
				"msg":   wantMsg,
			},
		},
		{
			name:     "WarnF",
			logFuncF: logger.Logger.WarnF,
			wantData: map[string]string{
				"level": levelWarn,
				"msg":   wantMsg,
			},
		},
		{
			name:     "WarnW",
			logFuncW: logger.Logger.WarnW,
			wantData: map[string]string{
				"level": levelWarn,
				"msg":   wantMsg,
				"key":   "value",
			},
		},
		{
			name:    "Error",
			logFunc: logger.Logger.Error,
			wantData: map[string]string{
				"level": levelError,
				"msg":   wantMsg,
			},
		},
		{
			name:     "ErrorF",
			logFuncF: logger.Logger.ErrorF,
			wantData: map[string]string{
				"level": levelError,
				"msg":   wantMsg,
			},
		},
		{
			name:     "ErrorW",
			logFuncW: logger.Logger.ErrorW,
			wantData: map[string]string{
				"level": levelError,
				"msg":   wantMsg,
				"key":   "value",
			},
		},
		{
			name:    "Fatal",
			logFunc: logger.Logger.Fatal,
			wantData: map[string]string{
				"level": levelFatal,
				"msg":   wantMsg,
			},
		},
		{
			name:     "FatalF",
			logFuncF: logger.Logger.FatalF,
			wantData: map[string]string{
				"level": levelFatal,
				"msg":   wantMsg,
			},
		},
		{
			name:     "FatalW",
			logFuncW: logger.Logger.FatalW,
			wantData: map[string]string{
				"level": levelFatal,
				"msg":   wantMsg,
				"key":   "value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tempOut, err := os.CreateTemp("", "test*.log")
			r.NoError(err)

			os.Stdout = tempOut

			defer func() {
				_ = tempOut.Close()
				_ = os.Remove(tempOut.Name())
				os.Stdout = stdout

				if tt.wantData["level"] == levelFatal {
					var panicOccurred bool

					if rec := recover(); rec != nil {
						panicOccurred = true
					}

					r.True(panicOccurred)
				}
			}()

			log, err := logger.New("prod")
			r.NoError(err)

			switch {
			case tt.logFunc != nil:
				tt.logFunc(log, tt.wantData["msg"])
			case tt.logFuncF != nil:
				tt.logFuncF(log, "hello %s", "world")
			case tt.logFuncW != nil:
				tt.logFuncW(log, tt.wantData["msg"], map[string]any{
					"key": "value",
				})
			}

			out, err := os.ReadFile(tempOut.Name())
			r.NoError(err)

			for k, v := range tt.wantData {
				r.Contains(string(out), fmt.Sprintf("%q:%q", k, v))
			}
		})
	}
}

func TestMockLogger(t *testing.T) {
	r := require.New(t)

	stdout := os.Stdout

	tempOut, err := os.CreateTemp("", "test.log")
	r.NoError(err)

	os.Stdout = tempOut

	defer func() {
		_ = tempOut.Close()
		_ = os.Remove(tempOut.Name())
		os.Stdout = stdout
	}()

	log, err := logger.New("mock")
	r.NoError(err)

	log.Info("hello world")

	out, err := os.ReadFile(tempOut.Name())
	r.NoError(err)

	r.Empty(string(out))
}

func TestSetLoggerLevel(t *testing.T) {
	r := require.New(t)

	stdout := os.Stdout

	const (
		debugLevel = "\033[35mDEBUG\033[0m"
		msg        = "hello world"
	)

	t.Run("level debug", func(t *testing.T) {
		tempOut, err := os.CreateTemp("", "test1.log")
		r.NoError(err)

		os.Stdout = tempOut

		defer func() {
			_ = tempOut.Close()
			_ = os.Remove(tempOut.Name())
			os.Stdout = stdout
		}()

		log, err := logger.New("dev")
		r.NoError(err)

		log.Debug(msg)

		out, err := os.ReadFile(tempOut.Name())
		r.NoError(err)

		r.Contains(string(out), debugLevel)
		r.Contains(string(out), msg)

		os.Stdout = stdout
		t.Run("level info", func(_ *testing.T) {
			tempOut, err = os.CreateTemp("", "test2.log")
			r.NoError(err)

			os.Stdout = tempOut

			err = log.SetLoggerLevel(logger.InfoLevel)
			r.NoError(err)

			log.Debug("test")

			out, err = os.ReadFile(tempOut.Name())
			r.NoError(err)

			r.Empty(string(out))
		})
	})
}

func TestWithAttrs(t *testing.T) {
	r := require.New(t)

	tests := []struct {
		name         string
		env          string
		args         map[string]any
		wantDevArgs  map[string]any
		wantProdArgs string
	}{
		{
			name:         "prod with attrs",
			env:          "prod",
			wantProdArgs: "\"key\":\"value\"",
			args:         map[string]any{"key": "value"},
		},
		{
			name: "dev with attrs",
			env:  "dev",
			wantDevArgs: map[string]any{
				"key": "value",
			},
			args: map[string]any{"key": "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			stdout := os.Stdout

			tempOut, err := os.CreateTemp("", "test*.log")
			r.NoError(err)

			os.Stdout = tempOut

			defer func() {
				_ = tempOut.Close()
				_ = os.Remove(tempOut.Name())
				os.Stdout = stdout
			}()

			log, err := logger.New(tt.env)
			r.NoError(err)

			log = log.WithAttrs(tt.args)
			log.Info("test")

			out, err := os.ReadFile(tempOut.Name())
			r.NoError(err)

			switch tt.env {
			case "prod":
				r.Contains(string(out), tt.wantProdArgs)
			case "dev":
				re := regexp.MustCompile(`{[^{}]*}`)
				jsonArgs := re.FindString(string(out))

				actualArgs := map[string]any{}

				err = json.Unmarshal([]byte(jsonArgs), &actualArgs)
				r.NoError(err)

				r.Equal(tt.wantDevArgs, actualArgs)
			}
		})
	}
}

func TestWithAttrsAfterSetLoggerLevel(t *testing.T) {
	r := require.New(t)

	const (
		debugLevel = "\033[35mDEBUG\033[0m"
		infoLevel  = "\033[34mINFO\033[0m"

		wantMsg = "hello world"
	)

	wantAttrs := map[string]any{
		"key": "value",
	}

	stdout := os.Stdout
	t.Run("dev with attrs", func(_ *testing.T) {
		tempOut, err := os.CreateTemp("", "test*.log")
		r.NoError(err)

		os.Stdout = tempOut

		defer func() {
			_ = tempOut.Close()
			_ = os.Remove(tempOut.Name())
			os.Stdout = stdout
		}()

		log, err := logger.New("dev")
		r.NoError(err)

		log = log.WithAttrs(wantAttrs)

		log.Debug(wantMsg)

		out, err := os.ReadFile(tempOut.Name())
		r.NoError(err)

		jsonArgs := regexp.MustCompile(`{[^{}]*}`).FindString(string(out))

		actualArgs := map[string]any{}

		err = json.Unmarshal([]byte(jsonArgs), &actualArgs)
		r.NoError(err)

		r.Contains(string(out), debugLevel)
		r.Contains(string(out), wantMsg)
		r.Equal(wantAttrs, actualArgs)

		os.Stdout = stdout

		t.Run("after set level", func(_ *testing.T) {
			tempout, err := os.CreateTemp("", "test2.log")
			r.NoError(err)

			os.Stdout = tempout

			err = log.SetLoggerLevel(logger.InfoLevel)
			r.NoError(err)

			log.Info(wantMsg)

			out, err = os.ReadFile(tempOut.Name())
			r.NoError(err)

			jsonArgs = regexp.MustCompile(`{[^{}]*}`).FindString(string(out))

			actualArgs = map[string]any{}

			err = json.Unmarshal([]byte(jsonArgs), &actualArgs)
			r.NoError(err)

			r.Contains(string(out), infoLevel)
			r.Contains(string(out), wantMsg)
			r.Equal(wantAttrs, actualArgs)
		})
	})

	t.Run("prod with attrs", func(_ *testing.T) {
		tempOut, err := os.CreateTemp("", "test*.log")
		r.NoError(err)

		os.Stdout = tempOut

		defer func() {
			_ = tempOut.Close()
			_ = os.Remove(tempOut.Name())
			os.Stdout = stdout
		}()

		log, err := logger.New("prod")
		r.NoError(err)

		log = log.WithAttrs(wantAttrs)

		log.Info(wantMsg)

		out, err := os.ReadFile(tempOut.Name())
		r.NoError(err)

		for k, v := range wantAttrs {
			r.Contains(string(out), fmt.Sprintf("%q:%q", k, v))
		}

		r.Contains(string(out), wantMsg)

		t.Run("after set level", func(_ *testing.T) {
			tempout, err := os.CreateTemp("", "test2.log")
			r.NoError(err)

			os.Stdout = tempout

			err = log.SetLoggerLevel(logger.WarnLevel)
			r.NoError(err)

			log.Warn(wantMsg)

			out, err = os.ReadFile(tempOut.Name())
			r.NoError(err)

			for k, v := range wantAttrs {
				r.Contains(string(out), fmt.Sprintf("%q:%q", k, v))
			}

			r.Contains(string(out), wantMsg)
		})
	})
}

func TestNew(t *testing.T) {
	r := require.New(t)

	tests := []struct {
		name    string
		mode    string
		wantErr bool
	}{
		{
			name:    "dev",
			mode:    "dev",
			wantErr: false,
		},
		{
			name:    "prod",
			mode:    "prod",
			wantErr: false,
		},
		{
			name:    "mock",
			mode:    "mock",
			wantErr: false,
		},
		{
			name:    "unknown",
			mode:    "unknown",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			_, err := logger.New(tt.mode)
			if tt.wantErr {
				r.Error(err)
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestToZapLevel(t *testing.T) {
	r := require.New(t)

	tests := []struct {
		name      string
		level     logger.Level
		wantLevel zapcore.Level
		err       error
	}{
		{
			name:      "debug",
			level:     logger.DebugLevel,
			wantLevel: zapcore.DebugLevel,
			err:       nil,
		},
		{
			name:      "info",
			level:     logger.InfoLevel,
			wantLevel: zapcore.InfoLevel,
			err:       nil,
		},
		{
			name:      "warn",
			level:     logger.WarnLevel,
			wantLevel: zapcore.WarnLevel,
			err:       nil,
		},
		{
			name:      "error",
			level:     logger.ErrorLevel,
			wantLevel: zapcore.ErrorLevel,
			err:       nil,
		},
		{
			name:      "fatal",
			level:     logger.FatalLevel,
			wantLevel: zapcore.FatalLevel,
			err:       nil,
		},
		{
			name:      "invalid level",
			level:     logger.Level(100),
			wantLevel: 0,
			err:       logger.ErrInvalidLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			actualLevel, err := tt.level.ToZapLevel()
			r.Equal(tt.wantLevel, actualLevel)
			r.Equal(tt.err, err)
		})
	}
}

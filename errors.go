package logger

import (
	"errors"
)

var (
	// ErrInvalidMode невалидный режим логирования
	ErrInvalidMode = errors.New("invalid logger mode")

	// ErrInvalidLevel невалидный уровень логирования
	ErrInvalidLevel = errors.New("invalid logger level")
)

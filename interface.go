package logger

type Logger interface {
	Debug(msg string)                       // Debug пишет лог уровня debug
	DebugF(format string, args ...any)      // DebugF пишет форматированный лог уровня debug
	DebugW(msg string, args map[string]any) // DebugW пишет лог уровня debug с аргументами

	Info(msg string)                       // Info пишет лог уровня info
	InfoF(format string, args ...any)      // InfoF пишет форматированный лог уровня info
	InfoW(msg string, args map[string]any) // InfoW пишет лог уровня info с аргументами

	Warn(msg string)                       // Warn пишет лог уровня warn
	WarnF(format string, args ...any)      // WarnF пишет форматированный лог уровня warn
	WarnW(msg string, args map[string]any) // WarnW пишет лог уровня warn с аргументами

	Error(msg string)                       // Error пишет лог уровня error
	ErrorF(format string, args ...any)      // ErrorF пишет форматированный лог уровня error
	ErrorW(msg string, args map[string]any) // ErrorW пишет лог уровня error с аргументами

	Fatal(msg string)                       // Fatal пишет лог уровня fatal и вызывает панику
	FatalF(format string, args ...any)      // FatalF пишет форматированный лог уровня fatal и вызывает панику
	FatalW(msg string, args map[string]any) // FatalW пишет лог уровня fatal с аргументами и вызывает панику

	WithAttrs(args map[string]any) Logger // WithAttrs создает новый логгер с атрибутами

	SetLoggerLevel(level Level) error // SetLoggerLevel меняет уровень логирования
}

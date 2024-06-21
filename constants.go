package logger

const (
	Dev  = "dev"  // Dev режим логгера для локальной разработки
	Prod = "prod" // Prod режим логгера для продакшн сборки
	Mock = "mock" // Mock режим логгера для тестирования
)

const (
	TimestampKey = "ts"    // TimestampKey ключ для timestamp
	LevelKey     = "level" // LevelKey ключ для уровня логирования
	MessageKey   = "msg"   // MessageKey ключ для сообщения
)

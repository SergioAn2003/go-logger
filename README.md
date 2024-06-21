# Logger

Утилита для печати логов. Под капотом использует пакет `go.uber.org/zap`.
Имеет три режима работы: `dev`, `prod` и `mock`:
- `dev` выводит в консоль стилизованные логи всех уровней;
- `prod` выводит в консоль логи уровнем выже `Debog` в формате `json`;
- `mock` игнорирует все логи, используется в тестах.

<br>

# Доступные команды

Для работы с логгером необходимо установить зависимости в проекте, для этого запустите команду `make install-deps`.

<br>

# Установка

```bash
  go get -u github.com/SergioAn2003/go-logger
```

<br>

# Критерий приёмки Pull Request

- Все имеющиеся тесты прошли успешно;
- Весь основной функционал покрыт тестами;
- Все изменения в коде прошли проверку с помощью `golangci-lint`;
- Написана необходимая документация в формате `godoc` с примерами использования;
- При необходимости доработан `README.md`.

Перед отправкой Pull Request необходимо выполнить следующие команды:

```
make test
make lint
```

<br>

# Пример использования

```go
package main

func main() {
	log, err := logger.New("dev")
	if err != nil {
		panic(err)
	}

	log.Debug("Hello world")
	log.DebugF("Hello %s", "world")
	log.DebugW("Hello world", map[string]any{
		"key": "value",
	})

	log.Info("Hello world")
	log.InfoF("Hello %s", "world")
	log.InfoW("Hello world", map[string]any{
		"key": "value",
	})

	log.Warn("Hello world")
	log.WarnF("Hello %s", "world")
	log.WarnW("Hello world", map[string]any{
		"key": "value",
	})

	log.Error("Hello world")
	log.ErrorF("Hello %s", "world")
	log.ErrorW("Hello world", map[string]any{
		"key": "value",
	})

	// Вызовет панику
	log.Fatal("Hello world")
	log.FatalF("Hello %s", "world")
	log.FatalW("Hello world", map[string]any{
		"key": "value",
	})
}
```

Также можно инициализировать в FX DI Framework

```go
package main

func main() {
    fx.App(
        logger.Module,
    ).Run()
}
```

<br>

# Banchmark

| Режим логгера | Метод  | Скорость    | Кол-во аллокаций |
|---------------|--------|-------------|------------------|
| `Prod`        | Debug  | 6.882 ns/op | 0 allocs/op      |
| `Prod`        | DebugF | 110.8 ns/op | 2 allocs/op      |
| `Prod`        | DebugW | 225.2 ns/op | 3 allocs/op      |
| `Prod`        | Info   | 5082 ns/op  | 0 allocs/op      |
| `Prod`        | InfoF  | 5196 ns/op  | 2 allocs/op      |
| `Prod`        | InfoW  | 6022 ns/op  | 3 allocs/op      |
| `Prod`        | Warn   | 4968 ns/op  | 0 allocs/op      |
| `Prod`        | WarnF  | 5562 ns/op  | 2 allocs/op      |
| `Prod`        | WarnW  | 6188 ns/op  | 3 allocs/op      |
| `Prod`        | Error  | 5072 ns/op  | 0 allocs/op      |
| `Prod`        | ErrorF | 5475 ns/op  | 2 allocs/op      |
| `Prod`        | ErrorW | 6308 ns/op  | 3 allocs/op      |

| Режим логгера | Метод  | Скорость   | Кол-во аллокаций |
|---------------|--------|------------|------------------|
| `Dev`         | Debug  | 5516 ns/op | 3 allocs/op      |
| `Dev`         | DebugF | 5461 ns/op | 5 allocs/op      |
| `Dev`         | DebugW | 6841 ns/op | 6 allocs/op      |
| `Dev`         | Info   | 5165 ns/op | 3 allocs/op      |
| `Dev`         | InfoF  | 5417 ns/op | 5 allocs/op      |
| `Dev`         | InfoW  | 7017 ns/op | 6 allocs/op      |
| `Dev`         | Warn   | 5157 ns/op | 3 allocs/op      |
| `Dev`         | WarnF  | 5361 ns/op | 5 allocs/op      |
| `Dev`         | WarnW  | 7025 ns/op | 6 allocs/op      |
| `Dev`         | Error  | 5465 ns/op | 3 allocs/op      |
| `Dev`         | ErrorF | 5474 ns/op | 5 allocs/op      |
| `Dev`         | ErrorW | 7009 ns/op | 6 allocs/op      |

<br>

# Доступные команды

| Команда             | Описание                                                       |
|---------------------|----------------------------------------------------------------|
| `make install-deps` | Устанавливает необходимые зависимости                          |
| `make doc`          | Открывает документацию с помощью утилиты `godoc` 9000 на порту |
| `lint`              | Запускает проверку кода с помощью `golangci-lint`              |
| `test`              | Запускает тесты с помощью `go test` с флагом на race detection |
| `test-cover`        | Выводит процент покрытия проекта тестами                       |
| `test-cover-svg`    | Создает карту покрытия проекта тастами в формате `svg`         |
| `test-cover-html`   | Открывает страницу с покрытием тестами в формате `html`        |
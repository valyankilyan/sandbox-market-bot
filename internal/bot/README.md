## Telegram BOT

Микросервис, управляющий ботом.

В основной дирректории хранятся интерфейсы для вложенных пакетов и методы, реализующие общение бота с пользователями

### TBot

``` golang
type TBot struct {
	token  string
	server Server
	invest Invest
}
```

Структура бота представляет из себя
- Телеграм токен
- `server` Интерфейс общения с сервером базы данных ([реализация](server_client/README.md))
- `invest` Интерфейс общения с микросервисом [myinvest](../myinvest/README.MD) ([реализация](myinvest_client/README.md))

### В кратце
- `basic_commands.go` - базовые команды: \start, \help...
- `currencies.go` - всё для вывода доступных валют и их курса
- `datastruct.go` - основные структуры микросервиса
- `get_updates.go` - каждые N ms ([config](../../config)) подтягивает обновления с апишики телеграма и запихивает их в канал, делящий с `HandleMessages`, который в свою очередь обрабатывает возможные кейсы и отвечает пользователю.
- `send_message.go` - метод для посылки сообщений 
- `tinkoff_account.go` - методы связанные с тинькофф акаунтом (`update_token`, `payin`)
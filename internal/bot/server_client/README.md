## Server Client
- `server.go` - структура  для общения с сервером и создание объекта
``` golang
type Server struct {
	client srv.UserServiceClient
	ctx    context.Context
}
```

- `user.go` - все, необходимые боту действия с базой данных
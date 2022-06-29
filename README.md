# SandBox MarketBot
## Основная идея
Бот, с помощью которого реализуется взаимодействие с песочницей tinkoff investapi

## Микросервисы
Здесь находится сразу несколько микросервисов, потому что они все небольшие, и в одном репозитории легче разглядеть то, как они работают и друг с другом взаимодействуют.
- [myinvest](/internal/myinvest/README.md) - передает только нужные данные с tinkoff investapi
- [telegram](/internal/bot/README.md) - телеграм бот
- [server](/internal/server/README.md) - ручки для общения с базой данных

Возможно для поставленной задачи не было необходимости в создании такого количества микросервисах, но мне хотелось по-лучше разобраться в том, как работает gRPC и protobuff


## Вью
В качестве интерфейса используется телеграм бот.
- Регистрация в боте происходит через предоставления sandbox ключей для подтягивания данных с TinkoffAPI
- Через бот можно будет посылать запросы о курсе валют и пополнению счета


### Граф
``` mermaid
flowchart TB
    subgraph Myinvest
    Myinvest_server-->InvestAPI_client
    end
    InvestAPI_client-->InvestAPI

    subgraph Server
    Server_hands
    end
    Server_hands--->PostgreSQL
    
    subgraph TelegramBot
    Telegram_client-->Myinvest_client
    Telegram_client-->Server_client
    end
    Myinvest_client-->Myinvest_server
    Server_client-->Server_hands
    TelegramBot---->TelegramAPI
```



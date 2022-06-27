# SandBox MarketBot
Здесь находится сразу несколько микросервисов, потому что они все небольшие, и в одном репозитории легче разглядеть то, как они работают и друг с другом взаимодействуют.
## Основная идея
Бот, с помощью которого реализуется взаимодействие с песочницей tinkoff investapi



## Вью
В качестве интерфейса будет использоваться телеграм бот.
- Регистрация в боте происходит через предоставления sandbox ключей для подтягивания данных с TinkoffAPI
- Через бот можно будет посылать запросы на покупку активов
- Бот по запросу будет посылать имеющиеся активы


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



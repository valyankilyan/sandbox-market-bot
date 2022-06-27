# SandBox MarketBot

## Основная идея
Бот, с помощью которого реализуется взаимодействие с песочницей tinkoff investapi



## Вью
В качестве интерфейса будет использоваться телеграм бот.
- Регистрация в боте происходит через предоставления sandbox ключей для подтягивания данных с TinkoffAPI
- Через бот можно будет посылать запросы на покупку активов
- Бот по запросу будет посылать имеющиеся активы


### Граф
``` mermaid
graph TD;
    TinkoffAPI --gRPC--> TinkoffAPI_Hand
    TinkoffAPI_Hand --gRPC--> Communicator
    Communicator --gRPC--> TelegramBot_Hand


    TelegramBot_Hand --> Server
    TinkoffAPI_Hand --> Server
    Communicator --> Server    
    Server --> PostgreSQL
```



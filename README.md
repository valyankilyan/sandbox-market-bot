# Homework 2. Market bot

## Основная идея
Бот с реализацией работы с песочницей tinkoff investapi
- Сервер будет отдавать данные с дб
- API handler будет собирать данные о состояннии настоящей стоимости портфеля
- Коммуникатор будет обрабатывать эти данные
- Вьюшка будет собирать данные (с апи или коммуникатора) с коммуникатора и присылать данные о портфеле потребителю
 

## Вью
В качестве интерфейса будет использоваться телеграм бот.
- Регистрация в боте происходит через предоставления sandbox ключей для подтягивания данных с TinkoffAPI
- Через бот можно будет посылать запросы на покупку активов
- Бот по запросу будет посылать имеющиеся активы
- По написании команды (скажем /status) будет выводится поисковая строка со списком возможных активов.


## Реализация
В реализации планируется использовать несколько различных модулей.

### Связной между view и api хендлерами
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
<!-- BinanceAPI --HTTP-- BinanceAPI_Hand
BinanceAPI_Hand --gRPC-- Communicator -->
<!-- Communicator --gRPC-- WEBSite -->


